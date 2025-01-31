package api

import (
	"context"
	"github.com/labstack/echo/v4"
	midtrans2 "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/veritrans/go-midtrans"
	"hotel-payments/external"
	"hotel-payments/helpers"
	"hotel-payments/internal/interfaces"
	"hotel-payments/internal/models"
	"net/http"
	"strconv"
)

type PaymentAPI struct {
	PaymentService interfaces.IPaymentsService
	SnapClient     *snap.Client
	CoreClient     *midtrans.CoreGateway
	External       interfaces.IExternal
}

func NewPaymentAPI(paymentSvc interfaces.IPaymentsService, snapClient *snap.Client, coreClient *midtrans.CoreGateway, ex interfaces.IExternal) *PaymentAPI {
	return &PaymentAPI{
		PaymentService: paymentSvc,
		SnapClient:     snapClient,
		CoreClient:     coreClient,
		External:       ex,
	}
}

func (api *PaymentAPI) ProcessPayment(req *models.Booking) error {
	snapReq := &snap.Request{
		TransactionDetails: midtrans2.TransactionDetails(midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(req.ID)),
			GrossAmt: int64(req.TotalPrice),
		}),
		CustomerDetail: &midtrans2.CustomerDetails{
			FName: req.FullName,
			Email: req.Email,
		},
	}

	snapResp, err := api.SnapClient.CreateTransaction(snapReq)
	if err != nil {
		return err
	}

	err2 := api.PaymentService.CreatePayment(context.Background(), req, snapResp.RedirectURL)
	if err2 != nil {
		return err2
	}

	return nil
}

func (api *PaymentAPI) ProcessPaymentCallback(e echo.Context) error {
	var (
		log = helpers.Logger
		req = map[string]interface{}{}
	)

	if err := e.Bind(&req); err != nil {
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

func (api *PaymentAPI) RefundPayment(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.RefundRequest{}
	)

	token := e.Get("token").(*external.User)
	if token == nil {
		log.Error("token not found")
		return helpers.SendResponse(e, http.StatusUnauthorized, "unauthorized", nil)
	}

	if err := e.Bind(&req); err != nil {
		log.Error("Failed to bind request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	paymentData, err := api.External.GetMidtransTransactionData(e.Request().Context(), strconv.Itoa(int(req.BookingID)))
	if err != nil {
		log.Error("Failed to get midtrans transaction data: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	paymentUser, err := api.PaymentService.GetPaymentByIdAndUserId(e.Request().Context(), int(req.BookingID), int(token.ID))
	if err != nil {
		log.Error("Failed to get transaction data in database: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	if paymentUser == nil {
		log.Error("Failed to get transaction data in database: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, "transaction not found", nil)
	}

	if paymentData == nil {
		log.Error("Failed to get midtrans transaction data: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, "transaction not found", nil)
	}

	if paymentData.TransactionStatus != "settlement" {
		log.Error("Failed to get midtrans transaction data: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	amountFloat, _ := strconv.ParseFloat(paymentData.GrossAmount, 64)
	amountInt := int64(amountFloat)
	refundReq := midtrans.RefundReq{
		RefundKey: paymentData.TransactionID,
		Amount:    amountInt,
		Reason:    req.Reason,
	}

	refundResp, err := api.CoreClient.Refund(strconv.Itoa(int(req.BookingID)), &refundReq)
	if err != nil {
		log.Error("Failed to refund payment: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	err = api.PaymentService.RefundPayment(e.Request().Context(), refundResp.TransactionStatus, int(req.BookingID))
	if err != nil {
		log.Error("Failed to refund payment: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}
