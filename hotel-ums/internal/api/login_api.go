package api

import (
	"github.com/labstack/echo/v4"
	"hotel-ums/helpers"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/models"
)

type LoginAPI struct {
	LoginService interfaces.IUserLoginService
}

func NewLoginAPI(loginSvc interfaces.IUserLoginService) *LoginAPI {
	return &LoginAPI{LoginService: loginSvc}
}

func (api *LoginAPI) Login(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.LoginRequest{}
	)

	if err := e.Bind(req); err != nil {
		log.Error("Failed to bind user login request: ", err)
		return helpers.SendResponse(e, 400, err.Error(), nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate user login request: ", err)
		return helpers.SendResponse(e, 400, err.Error(), nil)
	}

	resp, err := api.LoginService.Login(e.Request().Context(), req)
	if err != nil {
		log.Error("failed to login: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	return helpers.SendResponse(e, 200, "success", resp)
}

func (api *LoginAPI) RefreshToken(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	token := e.Request().Header.Get("Authorization")
	claims := e.Get("token")
	claimsToken, ok := claims.(*helpers.Claims)
	if !ok {
		log.Error("error getting claims token")
		return helpers.SendResponse(e, 500, "server error", nil)
	}

	resp, err := api.LoginService.RefreshToken(e.Request().Context(), token, claimsToken)
	if err != nil {
		log.Error("failed to update token: ", err)
		return helpers.SendResponse(e, 401, err.Error(), nil)
	}

	return helpers.SendResponse(e, 200, "success", resp)
}
