package main

import (
	"github.com/alexcogojocaru/cloud-computing-project/ride/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			NewRideService,
			service.NewRideGrpcService,
			zap.NewExample,
		), fx.Invoke(
			func(*RideService) {},
			func(*service.RideGrpcService) {},
		),
	).Run()
}
