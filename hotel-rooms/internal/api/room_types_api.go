package api

import (
	"github.com/labstack/echo/v4"
	"hotel-rooms/helpers"
	"hotel-rooms/internal/interfaces"
	"hotel-rooms/internal/models"
	"net/http"
	"strconv"
)

type RoomTypesAPI struct {
	RoomTypesSVC interfaces.IRoomTypesService
}

func NewRoomTypesAPI(roomTypesSvc interfaces.IRoomTypesService) *RoomTypesAPI {
	return &RoomTypesAPI{RoomTypesSVC: roomTypesSvc}
}

func (api *RoomTypesAPI) GetAllRoomTypes(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	resp, err := api.RoomTypesSVC.GetAllRoomTypes(e.Request().Context())
	if err != nil {
		log.Error("failed to get all room types: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}

func (api *RoomTypesAPI) GetRoomTypesDetails(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	id := e.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error("failed to convert id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	resp, err := api.RoomTypesSVC.GetRoomTypesDetails(e.Request().Context(), idInt)
	if err != nil {
		log.Error("failed to get room types details: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}

func (api *RoomTypesAPI) AddRoomType(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.RoomType{}
	)

	if err := e.Bind(req); err != nil {
		log.Error("failed to bind request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	err := api.RoomTypesSVC.AddRoomType(e.Request().Context(), req)
	if err != nil {
		log.Error("failed to add room type: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}
