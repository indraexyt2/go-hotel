package api

import (
	"github.com/labstack/echo/v4"
	"hotel-ums/helpers"
	"hotel-ums/internal/interfaces"
)

type LogoutAPI struct {
	LogoutService interfaces.IUserLogoutService
}

func NewLogoutAPI(logoutService interfaces.IUserLogoutService) *LogoutAPI {
	return &LogoutAPI{LogoutService: logoutService}
}

func (api *LogoutAPI) Logout(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	token := e.Request().Header.Get("Authorization")

	err := api.LogoutService.Logout(e.Request().Context(), token)
	if err != nil {
		log.Error("failed to logout: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	return helpers.SendResponse(e, 200, "success", nil)
}
