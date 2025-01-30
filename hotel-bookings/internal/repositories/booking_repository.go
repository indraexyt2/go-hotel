package repositories

import (
	"context"
	"gorm.io/gorm"
	"hotel-bookings/internal/models"
	"time"
)

type BookingRepository struct {
	DB *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{DB: db}
}

func (r *BookingRepository) AddBooking(ctx context.Context, booking *models.Booking) error {
	return r.DB.WithContext(ctx).Create(booking).Error
}

func (r *BookingRepository) GetBookingByID(ctx context.Context, bookingID int) (models.Booking, error) {
	var booking models.Booking
	err := r.DB.WithContext(ctx).Where("id = ?", bookingID).First(&booking).Error
	return booking, err
}

func (r *BookingRepository) GetBookingsByUserID(ctx context.Context, userID int) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.DB.WithContext(ctx).Where("guest_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func (r *BookingRepository) GetBookings(ctx context.Context) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.DB.WithContext(ctx).Find(&bookings).Error
	return bookings, err
}

func (r *BookingRepository) EditBooking(ctx context.Context, bookingID int, newData map[string]interface{}) error {
	return r.DB.WithContext(ctx).Model(&models.Booking{}).Where("id = ?", bookingID).Updates(newData).Error
}

func (r *BookingRepository) UpdateBookingStatus(ctx context.Context, bookingID int, status string) error {
	return r.DB.WithContext(ctx).Model(&models.Booking{}).Where("id = ?", bookingID).Update("status", status).Error
}

func (r *BookingRepository) GetTotalBookings(ctx context.Context, checkinDate time.Time, checkoutDate time.Time) ([]models.RoomTypeBooked, error) {
	var bookings []models.RoomTypeBooked
	err := r.DB.WithContext(ctx).Table("bookings").
		Select("room_type_id, COUNT(*) as total_booked").
		Where("date(checkin_date) >= ? AND date(checkout_date) <= ?", checkinDate, checkoutDate).
		Group("room_type_id").
		Order("room_type_id ASC").
		Find(&bookings).Error

	return bookings, err
}
