package interfaces

import (
	"context"
	"hotel-rooms/external"
	"hotel-rooms/internal/models"
)

type IExternal interface {
	GetTotalBooked(ctx context.Context, checkinDate string, checkoutDate string) ([]models.RoomBookedResponse, error)
	ValidateUser(ctx context.Context, token string) (*external.User, error)
}
