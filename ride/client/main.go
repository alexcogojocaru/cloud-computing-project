package main

import (
	"context"
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

	resp, err := client.Start(context.Background(), &pb.LocationMetadata{
		Latitude:  47.16129960502986,
		Longitude: 27.590637972547764,
		Radius:    2,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
}
