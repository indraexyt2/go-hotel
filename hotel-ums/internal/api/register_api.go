package api

import (
	"github.com/labstack/echo/v4"
	"hotel-ums/helpers"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/models"
)

type RegisterAPI struct {
	RegisterService interfaces.IUserRegisterService
}

func NewRegisterAPI(registerService interfaces.IUserRegisterService) *RegisterAPI {
	return &RegisterAPI{
		RegisterService: registerService,
	}
}

func (api *RegisterAPI) RegisterNewUser(e echo.Context) error {
	var (
		log  = helpers.Logger
		user = &models.User{}
	)

	if err := e.Bind(user); err != nil {
		log.Error("Failed to bind user: ", err)
		return helpers.SendResponse(e, 400, err.Error(), nil)
	}

	if err := user.Validate(); err != nil {
		log.Error("Failed to validate user: ", err)
		return helpers.SendResponse(e, 400, err.Error(), nil)
	}

	resp, err := api.RegisterService.RegisterNewUser(e.Request().Context(), user)
	if err != nil {
		log.Error("Failed to register new user: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	return helpers.SendResponse(e, 200, "success", resp)
}

func (api *RegisterAPI) EmailVerification(e echo.Context) error {
	var (
		log              = helpers.Logger
		emailVerifyToken = e.Param("token")
	)

	if emailVerifyToken == "" {
		log.Error("Failed to get token: ", emailVerifyToken)
		return helpers.SendResponse(e, 400, "token is empty", nil)
	}

	err := api.RegisterService.EmailVerification(e.Request().Context(), emailVerifyToken)
	if err != nil {
		log.Error("Failed to verify email: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	return helpers.SendResponse(e, 200, "success", nil)
}

func (api *RegisterAPI) ResendEmailVerification(e echo.Context) error {
	var (
		log = helpers.Logger
		req = &models.ResendEmailVerificationRequest{}
	)

	if err := e.Bind(req); err != nil {
		log.Error("Failed to bind request: ", err)
		return helpers.SendResponse(e, 400, err.Error(), nil)
	}

	err := api.RegisterService.ResendEmailVerification(e.Request().Context(), req.Email)
	if err != nil {
		log.Error("Failed to resend email verification: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	return helpers.SendResponse(e, 200, "success", nil)

}
