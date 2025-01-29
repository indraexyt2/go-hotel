package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-rooms/internal/models"
)

type IRoomFeaturesRepository interface {
	AddRoomFeature(ctx context.Context, roomFeature *models.RoomFeature) error
	GetAllRoomFeatures(ctx context.Context, roomTypeID int) ([]models.RoomFeature, error)
	EditRoomFeature(ctx context.Context, roomFeatureID int, newData map[string]interface{}) error
	DeleteRoomFeature(ctx context.Context, roomFeatureID int) error
}

type IRoomFeaturesService interface {
	AddRoomFeature(ctx context.Context, roomFeature *models.RoomFeature) error
	GetAllRoomFeatures(ctx context.Context, roomTypeID int) ([]models.RoomFeature, error)
	EditRoomFeature(ctx context.Context, roomFeatureID int, req *models.RoomFeature) error
	DeleteRoomFeature(ctx context.Context, roomFeatureID int) error
}

type IRoomFeaturesAPI interface {
	AddRoomFeature(e echo.Context) error
	GetAllRoomFeatures(e echo.Context) error
	EditRoomFeature(e echo.Context) error
	DeleteRoomFeature(e echo.Context) error
}
