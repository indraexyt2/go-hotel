package repositories

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"hotel-notification/internal/models"
)

type NotificationRepository struct {
	DB  *gorm.DB
	RDB *redis.ClusterClient
}

func NewNotificationRepository(db *gorm.DB, RDB *redis.ClusterClient) *NotificationRepository {
	return &NotificationRepository{
		DB:  db,
		RDB: RDB,
	}
}

func (r *NotificationRepository) GetTemplate(ctx context.Context, templateName string) (models.NotificationTemplate, error) {
	var template models.NotificationTemplate
	err := r.DB.WithContext(ctx).Where("template_name = ?", templateName).First(&template).Error
	return template, err
}

func (r *NotificationRepository) InsertNotificationHistory(ctx context.Context, history *models.NotificationHistory) error {
	return r.DB.WithContext(ctx).Create(history).Error
}
