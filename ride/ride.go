package main

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RideService struct {
	log          *zap.Logger
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
)

func NewRideService(lc fx.Lifecycle, log *zap.Logger) *RideService {
	pubsubClient, err := pubsub.NewClient(context.Background(), GCP_PROJECT)
	if err != nil {
		return nil
	}

	rs := &RideService{
		log:          log,
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

func (rs *RideService) Start(ctx context.Context) error {
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
		m.Ack()
	})

	return err
}
