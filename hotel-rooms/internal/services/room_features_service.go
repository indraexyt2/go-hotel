package services

import (
	"context"
	"encoding/json"
	"fmt"
	"hotel-rooms/internal/interfaces"
	"hotel-rooms/internal/models"
)

type RoomFeaturesService struct {
	roomFeaturesRepo interfaces.IRoomFeaturesRepository
}

func NewRoomFeaturesService(roomFeaturesRepo interfaces.IRoomFeaturesRepository) *RoomFeaturesService {
	return &RoomFeaturesService{roomFeaturesRepo: roomFeaturesRepo}
}

func (s *RoomFeaturesService) AddRoomFeature(ctx context.Context, roomFeature *models.RoomFeature) error {
	return s.roomFeaturesRepo.AddRoomFeature(ctx, roomFeature)
}

func (s *RoomFeaturesService) GetAllRoomFeatures(ctx context.Context, roomTypeID int) ([]models.RoomFeature, error) {
	return s.roomFeaturesRepo.GetAllRoomFeatures(ctx, roomTypeID)
}

func (s *RoomFeaturesService) EditRoomFeature(ctx context.Context, roomFeatureID int, req *models.RoomFeature) error {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	fmt.Println("req: ", req)

	var newData map[string]interface{}
	err = json.Unmarshal(jsonReq, &newData)
	if err != nil {
		return err
	}

	return s.roomFeaturesRepo.EditRoomFeature(ctx, roomFeatureID, newData)
}

func (s *RoomFeaturesService) DeleteRoomFeature(ctx context.Context, roomFeatureID int) error {
	return s.roomFeaturesRepo.DeleteRoomFeature(ctx, roomFeatureID)
}
