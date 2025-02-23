package endpoints

import (
	"context"

	"github.com/emaforlin/notification-service/models"
	"github.com/emaforlin/notification-service/service"
	gkendpoint "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SendEmailNotification gkendpoint.Endpoint
}

func MakeEndpoints(s service.MailService) Endpoints {
	return Endpoints{
		SendEmailNotification: makeSendEmailNotificationEndpoint(s),
	}
}

func makeSendEmailNotificationEndpoint(s service.MailService) gkendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.EmailDto)
		err = s.SendEmailNotification(ctx, req)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
