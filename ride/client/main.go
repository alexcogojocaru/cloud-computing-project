package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alexcogojocaru/cloud-computing-project/ride/client/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewRideClient(conn)
	defer conn.Close()

	resp, err := client.Start(context.Background(), &pb.LocationMetadata{
		Latitude:  47.16129960502986,
		Longitude: 27.590637972547764,
		Radius:    2,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
