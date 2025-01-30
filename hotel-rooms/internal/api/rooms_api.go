package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hotel-rooms/helpers"
	"hotel-rooms/internal/interfaces"
	"hotel-rooms/internal/models"
	"net/http"
	"strconv"
)

type RoomsAPI struct {
	RoomsSVC interfaces.IRoomsService
	External interfaces.IExternal
}

func NewRoomsAPI(roomsSvc interfaces.IRoomsService, ext interfaces.IExternal) *RoomsAPI {
	return &RoomsAPI{
		RoomsSVC: roomsSvc,
		External: ext,
	}
}

func (api *RoomsAPI) AddRoom(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.Room{}
	)

	if err := e.Bind(req); err != nil {
		log.Error("failed to bind request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	err := api.RoomsSVC.AddRoom(e.Request().Context(), req)
	if err != nil {
		log.Error("failed to add room: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}

func (api *RoomsAPI) GetRooms(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	resp, err := api.RoomsSVC.GetAllRooms(e.Request().Context())
	if err != nil {
		log.Error("failed to get rooms: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}

func (api *RoomsAPI) GetRoomDetails(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	id := e.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error("failed to convert id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	resp, err := api.RoomsSVC.GetRoomDetails(e.Request().Context(), idInt)
	if err != nil {
		log.Error("failed to get room details: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}

func (api *RoomsAPI) EditRoom(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.Room{}
	)

	id := e.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error("failed to convert id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := e.Bind(req); err != nil {
		log.Error("failed to bind request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	err = api.RoomsSVC.EditRoom(e.Request().Context(), idInt, req)
	if err != nil {
		log.Error("failed to edit room: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}

func (api *RoomsAPI) DeleteRoom(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	id := e.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error("failed to convert id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	err = api.RoomsSVC.DeleteRoom(e.Request().Context(), idInt)
	if err != nil {
		log.Error("failed to delete room: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}

func (api *RoomsAPI) GetRoomAvailability(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	checkinDate := e.QueryParam("checkinDate")
	checkoutDate := e.QueryParam("checkoutDate")

	if checkinDate == "" || checkoutDate == "" {
		log.Error("checkinDate or checkoutDate not found")
		return helpers.SendResponse(e, http.StatusBadRequest, "checkin date or checkout date not found", nil)
	}

	totalBooked, err := api.External.GetTotalBooked(e.Request().Context(), checkinDate, checkoutDate)
	if err != nil {
		log.Error("failed to get total booked: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	if totalBooked == nil {
		fmt.Println("totalBooked is nil")
		return helpers.SendResponse(e, http.StatusOK, "success", nil)
	}

	resp, err := api.RoomsSVC.GetRoomAvailability(e.Request().Context(), totalBooked)
	if err != nil {
		log.Error("failed to get room availability: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}
