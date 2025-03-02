package service

import (
	"context"
	"time"

	"github.com/emaforlin/notification-service/dto"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

func (ms *NotificationService) startMailDaemon() {
	var s gomail.SendCloser
	var err error
	open := false
	for {
		select {
		case m, ok := <-ms.mq:
			if !ok {
				return
			}
			if !open {
				if s, err = ms.emailDialer.Dial(); err != nil {
					ms.logger.Error("Failed to open SMTP connection", zap.Error(err))
				}
				open = true
			}
			if err := gomail.Send(s, m); err != nil {
				ms.logger.Error("Failed to send email", zap.Error(err))
			}
			ms.logger.Info("Email sent!")

		case <-time.After(30 * time.Second):
			if open {
				if err := s.Close(); err != nil {
					ms.logger.Error("Error closing SMTP connection", zap.Error(err))
				}
				open = false
			}
		}
	}
}

func (ms *NotificationService) Stop() {
	close(ms.mq)
}

// SendEmailNotification implements Service.
func (s *NotificationService) SendEmailNotification(ctx context.Context, data dto.EmailDto) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.emailDialer.Username)
	msg.SetHeader("To", data.Recipient)
	msg.SetHeader("Subject", data.Subject)
	msg.SetBody("text/html", data.Content)

	s.mq <- msg

	return nil
}
