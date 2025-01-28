package services

import (
	"context"
	"encoding/json"
	"hotel-rooms/internal/interfaces"
	"hotel-rooms/internal/models"
)

type RoomTypesService struct {
	roomTypesRepo interfaces.IRoomTypesRepository
}

func NewRoomTypesService(roomTypesRepo interfaces.IRoomTypesRepository) *RoomTypesService {
	return &RoomTypesService{roomTypesRepo: roomTypesRepo}
}

func (s *RoomTypesService) GetAllRoomTypes(ctx context.Context) ([]models.RoomType, error) {
	return s.roomTypesRepo.GetAllRoomTypes(ctx)
}

func (s *RoomTypesService) GetRoomTypesDetails(ctx context.Context, id int) (*models.RoomType, error) {
	return s.roomTypesRepo.GetRoomTypesDetails(ctx, id)
}

func (s *RoomTypesService) AddRoomType(ctx context.Context, roomType *models.RoomType) error {
	return s.roomTypesRepo.AddRoomType(ctx, roomType)
}

func (s *RoomTypesService) UpdateRoomType(ctx context.Context, roomType *models.RoomType, id int) error {
	jsonRoomType, err := json.Marshal(roomType)
	if err != nil {
		return err
	}

	newData := make(map[string]interface{})
	err = json.Unmarshal(jsonRoomType, &newData)
	if err != nil {
		return err
	}

	return s.roomTypesRepo.UpdateRoomType(ctx, newData, id)
}
