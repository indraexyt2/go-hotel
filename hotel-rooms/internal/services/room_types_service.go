package services

import (
	"context"
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
