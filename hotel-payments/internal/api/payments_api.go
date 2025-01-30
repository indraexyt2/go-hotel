package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	midtrans2 "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/veritrans/go-midtrans"
	"hotel-payments/helpers"
	"hotel-payments/internal/interfaces"
	"hotel-payments/internal/models"
	"net/http"
	"strconv"
)

type PaymentAPI struct {
	PaymentService interfaces.IPaymentsService
	SnapClient     *snap.Client
}

func NewPaymentAPI(paymentSvc interfaces.IPaymentsService, snapClient *snap.Client) *PaymentAPI {
	return &PaymentAPI{
		PaymentService: paymentSvc,
		SnapClient:     snapClient,
	}
}

func (api *PaymentAPI) ProcessPayment(req *models.Booking) error {
	var (
		log = helpers.Logger
	)

	snapReq := &snap.Request{
		TransactionDetails: midtrans2.TransactionDetails(midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(req.ID)),
			GrossAmt: int64(req.TotalPrice),
		}),
		CustomerDetail: &midtrans2.CustomerDetails{},
	}

	snapResp, err := api.SnapClient.CreateTransaction(snapReq)
	if err != nil {
		return err
	}

	err2 := api.PaymentService.CreatePayment(context.Background(), req, snapResp.RedirectURL)
	if err2 != nil {
		return err2
	}

	log.Info(fmt.Printf("Payment URL for Order, Token %s: %s", snapResp.Token, snapResp.RedirectURL))
	return nil
}

func (api *PaymentAPI) ProcessPaymentCallback(e echo.Context) error {
	var (
		log = helpers.Logger
		req = map[string]interface{}{}
	)

	if err := e.Bind(req); err != nil {
		log.Error("Failed to bind request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	err := api.PaymentService.UpdatePayment(e.Request().Context(), req)
	if err != nil {
		log.Error("Failed to update payment: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}
