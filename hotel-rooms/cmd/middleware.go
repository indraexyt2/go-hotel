package cmd

import (
	"github.com/labstack/echo/v4"
	"hotel-rooms/helpers"
	"net/http"
)

func (d *Dependencies) MiddlewareAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			log = helpers.Logger
		)

		token := e.Request().Header.Get("Authorization")
		if token == "" {
			log.Error("token not found")
			return helpers.SendResponse(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		userData, err := d.External.ValidateUser(e.Request().Context(), token)
		if err != nil {
			log.Error("failed to validate user: ", err)
			return helpers.SendResponse(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		e.Set("token", userData)
		return next(e)
	}
}

func (d *Dependencies) MiddlewareAdminAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			log = helpers.Logger
		)

		token := e.Get("Authorization")
		if token == "" {
			log.Error("token not found")
			return helpers.SendResponse(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		userData, err := d.External.ValidateUser(e.Request().Context(), token.(string))
		if err != nil {
			log.Error("failed to validate user: ", err)
			return helpers.SendResponse(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		if userData.Role != "admin" {
			log.Error("only admin can access!")
			return helpers.SendResponse(e, http.StatusUnauthorized, "unauthorized", nil)
		}

		e.Set("token", userData)
		return next(e)
	}
}
