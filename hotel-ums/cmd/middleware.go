package cmd

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
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

		session, err := d.UserRepo.GetUserSessionByToken(e.Request().Context(), token)
		if err != nil {
			log.Error("user session not found: ", err)
			return helpers.SendResponse(e, 401, "unauthorized", nil)
		}

		if session.Source == "internal" {
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
		} else if session.Source == "google" {
			tokenGoogle := oauth2.Token{
				AccessToken:  session.Token,
				RefreshToken: session.RefreshToken,
				Expiry:       session.ExpiresAt,
			}

			fmt.Println("expired: ", session.ExpiresAt)
			if session.ExpiresAt.Unix() < time.Now().Unix() {
				fmt.Println("kesini")
				if session.RefreshToken != "" {
					accessToken, err := helpers.RefreshAccessToken(e.Request().Context(), session.RefreshToken)
					if err != nil {
						log.Error("failed to refresh access token: ", err)
						return helpers.SendResponse(e, 401, "unauthorized", nil)
					}

					tokenGoogle.AccessToken = accessToken.AccessToken
					tokenGoogle.Expiry = accessToken.Expiry
					tokenGoogle.RefreshToken = accessToken.RefreshToken

					err = d.UserRepo.UpdateUserSession(e.Request().Context(), tokenGoogle.AccessToken, session.RefreshToken, tokenGoogle.Expiry)
					if err != nil {
						log.Error("failed to update user session: ", err)
						return helpers.SendResponse(e, 401, "unauthorized", nil)
					}
				} else {
					log.Error("refresh token not found")
					return helpers.SendResponse(e, 401, "unauthorized", nil)
				}
			}

			response, err := helpers.ValidateTokenGoogle(e.Request().Context(), &tokenGoogle)
			if err != nil {
				log.Error("invalid token: ", err)
				return helpers.SendResponse(e, 401, "unauthorized", nil)
			}

			e.Set("token", response)
			return next(e)
		} else {
			return helpers.SendResponse(e, 401, "unauthorized", nil)
		}
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

func (d *Dependencies) MiddlewareValidateAdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			log = helpers.Logger
		)

		token := e.Request().Header.Get("Authorization")
		if token == "" {
			log.Error("refresh token not found")
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

		if tokenWithClaims.Role != "admin" {
			log.Error("only admin can access!")
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
