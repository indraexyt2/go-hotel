package interfaces

import (
	"context"
	"hotel-notification/internal/models"
)

type INotificationRepository interface {
	GetTemplate(ctx context.Context, templateName string) (models.NotificationTemplate, error)
	InsertNotificationHistory(ctx context.Context, history *models.NotificationHistory) error
}

type INotificationService interface {
	SendNotification(ctx context.Context, req *models.InternalNotificationRequest) error
}

type INotificationHandler interface {
	SendNotification(ctx context.Context, req *models.InternalNotificationRequest) error
}
