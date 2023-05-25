package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/alexcogojocaru/cloud-computing-project/driver/client/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	PUBSUB_TOPIC    = "rider-streaming"
	LOWER_LATITUDE  = 47.14239230121294
	UPPER_LATITUDE  = 47.183312148060274
	LOWER_LONGITUDE = 27.531191096268163
	UPPER_LONGITUDE = 27.660349147474353
)

type GeolocationCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type DriverDetails struct {
	ID     string                 `json:"id"`
	Coords GeolocationCoordinates `json:"coords"`
}

var (
	DRIVER_SERVICE_ADDR = os.Getenv("DRIVER_SERVICE_ADDR")
)

func main() {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, "cloudcomputing-386413")
	if err != nil {
		log.Fatal(err)
	}
	defer pubsubClient.Close()

	conn, err := grpc.Dial("cc-driver-server-qqtr4ife4a-lz.a.run.app:443", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewDriverClient(conn)
	drivername := uuid.New().String()

	_, err = client.SetStatus(ctx, &pb.DriverStatusMetadata{
		Name:   drivername,
		Status: pb.DriverStatus_FREE,
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		latitude := rand.Float64()*(UPPER_LATITUDE-LOWER_LATITUDE) + LOWER_LATITUDE
		longitude := rand.Float64()*(UPPER_LONGITUDE-LOWER_LONGITUDE) + LOWER_LONGITUDE

		details, _ := json.Marshal(DriverDetails{
			ID: uuid.New().String(),
			Coords: GeolocationCoordinates{
				Latitude:  latitude,
				Longitude: longitude,
			},
		})

		topic := pubsubClient.Topic(PUBSUB_TOPIC)
		res := topic.Publish(ctx, &pubsub.Message{
			Data: details,
		})
		id, err := res.Get(ctx)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("name=%s pubsubId=%s\n", drivername, id)
		time.Sleep(10 * time.Second)
	}
}
