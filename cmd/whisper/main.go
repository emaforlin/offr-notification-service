package main

import (
	"github.com/emaforlin/notification-service/config"
	"github.com/emaforlin/notification-service/grpc"
	"github.com/emaforlin/notification-service/logger"
	"github.com/emaforlin/notification-service/service"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		logger.ProvideZapLogger(),
		config.ProvideConfig(),
		service.ProvideNotificationService(),
		grpc.StartGRPCServer(),
	).Run()
}
