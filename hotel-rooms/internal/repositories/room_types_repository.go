package repositories

import (
	"context"
	"gorm.io/gorm"
	"hotel-rooms/internal/models"
)

type RoomTypesRepository struct {
	DB *gorm.DB
}

func NewRoomTypesRepository(db *gorm.DB) *RoomTypesRepository {
	return &RoomTypesRepository{
		DB: db,
	}
}

func (r *RoomTypesRepository) GetAllRoomTypes(ctx context.Context) ([]models.RoomType, error) {
	var roomTypes []models.RoomType
	err := r.DB.WithContext(ctx).Preload("RoomPhotos").Preload("RoomFeatures").Find(&roomTypes).Error
	if err != nil {
		return nil, err
	}
	return roomTypes, nil
}

func (r *RoomTypesRepository) GetRoomTypesDetails(ctx context.Context, id int) (*models.RoomType, error) {
	var roomType models.RoomType
	r.DB.WithContext(ctx).Model(&models.RoomType{}).Where("id = ?", id).Preload("RoomPhotos").Preload("RoomFeatures").First(&roomType)
	return &roomType, nil
}

func (r *RoomTypesRepository) AddRoomType(ctx context.Context, roomType *models.RoomType) error {
	return r.DB.WithContext(ctx).Create(roomType).Error
}
