package main

import (
	"context"
	"io"
	"log"

	"github.com/alexcogojocaru/cloud-computing-project/ride/client/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRideClient(conn)

	stream, err := client.Start(context.Background(), &pb.StartRideRequest{
		Username: "alexcogojocaru",
		StartLocation: &pb.LocationMetadata{
			Latitude:  47.16129960502986,
			Longitude: 27.590637972547764,
			Radius:    2,
		},
		EndLocation: &pb.LocationMetadata{
			Latitude:  47.172983080034896,
			Longitude: 27.54453466623929,
			Radius:    2,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
	}
}
