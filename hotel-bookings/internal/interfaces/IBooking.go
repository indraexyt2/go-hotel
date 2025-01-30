package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-bookings/internal/models"
	"time"
)

type IBookingRepository interface {
	AddBooking(ctx context.Context, booking *models.Booking) error
	GetBookingByID(ctx context.Context, bookingID int) (models.Booking, error)
	GetBookingsByUserID(ctx context.Context, userID int) ([]models.Booking, error)
	GetBookings(ctx context.Context) ([]models.Booking, error)
	EditBooking(ctx context.Context, bookingID int, newData map[string]interface{}) error
	UpdateBookingStatus(ctx context.Context, bookingID int, status string) error
	GetTotalBookings(ctx context.Context, checkinDate time.Time, checkoutDate time.Time) ([]models.RoomTypeBooked, error)
}

type IBookingService interface {
	AddBooking(ctx context.Context, booking *models.Booking) error
	GetBookingByID(ctx context.Context, bookingID int) (models.Booking, error)
	GetBookingsByUserID(ctx context.Context, userID int) ([]models.Booking, error)
	GetBookings(ctx context.Context) ([]models.Booking, error)
	EditBooking(ctx context.Context, bookingID int, req *models.Booking) error
	UpdateBookingStatus(ctx context.Context, bookingID int, status string) error
	GetTotalBookings(ctx context.Context, checkinDate time.Time, checkoutDate time.Time) ([]models.RoomTypeBooked, error)
}

type IBookingAPI interface {
	AddBooking(e echo.Context) error
	GetBookingById(e echo.Context) error
	GetBookingByUserId(e echo.Context) error
	GetBookings(e echo.Context) error
	EditBooking(e echo.Context) error
	UpdateBookingStatus(e echo.Context) error
	GetTotalBookings(e echo.Context) error
}
