package repositories

import (
	"context"
	"gorm.io/gorm"
	"hotel-rooms/internal/models"
)

type RoomPhotosRepository struct {
	DB *gorm.DB
}

func NewRoomPhotosRepository(db *gorm.DB) *RoomPhotosRepository {
	return &RoomPhotosRepository{DB: db}
}

func (r *RoomPhotosRepository) AddRoomTypePhotos(ctx context.Context, roomPhotos []models.RoomPhoto) error {
	return r.DB.WithContext(ctx).Create(&roomPhotos).Error
}

func (r *RoomPhotosRepository) GetRoomTypePhotos(ctx context.Context, roomTypeID int) ([]models.RoomPhoto, error) {
	var roomPhotos []models.RoomPhoto
	err := r.DB.WithContext(ctx).Where("room_type_id = ?", roomTypeID).Find(&roomPhotos).Error
	if err != nil {
		return nil, err
	}
	return roomPhotos, nil
}

func (r *RoomPhotosRepository) DeletePhotos(ctx context.Context, photosID int) error {
	return r.DB.WithContext(ctx).Where("id = ?", photosID).Delete(&models.RoomPhoto{}).Error
}
