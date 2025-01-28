package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-rooms/internal/models"
)

type IRoomTypesRepository interface {
	GetAllRoomTypes(ctx context.Context) ([]models.RoomType, error)
	GetRoomTypesDetails(ctx context.Context, id int) (*models.RoomType, error)
}

type IRoomTypesService interface {
	GetAllRoomTypes(ctx context.Context) ([]models.RoomType, error)
	GetRoomTypesDetails(ctx context.Context, id int) (*models.RoomType, error)
}

type IRoomTypesAPI interface {
	GetAllRoomTypes(e echo.Context) error
	GetRoomTypesDetails(e echo.Context) error
}
