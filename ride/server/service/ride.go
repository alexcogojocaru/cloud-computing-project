package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/alexcogojocaru/cloud-computing-project/ride/proto-gen/pb"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type RideGrpcService struct {
	pb.UnimplementedRideServer

	log          *zap.Logger
	driverConn   *grpc.ClientConn
	driverClient pb.DriverClient
	rdb          *redis.Client
	pubsub       *pubsub.Client
	firestore    *FirestoreWrapper
}

type NotificationMessage struct {
	RideID     string    `json:"rideid"`
	RiderName  string    `json:"rider"`
	DriverName string    `json:"driver"`
	Distance   float64   `json:"distance"`
	Timestamp  time.Time `json:"timestamp"`
}

var (
	DRIVER_ADDR = os.Getenv("DRIVER_ADDR")
)

func NewRideGrpcService(
	lc fx.Lifecycle,
	log *zap.Logger,
) *RideGrpcService {
	if DRIVER_ADDR == "" {
		DRIVER_ADDR = "localhost:8081"
	}

	conn, err := grpc.Dial(DRIVER_ADDR, grpc.WithInsecure())
	if err != nil {
		return nil
	}

	pubsubClient, err := pubsub.NewClient(context.Background(), "cloudcomputing-386413")
	if err != nil {
		return nil
	}

	client := pb.NewDriverClient(conn)
	r := &RideGrpcService{
		log:          log,
		driverConn:   conn,
		driverClient: client,
		rdb:          NewRedisClient(),
		pubsub:       pubsubClient,
		firestore:    NewFirestoreDb(),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", ":8082")
			if err != nil {
				log.Fatal("Cannot listen to port 8082", zap.Error(err))
			}

			server := grpc.NewServer()
			pb.RegisterRideServer(server, r)

			go func() {
				log.Info("Starting grpc service...")
				if err := server.Serve(lis); err != nil {
					log.Fatal("Failed to serve", zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			r.driverConn.Close()
			return nil
		},
	})

	return r
}

func (r *RideGrpcService) Start(location *pb.StartRideRequest, stream pb.Ride_StartServer) error {
	ctx := context.Background()
	r.log.Info("Received start request", zap.String("method", "Start"))
	driverData, err := r.driverClient.GetClosest(ctx, location.StartLocation)
	if err != nil {
		return err
	}

	if len(driverData.Locations) == 0 {
		return errors.New("No drivers available")
	}

	r.log.Info("Received driver data", zap.Int("size", len(driverData.Locations)))

	var closestDriver *pb.DriverLocation
	for _, driver := range driverData.Locations {
		metadata, err := r.driverClient.GetStatus(ctx, &pb.DriverStatusMetadata{
			Name: driver.Name,
		})
		if err != nil {
			r.log.Fatal(
				"Cannot retrieve the driver's status",
				zap.String("drivername", driver.Name),
				zap.Error(err),
			)
		}

		r.log.Info("Driver metadata", zap.Any("driver", metadata))

		if metadata.Status == pb.DriverStatus_FREE {
			closestDriver = driver
			break
		}
	}

	r.log.Info("Closest driver", zap.Any("driver", closestDriver))

	if closestDriver == nil {
		return errors.New("No free drivers")
	}

	r.driverClient.SetStatus(ctx, &pb.DriverStatusMetadata{
		Name:   closestDriver.Name,
		Status: pb.DriverStatus_BUSY,
	})

	r.rdb.GeoAdd(ctx, "riders/location", &redis.GeoLocation{
		Name:      fmt.Sprintf("start-%s", location.Username),
		Longitude: location.StartLocation.Longitude,
		Latitude:  location.StartLocation.Latitude,
	}, &redis.GeoLocation{
		Name:      fmt.Sprintf("stop-%s", location.Username),
		Longitude: location.EndLocation.Longitude,
		Latitude:  location.EndLocation.Latitude,
	})

	r.log.Info("Starting ride",
		zap.String("rider", location.Username),
		zap.String("driver", closestDriver.Name),
	)

	rideId := uuid.New().String()

	stream.Send(&pb.StartRideResponse{
		Matched:  true,
		Location: closestDriver,
	})

	startName := fmt.Sprintf("start-%s", location.Username)
	stopName := fmt.Sprintf("stop-%s", location.Username)

	distance, err := r.rdb.GeoDist(context.Background(), "riders/location", startName, stopName, "km").Result()
	if err != nil {
		r.log.Error("Cannot compute the distance", zap.Error(err))
	}

	r.log.Info("Computed the ride distance", zap.Float64("distance", distance))

	for i := 0; i < int(distance); i++ {
		err = stream.Send(&pb.StartRideResponse{
			Matched:  true,
			Location: closestDriver,
		})
		if err != nil {
			return err
		}

		r.log.Info("Ongoing ride",
			zap.String("rider", location.Username),
			zap.String("driver", closestDriver.Name),
		)

		time.Sleep(1 * time.Second)
	}

	r.log.Info("Finished ride",
		zap.String("rider", location.Username),
		zap.String("driver", closestDriver.Name),
	)

	r.driverClient.SetStatus(ctx, &pb.DriverStatusMetadata{
		Name:   closestDriver.Name,
		Status: pb.DriverStatus_FREE,
	})

	msg := NotificationMessage{
		RideID:     rideId,
		RiderName:  location.Username,
		DriverName: closestDriver.Name,
		Distance:   distance,
		Timestamp:  time.Now(),
	}
	details, _ := json.Marshal(msg)
	topic := r.pubsub.Topic("notification-stream")
	res := topic.Publish(ctx, &pubsub.Message{
		Data: details,
	})
	serverid, _ := res.Get(ctx)
	r.log.Info("PubSub id", zap.String("id", serverid))

	_, _, err = r.firestore.Client.Collection("rides").Add(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}
