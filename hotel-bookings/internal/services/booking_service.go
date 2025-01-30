package services

import (
	"context"
	"encoding/json"
	"errors"
	"hotel-bookings/constants"
	"hotel-bookings/internal/interfaces"
	"hotel-bookings/internal/models"
	"os"
	"time"
)

type BookingService struct {
	BookingRepository interfaces.IBookingRepository
	External          interfaces.IExternal
}

func NewBookingService(bookingRepo interfaces.IBookingRepository, ext interfaces.IExternal) *BookingService {
	return &BookingService{
		BookingRepository: bookingRepo,
		External:          ext,
	}
}

func (s *BookingService) AddBooking(ctx context.Context, booking *models.Booking) error {
	err := s.BookingRepository.AddBooking(ctx, booking)
	if err != nil {
		return err
	}

	bookingData, err := json.Marshal(booking)
	if err != nil {
		return err
	}

	err = s.External.ProduceKafkaMessage(ctx, os.Getenv("KAFKA_TOPIC_INITIATE_BOOKING"), bookingData)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookingService) GetBookingByID(ctx context.Context, bookingID int) (models.Booking, error) {
	return s.BookingRepository.GetBookingByID(ctx, bookingID)
}

func (s *BookingService) GetBookingsByUserID(ctx context.Context, userID int) ([]models.Booking, error) {
	return s.BookingRepository.GetBookingsByUserID(ctx, userID)
}

func (s *BookingService) GetBookings(ctx context.Context) ([]models.Booking, error) {
	return s.BookingRepository.GetBookings(ctx)
}

func (s *BookingService) EditBooking(ctx context.Context, bookingID int, req *models.Booking) error {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newData map[string]interface{}
	err = json.Unmarshal(jsonData, &newData)
	if err != nil {
		return err
	}

	return s.BookingRepository.EditBooking(ctx, bookingID, newData)
}

func (s *BookingService) UpdateBookingStatus(ctx context.Context, bookingID int, status string) error {
	bookingOrder, err := s.BookingRepository.GetBookingByID(ctx, bookingID)
	if err != nil {
		return err
	}

	updateStatus := false
	statusMapping := constants.UpdateStatusMapping[bookingOrder.Status]
	for _, v := range statusMapping {
		if v == status {
			updateStatus = true
			break
		}
	}

	if !updateStatus {
		return errors.New("invalid status")
	}

	return s.BookingRepository.UpdateBookingStatus(ctx, bookingID, status)
}

func (s *BookingService) GetTotalBookings(ctx context.Context, checkinDate time.Time, checkoutDate time.Time) ([]models.RoomTypeBooked, error) {
	return s.BookingRepository.GetTotalBookings(ctx, checkinDate, checkoutDate)
}
