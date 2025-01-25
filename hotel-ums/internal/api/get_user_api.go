package api

import (
	"github.com/labstack/echo/v4"
	"hotel-ums/helpers"
	"hotel-ums/internal/interfaces"
)

type GetUserAPI struct {
	GetUserService interfaces.IGetUserService
}

func NewGetUserAPI(getUserSvc interfaces.IGetUserService) *GetUserAPI {
	return &GetUserAPI{GetUserService: getUserSvc}
}

func (api *GetUserAPI) GetUser(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	token := e.Get("token")
	claimsToken, ok := token.(*helpers.Claims)
	if !ok {
		log.Error("error getting token")
		return helpers.SendResponse(e, 500, "server error", nil)
	}

	resp, err := api.GetUserService.GetUser(e.Request().Context(), claimsToken.UserID)
	if err != nil {
		log.Error("failed to get user: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	return helpers.SendResponse(e, 200, "success", resp)
}

func (api *GetUserAPI) GetAllUsers(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	resp, err := api.GetUserService.GetAllUsers(e.Request().Context())
	if err != nil {
		log.Error("failed to get all user: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	return helpers.SendResponse(e, 200, "success", resp)
}
