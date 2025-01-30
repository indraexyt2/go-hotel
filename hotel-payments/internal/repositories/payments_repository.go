package repositories

import (
	"context"
	"gorm.io/gorm"
	"hotel-payments/internal/models"
)

type PaymentsRepository struct {
	DB *gorm.DB
}

func NewPaymentsRepository(db *gorm.DB) *PaymentsRepository {
	return &PaymentsRepository{
		DB: db,
	}
}

func (r *PaymentsRepository) CreatePayment(ctx context.Context, req *models.Payment) error {
	return r.DB.WithContext(ctx).Create(req).Error
}

func (r *PaymentsRepository) GetPaymentById(ctx context.Context, bookingID int) (*models.Payment, error) {
	var payment models.Payment
	err := r.DB.WithContext(ctx).Where("booking_id = ?", bookingID).First(&payment).Error
	return &payment, err
}

func (r *PaymentsRepository) UpdatePayment(ctx context.Context, req map[string]interface{}) error {
	return r.DB.WithContext(ctx).Model(&models.Payment{}).Where("booking_id = ?", req["booking_id"]).Updates(req).Error
}
