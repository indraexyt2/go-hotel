package repositories

import (
	"context"
	"gorm.io/gorm"
	"hotel-rooms/internal/models"
)

type RoomFeaturesRepository struct {
	DB *gorm.DB
}

func NewRoomFeaturesRepository(db *gorm.DB) *RoomFeaturesRepository {
	return &RoomFeaturesRepository{
		DB: db,
	}
}

func (r *RoomFeaturesRepository) AddRoomFeature(ctx context.Context, roomFeature *models.RoomFeature) error {
	return r.DB.WithContext(ctx).Create(roomFeature).Error
}

func (r *RoomFeaturesRepository) GetAllRoomFeatures(ctx context.Context, roomTypeID int) ([]models.RoomFeature, error) {
	var roomFeatures []models.RoomFeature
	err := r.DB.WithContext(ctx).Where("room_type_id = ?", roomTypeID).Find(&roomFeatures).Error
	if err != nil {
		return nil, err
	}
	return roomFeatures, nil
}

func (r *RoomFeaturesRepository) EditRoomFeature(ctx context.Context, roomFeatureID int, newData map[string]interface{}) error {
	return r.DB.Debug().WithContext(ctx).Model(&models.RoomFeature{}).Where("id = ?", roomFeatureID).Updates(newData).Error
}

func (r *RoomFeaturesRepository) DeleteRoomFeature(ctx context.Context, roomFeatureID int) error {
	return r.DB.WithContext(ctx).Where("id = ?", roomFeatureID).Delete(&models.RoomFeature{}).Error
}
