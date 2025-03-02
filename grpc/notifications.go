package grpc

import (
	"context"

	"github.com/emaforlin/notification-service/dto"
	"github.com/emaforlin/notification-service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GrpcServer) SendEmailNotification(ctx context.Context, req *pb.SendEmailNotificationReq) (*pb.SendEmailNotificationRes, error) {
	err := s.notificationSvc.SendEmailNotification(ctx, dto.EmailDto{
		Recipient: req.GetRecipient(),
		Subject:   req.GetSubject(),
		Content:   req.GetBody(),
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}
	return &pb.SendEmailNotificationRes{}, nil
}
