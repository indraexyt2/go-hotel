package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Payment struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	TransactionID     string    `gorm:"not null;index" json:"transaction_id,omitempty" validate:"required"`
	BookingID         int       `gorm:"not null" json:"booking_id,omitempty" validate:"required"`
	GuestID           int       `gorm:"not null" json:"guest_id,omitempty" validate:"required"`
	FullName          string    `gorm:"not null" json:"full_name,omitempty" validate:"required"`
	RoomID            int       `gorm:"not null" json:"room_id,omitempty" validate:"required"`
	GrossAmount       float64   `gorm:"type:decimal(10,2)" json:"gross_amount,omitempty"`
	PaymentType       string    `gorm:"type:varchar(50)" json:"payment_type,omitempty"`
	FraudStatus       string    `gorm:"type:varchar(50)" json:"fraud_status,omitempty"`
	TransactionStatus string    `gorm:"type:varchar(50);not null" json:"transaction_status" validate:"required"`
	Currency          string    `gorm:"type:varchar(10)" json:"currency,omitempty"`
	TransactionTime   time.Time `gorm:"not null" json:"transaction_time,omitempty" validate:"required"`
	SnapLink          string    `gorm:"type:text" json:"snap_link,omitempty" validate:"required"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"-"`
}

func (p *Payment) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

type Booking struct {
	ID           uint      `json:"id"`
	RoomTypeID   uint      `json:"room_type_id" validate:"required"`
	RoomID       uint      `json:"room_id" validate:"required"`
	GuestID      uint      `json:"guest_id" validate:"required"`
	FullName     string    `json:"full_name" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	CheckinDate  time.Time `json:"checkin_date"`
	CheckoutDate time.Time `json:"checkout_date"`
	TotalPrice   float64   `json:"total_price"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RefundRequest struct {
	BookingID uint   `json:"booking_id" validate:"required"`
	Reason    string `json:"reason" validate:"required"`
}
