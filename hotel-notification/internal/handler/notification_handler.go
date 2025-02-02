package handler

import (
	"context"
	"hotel-notification/helpers"
	"hotel-notification/internal/interfaces"
	"hotel-notification/internal/models"
)

type NotificationHandler struct {
	NotificationService interfaces.INotificationService
}

func NewNotificationHandler(notificationService interfaces.INotificationService) *NotificationHandler {
	return &NotificationHandler{
		NotificationService: notificationService,
	}
}

func (s *NotificationHandler) SendNotification(ctx context.Context, req *models.InternalNotificationRequest) error {
	var (
		log = helpers.Logger
	)

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return err
	}

	return s.NotificationService.SendNotification(ctx, req)
}
