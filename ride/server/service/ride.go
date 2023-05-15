package service

import (
	"context"
	"net"

	"github.com/alexcogojocaru/cloud-computing-project/ride/proto-gen/pb"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type RideGrpcService struct {
	pb.UnimplementedRideServer

	log          *zap.Logger
	driverConn   *grpc.ClientConn
	driverClient pb.DriverClient
}

func NewRideGrpcService(
	lc fx.Lifecycle,
	log *zap.Logger,
) *RideGrpcService {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		return nil
	}

	client := pb.NewDriverClient(conn)
	r := &RideGrpcService{
		log:          log,
		driverConn:   conn,
		driverClient: client,
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

func (r *RideGrpcService) Start(ctx context.Context, location *pb.LocationMetadata) (*pb.StartRideResponse, error) {
	driverData, err := r.driverClient.GetClosest(ctx, location)
	if err != nil {
		return nil, err
	}

	if len(driverData.Locations) == 0 {
		return &pb.StartRideResponse{
			Matched: false,
		}, nil
	}

	var closestDriver *pb.DriverLocation
	for _, driver := range driverData.Locations {
		metadata, err := r.driverClient.GetStatus(ctx, &pb.DriverStatusMetadata{
			Name: driver.Name,
		})
		if err != nil {
			r.log.Fatal(
				"Cannot retrieve the driver's status",
				zap.String("drivername", driver.Name),
			)
		}

		if metadata.Status == pb.DriverStatus_FREE {
			closestDriver = driver
			break
		}
	}

	if closestDriver == nil {
		return &pb.StartRideResponse{
			Matched: false,
		}, nil
	}

	return &pb.StartRideResponse{
		Matched:  true,
		Location: closestDriver,
	}, nil
}
