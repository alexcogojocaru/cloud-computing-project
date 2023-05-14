package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
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

func main() {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, "cloudcomputing-386413")
	if err != nil {
		log.Fatal(err)
	}
	defer pubsubClient.Close()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

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

			log.Println(id)
		}()
	}

	wg.Wait()
}