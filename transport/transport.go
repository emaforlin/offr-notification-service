package transport

import (
	"context"

	"github.com/emaforlin/notification-service/endpoints"
	"github.com/emaforlin/notification-service/models"
	"github.com/emaforlin/notification-service/pb"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	pb.UnimplementedNotificationServer
	sendEmailNotification gt.Handler
}

func (s *gRPCServer) SendEmailNotification(ctx context.Context, req *pb.SendEmailNotificationReq) (*pb.SendEmailNotificationRes, error) {
	_, resp, err := s.sendEmailNotification.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SendEmailNotificationRes), nil
}

func decodeSendEmailReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SendEmailNotificationReq)
	return models.EmailDto{
		Recipient: req.GetRecipient(),
		Subject:   req.GetSubject(),
		Content:   req.GetBody(),
	}, nil
}

func encodeSendEmailRes(_ context.Context, _ interface{}) (interface{}, error) {
	return &pb.SendEmailNotificationRes{}, nil
}

func NewGRPCServer(endpoints endpoints.Endpoints) pb.NotificationServer {
	return &gRPCServer{
		sendEmailNotification: gt.NewServer(
			endpoints.SendEmailNotification,
			decodeSendEmailReq,
			encodeSendEmailRes,
		),
	}
}
