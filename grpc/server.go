package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/emaforlin/notification-service/config"
	"github.com/emaforlin/notification-service/pb"
	"github.com/emaforlin/notification-service/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	pb.UnimplementedNotificationServer
	notificationSvc *service.NotificationService
}

func NewGRPCServer(cfg *config.Config, nt *service.NotificationService) *GrpcServer {
	return &GrpcServer{
		notificationSvc: nt,
	}
}

func StartGRPCServer() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, cfg *config.Config, notificationSvc *service.NotificationService, log *zap.Logger) {
		grpcServer := grpc.NewServer()

		serverImpl := NewGRPCServer(cfg, notificationSvc)

		// DISABLE for prod
		reflection.Register(grpcServer)

		pb.RegisterNotificationServer(grpcServer, serverImpl)

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				addr := fmt.Sprintf(":%d", cfg.App.Port)
				lis, err := net.Listen("tcp", addr)
				if err != nil {
					return err
				}

				go grpcServer.Serve(lis)
				log.Info("Listening on ", zap.String("address", addr))
				return nil
			},
			OnStop: func(ctx context.Context) error {
				grpcServer.GracefulStop()
				log.Info("Server stopped.")
				return nil
			},
		})

	})
}
