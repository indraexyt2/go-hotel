package services

import (
	"context"
	"encoding/json"
	"hotel-rooms/internal/interfaces"
	"hotel-rooms/internal/models"
)

type RoomService struct {
	RoomsRepository interfaces.IRoomsRepository
}

func NewRoomService(roomRepo interfaces.IRoomsRepository) *RoomService {
	return &RoomService{RoomsRepository: roomRepo}
}

func (s *RoomService) AddRoom(ctx context.Context, room *models.Room) error {
	return s.RoomsRepository.AddRoom(ctx, room)
}

func (s *RoomService) GetAllRooms(ctx context.Context) ([]models.Room, error) {
	return s.RoomsRepository.GetAllRooms(ctx)
}

func (s *RoomService) GetRoomDetails(ctx context.Context, id int) (*models.Room, error) {
	return s.RoomsRepository.GetRoomDetails(ctx, id)
}

func (s *RoomService) EditRoom(ctx context.Context, roomID int, req *models.Room) error {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newData map[string]interface{}
	err = json.Unmarshal(jsonReq, &newData)
	if err != nil {
		return err
	}
	return s.RoomsRepository.EditRoom(ctx, roomID, newData)
}

func (s *RoomService) DeleteRoom(ctx context.Context, roomID int) error {
	return s.RoomsRepository.DeleteRoom(ctx, roomID)
}

func (s *RoomService) GetRoomAvailability(ctx context.Context, totalBooked []models.RoomBookedResponse) ([]*models.Room, error) {
	resp, err := s.RoomsRepository.GetRoomAvailability(ctx, totalBooked)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
