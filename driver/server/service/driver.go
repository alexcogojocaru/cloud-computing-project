package service

import (
	"context"
	"net"
	"time"

	"github.com/alexcogojocaru/cloud-computing-project/driver/pb"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type DriverGrpcService struct {
	pb.UnimplementedDriverServer

	log *zap.Logger
	rdb *redis.Client
}

func NewDriverGrpcService(
	lc fx.Lifecycle,
	log *zap.Logger,
	rdb *redis.Client,
) *DriverGrpcService {
	d := &DriverGrpcService{
		log: log,
		rdb: rdb,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", "0.0.0.0:8081")
			if err != nil {
				log.Fatal("Cannot listen to port 8081", zap.Error(err))
			}

			grpcServer := grpc.NewServer()
			pb.RegisterDriverServer(grpcServer, d)

			go func() {
				log.Info("Starting grpc service...")
				if err := grpcServer.Serve(lis); err != nil {
					log.Fatal("Failed to serve", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {

			return nil
		},
	})

	return d
}

func (d *DriverGrpcService) GetClosest(ctx context.Context, location *pb.LocationMetadata) (*pb.DriverLocationList, error) {
	d.log.Info("Received getClosest request", zap.String("method", "GetClosest"))
	geoLocations, err := GetClosestDriver(d.rdb, location)
	if err != nil {
		return nil, err
	}

	driverLocations := &pb.DriverLocationList{
		Locations: make([]*pb.DriverLocation, len(geoLocations)),
	}

	for idx, geoLocation := range geoLocations {
		driverLocations.Locations[idx] = &pb.DriverLocation{
			Name:      geoLocation.Name,
			Latitude:  geoLocation.Latitude,
			Longitude: geoLocation.Longitude,
			Distance:  geoLocation.Dist,
		}
	}

	d.log.Info("Found drivers", zap.Int("size", len(driverLocations.Locations)))

	return driverLocations, nil
}

func (d *DriverGrpcService) GetStatus(ctx context.Context, metadata *pb.DriverStatusMetadata) (*pb.DriverStatusMetadata, error) {
	value, err := d.rdb.Get(ctx, metadata.Name).Result()
	if err != nil {
		// d.log.Warn("No data in redis for driver", zap.String("driver", metadata.Name))
		// return &pb.DriverStatusMetadata{
		// 	Status: pb.DriverStatus_UNKNOWN,
		// }, err

		// for demo purposes, set the driver's status to FREE
		d.rdb.Set(ctx, metadata.Name, pb.DriverStatus_FREE.String(), 0*time.Second).Err()
	}

	var status pb.DriverStatus
	if value == "BUSY" {
		status = pb.DriverStatus_BUSY
	} else {
		status = pb.DriverStatus_FREE
	}

	d.log.Info(
		"GetStatus",
		zap.String("name", metadata.Name),
		zap.String("redisValue", value),
		zap.String("status", status.String()),
	)

	return &pb.DriverStatusMetadata{
		Name:   metadata.Name,
		Status: status,
	}, nil
}

func (d *DriverGrpcService) SetStatus(ctx context.Context, metadata *pb.DriverStatusMetadata) (*pb.Empty, error) {
	err := d.rdb.Set(ctx, metadata.Name, metadata.Status.String(), 0*time.Second).Err()
	if err != nil {
		return nil, err
	}

	d.log.Info("SetStatus", zap.String("drivername", metadata.Name), zap.String("status", metadata.Status.String()))

	return &pb.Empty{}, nil
}

func GetClosestDriver(rdb *redis.Client, location *pb.LocationMetadata) ([]redis.GeoLocation, error) {
	locations, err := rdb.GeoRadius(
		context.Background(),
		"drivers/location",
		location.Longitude,
		location.Latitude,
		&redis.GeoRadiusQuery{
			Radius:      location.Radius,
			Unit:        "km",
			WithGeoHash: true,
			WithCoord:   true,
			WithDist:    true,
			Count:       10,
			Sort:        "ASC",
		},
	).Result()
	if err != nil {
		return nil, err
	}

	return locations, nil
}
