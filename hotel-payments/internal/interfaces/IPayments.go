package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-payments/internal/models"
)

type IPaymentsRepository interface {
	CreatePayment(ctx context.Context, req *models.Payment) error
	GetPaymentById(ctx context.Context, bookingID int) (*models.Payment, error)
	UpdatePayment(ctx context.Context, req map[string]interface{}, bookingID string) error
	UpdateStatusTransaction(ctx context.Context, newStatus string, bookingID int) error
	GetPaymentByIdAndUserId(ctx context.Context, bookingID, guestID int) (*models.Payment, error)
}

type IPaymentsService interface {
	CreatePayment(ctx context.Context, req *models.Booking, snapURL string) error
	GetPaymentById(ctx context.Context, bookingID int) (*models.Payment, error)
	GetPaymentByIdAndUserId(ctx context.Context, bookingID, guestID int) (*models.Payment, error)
	UpdatePayment(ctx context.Context, req map[string]interface{}) error
	RefundPayment(ctx context.Context, newStatus string, bookingID int) error
}

type IPaymentsAPI interface {
	ProcessPayment(req *models.Booking) error
	ProcessPaymentCallback(e echo.Context) error
	RefundPayment(e echo.Context) error
}
