package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Booking struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	RoomTypeID   uint      `gorm:"not null;index" json:"room_type_id" validate:"required"`
	RoomID       uint      `gorm:"not null;index" json:"room_id" validate:"required"`
	GuestID      uint      `gorm:"not null;index" json:"guest_id" validate:"required"`
	FullName     string    `gorm:"not null" json:"full_name" validate:"required"`
	CheckinDate  time.Time `gorm:"not null" json:"checkin_date"`
	CheckoutDate time.Time `gorm:"not null" json:"checkout_date"`
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

type UpdateBookingStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

func (I *UpdateBookingStatusRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

type RoomTypeBooked struct {
	RoomTypeID  int `json:"room_type_id"`
	TotalBooked int `json:"total_booked"`
}

// BookingRequest
type BookingRequest struct {
	RoomTypeID   uint    `json:"room_type_id" validate:"required"`
	RoomID       uint    `json:"room_id" validate:"required"`
	CheckinDate  string  `json:"checkin_date" validate:"required"`
	CheckoutDate string  `json:"checkout_date" validate:"required"`
	TotalPrice   float64 `json:"total_price"`
	Status       string  `json:"status,omitempty"`
}

func (I *BookingRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}
