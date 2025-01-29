package repositories

import (
	"context"
	"gorm.io/gorm"
	"hotel-rooms/internal/models"
)

type RoomsRepository struct {
	DB *gorm.DB
}

func NewRoomsRepository(db *gorm.DB) *RoomsRepository {
	return &RoomsRepository{DB: db}
}

func (r *RoomsRepository) AddRoom(ctx context.Context, room *models.Room) error {
	return r.DB.WithContext(ctx).Create(room).Error
}

func (r *RoomsRepository) GetAllRooms(ctx context.Context) ([]models.Room, error) {
	var rooms []models.Room
	subquery := r.DB.Table("rooms").
		Select("MIN(id) as id").
		Group("room_type_id")

	err := r.DB.WithContext(ctx).
		Joins("JOIN (?) as r2 ON rooms.id = r2.id", subquery).
		Preload("RoomType.RoomPhotos").Preload("RoomType.RoomFeatures").
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *RoomsRepository) GetRoomDetails(ctx context.Context, id int) (*models.Room, error) {
	var room models.Room
	err := r.DB.WithContext(ctx).Where("id = ?", id).Preload("RoomType.RoomPhotos").Preload("RoomType.RoomFeatures").First(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomsRepository) EditRoom(ctx context.Context, roomID int, newData map[string]interface{}) error {
	return r.DB.WithContext(ctx).Model(&models.Room{}).Where("id = ?", roomID).Updates(newData).Error
}

func (r *RoomsRepository) DeleteRoom(ctx context.Context, roomID int) error {
	return r.DB.WithContext(ctx).Where("id = ?", roomID).Delete(&models.Room{}).Error
}

func (r *RoomsRepository) GetRoomAvailability(ctx context.Context, totalBooked []models.RoomBookedResponse) ([]*models.Room, error) {
	var rooms []*models.Room

	subquery := r.DB.Table("rooms").
		Select("MIN(id) as id").
		Group("room_type_id")

	err := r.DB.WithContext(ctx).
		Joins("JOIN (?) as r2 ON rooms.id = r2.id", subquery).
		Preload("RoomType.RoomPhotos").Preload("RoomType.RoomFeatures").
		Order("rooms.room_type_id ASC").
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	for i, room := range rooms {
		if int(room.RoomType.ID) == totalBooked[i].RoomTypeID && room.RoomType.TotalRooms-totalBooked[i].TotalBooked >= 0 {
			room.Available = true
		} else {
			room.Available = false
		}
	}

	return rooms, nil
}
