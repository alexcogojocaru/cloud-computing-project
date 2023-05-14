package main

import (
	"github.com/alexcogojocaru/cloud-computing-project/driver/service"
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
			NewDriverService,
			NewRedisClient,
			service.NewDriverGrpcService,
			zap.NewExample,
		), fx.Invoke(
			func(*DriverService) {},
			func(*service.DriverGrpcService) {},
		),
	).Run()
}
