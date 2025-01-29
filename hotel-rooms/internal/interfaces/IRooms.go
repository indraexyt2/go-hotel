package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-rooms/internal/models"
)

type IRoomsRepository interface {
	AddRoom(ctx context.Context, room *models.Room) error
	GetAllRooms(ctx context.Context) ([]models.Room, error)
	GetRoomDetails(ctx context.Context, id int) (*models.Room, error)
	EditRoom(ctx context.Context, roomID int, newData map[string]interface{}) error
	DeleteRoom(ctx context.Context, roomID int) error
	GetRoomAvailability(ctx context.Context, totalBooked []models.RoomBookedResponse) ([]*models.Room, error)
}

type IRoomsService interface {
	AddRoom(ctx context.Context, room *models.Room) error
	GetAllRooms(ctx context.Context) ([]models.Room, error)
	GetRoomDetails(ctx context.Context, id int) (*models.Room, error)
	EditRoom(ctx context.Context, roomID int, req *models.Room) error
	DeleteRoom(ctx context.Context, roomID int) error
	GetRoomAvailability(ctx context.Context, totalBooked []models.RoomBookedResponse) ([]*models.Room, error)
}

type IRoomsAPI interface {
	AddRoom(e echo.Context) error
	GetRooms(e echo.Context) error
	GetRoomDetails(e echo.Context) error
	EditRoom(e echo.Context) error
	DeleteRoom(e echo.Context) error
	GetRoomAvailability(e echo.Context) error
}
