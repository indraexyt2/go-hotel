package services

import (
	"bytes"
	"context"
	"hotel-notification/helpers"
	"hotel-notification/internal/interfaces"
	"hotel-notification/internal/models"
	"html/template"
)

type NotificationService struct {
	NotificationRepository interfaces.INotificationRepository
}

func NewNotificationService(notificationRepo interfaces.INotificationRepository) *NotificationService {
	return &NotificationService{
		NotificationRepository: notificationRepo,
	}
}

func (s *NotificationService) SendNotification(ctx context.Context, req *models.InternalNotificationRequest) error {
	emailTemplate, err := s.NotificationRepository.GetTemplate(ctx, req.TemplateName)
	if err != nil {
		return err
	}

	parse, err := template.New("emailTemplate").Parse(emailTemplate.Body)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	err = parse.Execute(&tpl, req.Placeholder)
	if err != nil {
		return err
	}

	email := helpers.Email{
		To:      req.Recipient,
		Subject: emailTemplate.Subject,
		Body:    tpl.String(),
	}

	err = email.SendEmailNotification()
	if err != nil {
		notifyHistory := &models.NotificationHistory{
			Recipient:    req.Recipient,
			TemplateID:   emailTemplate.ID,
			Status:       "failed",
			ErrorMessage: err.Error(),
		}

		err2 := s.NotificationRepository.InsertNotificationHistory(ctx, notifyHistory)
		if err2 != nil {
			return err2
		}

		return err
	}

	notifyHistory := &models.NotificationHistory{
		Recipient:  req.Recipient,
		TemplateID: emailTemplate.ID,
		Status:     "success",
	}

	err = s.NotificationRepository.InsertNotificationHistory(ctx, notifyHistory)
	if err != nil {
		return err
	}

	return nil
}
