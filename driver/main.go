package main

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
)

const PUBSUB_TOPIC = "rider-streaming"

type GeolocationCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type DriverDetails struct {
	ID     string                 `json:"id"`
	Coords GeolocationCoordinates `json:"coords"`
}

func main() {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, "cloudcomputing-386413")
	if err != nil {
		log.Fatal(err)
	}
	defer pubsubClient.Close()

	details, _ := json.Marshal(DriverDetails{
		ID: uuid.New().String(),
		Coords: GeolocationCoordinates{
			Latitude:  30.2,
			Longitude: 27.1,
		},
	})

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			topic := pubsubClient.Topic(PUBSUB_TOPIC)
			res := topic.Publish(ctx, &pubsub.Message{
				Data: details,
			})
			id, err := res.Get(ctx)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(id)
		}()
	}

	wg.Wait()
}
