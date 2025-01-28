package api

import (
	"github.com/labstack/echo/v4"
	"hotel-rooms/helpers"
	"hotel-rooms/internal/interfaces"
	"net/http"
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
