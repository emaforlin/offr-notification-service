package service

import (
	"github.com/emaforlin/notification-service/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	gomail "gopkg.in/gomail.v2"
)

const (
	MAIL_QUEUE_BUFFER = 100
)

type NotificationService struct {
	logger      *zap.Logger
	emailDialer *gomail.Dialer
	mq          chan *gomail.Message
}

func NewNotificationService(l *zap.Logger) *NotificationService {
	cfg := config.GetConfig()
	d := gomail.NewDialer(cfg.SMTP.Host, int(cfg.SMTP.Port), cfg.SMTP.User, cfg.SMTP.Pass)

	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	service := &NotificationService{
		logger:      l,
		emailDialer: d,
		mq:          make(chan *gomail.Message, MAIL_QUEUE_BUFFER),
	}

	go service.startMailDaemon()

	return service
}

func ProvideNotificationService() fx.Option {
	return fx.Provide(NewNotificationService)
}
