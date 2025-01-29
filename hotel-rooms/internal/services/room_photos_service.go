package services

import (
	"context"
	"hotel-rooms/internal/interfaces"
	"hotel-rooms/internal/models"
)

type RoomPhotosService struct {
	RoomPhotosRepo interfaces.IRoomPhotosRepository
}

func NewRoomPhotosService(roomPhotosRepo interfaces.IRoomPhotosRepository) *RoomPhotosService {
	return &RoomPhotosService{RoomPhotosRepo: roomPhotosRepo}
}

func (r *RoomPhotosService) AddRoomTypePhotos(ctx context.Context, roomPhotos []models.RoomPhoto) error {
	return r.RoomPhotosRepo.AddRoomTypePhotos(ctx, roomPhotos)
}

func (r *RoomPhotosService) GetRoomTypePhotos(ctx context.Context, roomTypeID int) ([]models.RoomPhoto, error) {
	resp, err := r.RoomPhotosRepo.GetRoomTypePhotos(ctx, roomTypeID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RoomPhotosService) DeletePhotos(ctx context.Context, photosID int) error {
	return r.RoomPhotosRepo.DeletePhotos(ctx, photosID)
}
