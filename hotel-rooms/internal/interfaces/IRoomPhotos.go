package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-rooms/internal/models"
)

type IRoomPhotosRepository interface {
	AddRoomTypePhotos(ctx context.Context, roomPhotos []models.RoomPhoto) error
	GetRoomTypePhotos(ctx context.Context, roomTypeID int) ([]models.RoomPhoto, error)
	DeletePhotos(ctx context.Context, photosID int) error
}

type IRoomPhotosService interface {
	AddRoomTypePhotos(ctx context.Context, roomPhotos []models.RoomPhoto) error
	GetRoomTypePhotos(ctx context.Context, roomTypeID int) ([]models.RoomPhoto, error)
	DeletePhotos(ctx context.Context, photosID int) error
}

type IRoomPhotosAPI interface {
	AddRoomTypePhotos(e echo.Context) error
	GetRoomTypePhotos(e echo.Context) error
	DeleteRoomTypePhoto(e echo.Context) error
}
