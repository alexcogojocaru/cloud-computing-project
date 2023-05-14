package main

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type DriverService struct {
	log          *zap.Logger
	rdb          *redis.Client
	PubSubClient *pubsub.Client
}

type DriverDetails struct {
	ID     string `json:"id"`
	Coords struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
}

const (
	PUBSUB_SUBSCRIPTION = "rider-streaming-sub"
	GCP_PROJECT         = "cloudcomputing-386413"
	DEFAULT_CACHE_TTL   = 5 * time.Minute
)

func NewDriverService(
	lc fx.Lifecycle,
	log *zap.Logger,
	rdb *redis.Client,
) *DriverService {
	pubsubClient, err := pubsub.NewClient(context.Background(), GCP_PROJECT)
	if err != nil {
		return nil
	}

	rs := &DriverService{
		log:          log,
		rdb:          rdb,
		PubSubClient: pubsubClient,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			rs.log.Info("Starting ride service...")
			go rs.Start(context.Background())
			return nil
		},
		OnStop: func(ctx context.Context) error {
			rs.PubSubClient.Close()

			rs.log.Info("Shutting down ride service...")
			return nil
		},
	})

	return rs
}

func (rs *DriverService) Start(ctx context.Context) error {
	subscription := rs.PubSubClient.Subscription(PUBSUB_SUBSCRIPTION)
	err := subscription.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		var details DriverDetails
		json.Unmarshal(m.Data, &details)

		rs.log.Info(
			"Received data",
			zap.String("msgID", m.ID),
			zap.String("subscription", PUBSUB_SUBSCRIPTION),
			zap.Any("data", details),
		)

		// cache the driver's location
		err := rs.rdb.GeoAdd(ctx, "drivers/location", &redis.GeoLocation{
			Name:      details.ID,
			Longitude: details.Coords.Longitude,
			Latitude:  details.Coords.Latitude,
		}).Err()
		if err != nil {
			rs.log.Error(
				"Cannot cache the driver's location",
				zap.String("driverid", details.ID),
				zap.Error(err),
			)
		}

		m.Ack()
	})

	return err
}
