package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/emaforlin/notification-service/config"
	"github.com/emaforlin/notification-service/endpoints"
	"github.com/emaforlin/notification-service/pb"
	"github.com/emaforlin/notification-service/service"
	"github.com/emaforlin/notification-service/transport"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.Init()
	cfg := config.GetConfig()

	logger := zap.NewExample()
	defer logger.Sync()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	svc := service.NewNotificationService(logger)
	endpoints := endpoints.MakeEndpoints(svc)

	grpcServer := transport.NewGRPCServer(endpoints)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		logger.Fatal("Failed to start the server", zap.Error(err))
	}

	baseServer := grpc.NewServer()
	pb.RegisterNotificationServer(baseServer, grpcServer)

	reflection.Register(baseServer)

	go func() {
		if err := baseServer.Serve(listener); err != nil {
			logger.Fatal("Failed to start the server", zap.Error(err))
		}
	}()

	logger.Info("Server started successfully")
	logger.Info("Listening on", zap.Uint16("port", cfg.App.Port))

	// Handle graceful shutdown
	<-quit
	logger.Info("Shutting down the server")
	baseServer.GracefulStop()
	listener.Close()
	close(quit)
	logger.Info("Server stopped successfully")

}
