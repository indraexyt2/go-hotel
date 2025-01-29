package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Booking struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	RoomID       uint      `gorm:"not null;index" json:"room_id" validate:"required"`
	GuestID      uint      `gorm:"not null;index" json:"guest_id" validate:"required"`
	CheckinDate  time.Time `gorm:"not null" json:"checkin_date" validate:"required"`
	CheckoutDate time.Time `gorm:"not null" json:"checkout_date" validate:"required"`
	TotalPrice   float64   `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Status       string    `gorm:"type:booking_status;default:'pending'" json:"status"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (*Booking) TableName() string {
	return "bookings"
}

func (I *Booking) Validate() error {
	v := validator.New()
	return v.Struct(I)
}
