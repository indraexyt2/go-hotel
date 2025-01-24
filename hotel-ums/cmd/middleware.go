package cmd

import (
	"github.com/labstack/echo/v4"
	"hotel-ums/helpers"
	"time"
)

func (d *Dependencies) MiddlewareValidateAuthByToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			log = helpers.Logger
		)

		token := e.Request().Header.Get("Authorization")
		if token == "" {
			log.Error("token not found")
			return helpers.SendResponse(e, 401, "unauthorized", nil)
		}

		_, err := d.UserRepo.GetUserSessionByToken(e.Request().Context(), token)
		if err != nil {
			log.Error("user session not found: ", err)
			return helpers.SendResponse(e, 401, "unauthorized", nil)
		}

		tokenWithClaims, err := helpers.ValidateToken(e.Request().Context(), token)
		if err != nil {
			log.Error("invalid token: ", err)
			return helpers.SendResponse(e, 401, "unauthorized", nil)
		}

		if time.Now().Unix() > tokenWithClaims.ExpiresAt.Unix() {
			log.Error("token expired: ", err)
			return helpers.SendResponse(e, 401, "unauthorized", nil)
		}

		e.Set("token", tokenWithClaims)
		return next(e)
	}
}

func (d *Dependencies) MiddlewareValidateAuthByRefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			log = helpers.Logger
		)

		token := e.Request().Header.Get("Authorization")
		if token == "" {
			log.Error("refresh token not found")
			return helpers.SendResponse(e, 401, "unauthorized", nil)
		}

		_, err := d.UserRepo.GetUserSessionByRefreshToken(e.Request().Context(), token)
		if err != nil {
			log.Error("user session not found: ", err)
			return helpers.SendResponse(e, 401, "unauthorized", nil)
		}

		tokenWithClaims, err := helpers.ValidateToken(e.Request().Context(), token)
		if err != nil {
			log.Error("invalid token: ", err)
			return helpers.SendResponse(e, 401, "unauthorized", nil)
		}

		if time.Now().Unix() > tokenWithClaims.ExpiresAt.Unix() {
			log.Error("token expired: ", err)
			return helpers.SendResponse(e, 401, "unauthorized", nil)
		}

		e.Set("token", tokenWithClaims)
		return next(e)
	}
}
