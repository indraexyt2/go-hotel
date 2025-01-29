package api

import (
	"github.com/labstack/echo/v4"
	"hotel-rooms/helpers"
	"hotel-rooms/internal/interfaces"
	"hotel-rooms/internal/models"
	"net/http"
	"strconv"
)

type RoomFeaturesAPI struct {
	RoomFeaturesSVC interfaces.IRoomFeaturesService
}

func NewRoomFeaturesAPI(roomFeaturesSvc interfaces.IRoomFeaturesService) *RoomFeaturesAPI {
	return &RoomFeaturesAPI{RoomFeaturesSVC: roomFeaturesSvc}
}

func (api *RoomFeaturesAPI) AddRoomFeature(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.RoomFeature{}
	)

	if err := e.Bind(req); err != nil {
		log.Error("failed to bind request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	err := api.RoomFeaturesSVC.AddRoomFeature(e.Request().Context(), req)
	if err != nil {
		log.Error("failed to add room feature: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}

func (api *RoomFeaturesAPI) GetAllRoomFeatures(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	IdStr := e.Param("roomTypeId")
	IdInt, err := strconv.Atoi(IdStr)
	if err != nil {
		log.Info("failed to change id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	resp, err := api.RoomFeaturesSVC.GetAllRoomFeatures(e.Request().Context(), IdInt)
	if err != nil {
		log.Error("failed to get all room features: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", resp)
}

func (api *RoomFeaturesAPI) EditRoomFeature(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.RoomFeature{}
	)

	roomFeatureIdStr := e.Param("id")
	roomFeatureIdInt, err := strconv.Atoi(roomFeatureIdStr)
	if err != nil {
		log.Info("failed to change id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	if err := e.Bind(req); err != nil {
		log.Error("failed to bind request: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	err = api.RoomFeaturesSVC.EditRoomFeature(e.Request().Context(), roomFeatureIdInt, req)
	if err != nil {
		log.Error("failed to edit room feature: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}

func (api *RoomFeaturesAPI) DeleteRoomFeature(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	roomFeatureIdStr := e.Param("id")
	roomFeatureIdInt, err := strconv.Atoi(roomFeatureIdStr)
	if err != nil {
		log.Info("failed to change id to int: ", err)
		return helpers.SendResponse(e, http.StatusBadRequest, err.Error(), nil)
	}

	err = api.RoomFeaturesSVC.DeleteRoomFeature(e.Request().Context(), roomFeatureIdInt)
	if err != nil {
		log.Error("failed to delete room feature: ", err)
		return helpers.SendResponse(e, http.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(e, http.StatusOK, "success", nil)
}
