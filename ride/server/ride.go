package main

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RideService struct {
	log *zap.Logger
}

const (
	PUBSUB_SUBSCRIPTION = "rider-streaming-sub"
	GCP_PROJECT         = "cloudcomputing-386413"
)

func NewRideService(lc fx.Lifecycle, log *zap.Logger) *RideService {
	rs := &RideService{
		log: log,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			rs.log.Info("Starting ride service...")
			go rs.Start(context.Background())
			return nil
		},
		OnStop: func(ctx context.Context) error {
			rs.log.Info("Shutting down ride service...")
			return nil
		},
	})

	return rs
}

func (rs *RideService) Start(ctx context.Context) error {
	return nil
}
