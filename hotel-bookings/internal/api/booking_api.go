package api

import (
	"github.com/labstack/echo/v4"
	"hotel-bookings/external"
	"hotel-bookings/helpers"
	"hotel-bookings/internal/interfaces"
	"hotel-bookings/internal/models"
	"net/http"
	"strconv"
	"time"
)

type BookingAPI struct {
	BookingSVC interfaces.IBookingService
}

func NewBookingAPI(bookingSVC interfaces.IBookingService) *BookingAPI {
	return &BookingAPI{BookingSVC: bookingSVC}
}

func (api *BookingAPI) AddBooking(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.BookingRequest{}
	)

	claims, ok := e.Get("token").(*external.User)
	if !ok {
		log.Error("failed to get token")
		return helpers.SendResponse(e, http.StatusUnauthorized, "unauthorized", nil)
	}

	if err := e.Bind(req); err != nil {
		log.Error("failed to binding request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, "invalid request", nil)
	}

	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Jakarta")
	checkinDateParsed, err := time.ParseInLocation(layout, req.CheckinDate, loc)
	if err != nil {
		log.Error("failed to parse checkin date: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	checkoutDateParsed, err := time.ParseInLocation(layout, req.CheckoutDate, loc)
	if err != nil {
		log.Error("failed to parse checkout date: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	newBooking := &models.Booking{
		RoomTypeID:   req.RoomTypeID,
		RoomID:       req.RoomID,
		GuestID:      claims.ID,
		CheckinDate:  checkinDateParsed,
		CheckoutDate: checkoutDateParsed,
		TotalPrice:   req.TotalPrice,
	}

	if err := api.BookingSVC.AddBooking(e.Request().Context(), newBooking); err != nil {
		log.Error("failed to add booking: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}

func (api *BookingAPI) GetBookingById(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	id := e.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error("failed to convert id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	resp, err := api.BookingSVC.GetBookingByID(e.Request().Context(), idInt)
	if err != nil {
		log.Error("failed to get booking by id: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}

func (api *BookingAPI) GetBookingByUserId(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	id := e.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error("failed to convert id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	resp, err := api.BookingSVC.GetBookingsByUserID(e.Request().Context(), idInt)
	if err != nil {
		log.Error("failed to get booking by user id: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}

func (api *BookingAPI) GetBookings(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	resp, err := api.BookingSVC.GetBookings(e.Request().Context())
	if err != nil {
		log.Error("failed to get all bookings: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}

func (api *BookingAPI) EditBooking(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.BookingRequest{}
	)

	claims, ok := e.Get("token").(*external.User)
	if !ok {
		log.Error("failed to get token")
		return helpers.SendResponse(e, http.StatusUnauthorized, "unauthorized", nil)
	}

	id := e.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error("failed to convert id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := e.Bind(req); err != nil {
		log.Error("failed to binding request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Jakarta")
	checkinDateParsed, err := time.ParseInLocation(layout, req.CheckinDate, loc)
	if err != nil {
		log.Error("failed to parse checkin date: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	checkoutDateParsed, err := time.ParseInLocation(layout, req.CheckoutDate, loc)
	if err != nil {
		log.Error("failed to parse checkout date: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	bookingData := &models.Booking{
		ID:           uint(idInt),
		GuestID:      claims.ID,
		RoomTypeID:   req.RoomTypeID,
		RoomID:       req.RoomID,
		CheckinDate:  checkinDateParsed,
		CheckoutDate: checkoutDateParsed,
		TotalPrice:   req.TotalPrice,
		Status:       req.Status,
	}

	if err := api.BookingSVC.EditBooking(e.Request().Context(), idInt, bookingData); err != nil {
		log.Error("failed to edit booking: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}

func (api *BookingAPI) UpdateBookingStatus(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.UpdateBookingStatusRequest{}
	)

	id := e.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error("failed to convert id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := e.Bind(req); err != nil {
		log.Error("failed to binding request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := api.BookingSVC.UpdateBookingStatus(e.Request().Context(), idInt, req.Status); err != nil {
		log.Error("failed to update booking status: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}

func (api *BookingAPI) GetTotalBookings(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	checkinDate := e.QueryParam("checkinDate")
	checkoutDate := e.QueryParam("checkoutDate")

	if checkinDate == "" || checkoutDate == "" {
		log.Error("checkinDate or checkoutDate not found")
		return helpers.SendResponse(e, http.StatusBadRequest, "checkin date or checkout date not found", nil)
	}

	layout := "2006-01-02"
	checkinDateParsed, err := time.Parse(layout, checkinDate)
	if err != nil {
		log.Error("failed to parse checkin date: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	checkoutDateParsed, err := time.Parse(layout, checkoutDate)
	if err != nil {
		log.Error("failed to parse checkout date: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	resp, err := api.BookingSVC.GetTotalBookings(e.Request().Context(), checkinDateParsed, checkoutDateParsed)
	if err != nil {
		log.Error("failed to get total bookings: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}
