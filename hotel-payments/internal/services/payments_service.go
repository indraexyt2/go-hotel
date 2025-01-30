package services

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"hotel-payments/internal/interfaces"
	"hotel-payments/internal/models"
	"math/rand"
	"strconv"
	"time"
)

type PaymentService struct {
	PaymentRepository interfaces.IPaymentsRepository
}

func NewPaymentService(paymentRepo interfaces.IPaymentsRepository) *PaymentService {
	return &PaymentService{PaymentRepository: paymentRepo}
}

func (s *PaymentService) CreatePayment(ctx context.Context, req *models.Booking, snapURL string) error {
	timeNow := int(time.Now().Unix()) + rand.Intn(100)
	transactionID := strconv.Itoa(timeNow)

	newPayment := &models.Payment{
		TransactionID: transactionID,
		BookingID:     int(req.ID),
		GuestID:       int(req.GuestID),
		FullName:      req.FullName,
		RoomID:        int(req.RoomID),
		GrossAmount:   req.TotalPrice,
		SnapLink:      snapURL,
	}
	return s.PaymentRepository.CreatePayment(ctx, newPayment)
}

func (s *PaymentService) GetPaymentById(ctx context.Context, bookingID int) (*models.Payment, error) {
	return s.PaymentRepository.GetPaymentById(ctx, bookingID)
}

func (s *PaymentService) UpdatePayment(ctx context.Context, req map[string]interface{}) error {
	orderIdStr, exists := req["order_id"].(string)
	if !exists {
		log.Error("Failed to get order id")
		return errors.New("invalid request")
	}

	transactionTime, _ := req["transaction_time"].(string)
	paymentType, _ := req["payment_type"].(string)
	currency, _ := req["currency"].(string)
	transactionStatus, _ := req["transaction_status"].(string)
	fraudStatus, _ := req["fraud_status"].(string)
	orderId, _ := strconv.Atoi(orderIdStr)

	payment, err := s.GetPaymentById(ctx, orderId)
	if err != nil {
		log.Error("Failed to get payment by order id: ", err)
		return err
	} else {
		if payment != nil {
			if transactionStatus == "capture" {
				if fraudStatus == "challenge" {
					transactionStatus = "challenge"
				} else if fraudStatus == "accept" {
					transactionStatus = "success"
				}
			} else if transactionStatus == "settlement" {
				transactionStatus = "success"
			} else if transactionStatus == "deny" {

			} else if transactionStatus == "cancel" || transactionStatus == "expire" {
				transactionStatus = "failure"
			} else if transactionStatus == "pending" {
				transactionStatus = "pending"
			}
		}
	}

	updateData := map[string]interface{}{
		"transaction_time":   transactionTime,
		"payment_type":       paymentType,
		"currency":           currency,
		"transaction_status": transactionStatus,
		"fraud_status":       fraudStatus,
	}

	return s.PaymentRepository.UpdatePayment(ctx, updateData, orderIdStr)
}
