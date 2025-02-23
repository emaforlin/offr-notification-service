package service

import (
	"context"

	"github.com/emaforlin/notification-service/config"
	"github.com/emaforlin/notification-service/models"
	"go.uber.org/zap"
	gomail "gopkg.in/gomail.v2"
)

const (
	MAIL_QUEUE_BUFFER = 100
)

type MailService interface {
	SendEmailNotification(ctx context.Context, data models.EmailDto) error
}

type mailService struct {
	logger      *zap.Logger
	emailDialer *gomail.Dialer
	mq          chan *gomail.Message
}

func NewNotificationService(l *zap.Logger) MailService {
	cfg := config.GetConfig()
	d := gomail.NewDialer(cfg.SMTP.Host, int(cfg.SMTP.Port), cfg.SMTP.User, cfg.SMTP.Pass)

	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	service := &mailService{
		logger:      l,
		emailDialer: d,
		mq:          make(chan *gomail.Message, MAIL_QUEUE_BUFFER),
	}

	go service.startMailDaemon()

	return service
}
