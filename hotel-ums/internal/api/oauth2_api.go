package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"hotel-ums/helpers"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/models"
)

type OAuth2API struct {
	OAuth2Service interfaces.IOAuth2Service
}

func NewOAuth2API(oauth2Service interfaces.IOAuth2Service) *OAuth2API {
	return &OAuth2API{OAuth2Service: oauth2Service}
}

func (api *OAuth2API) Login(e echo.Context) error {
	url := helpers.GoogleOauth2Config.AuthCodeURL(helpers.State, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return e.Redirect(302, url)
}

func (api *OAuth2API) LoginCallback(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	state := e.QueryParam("state")
	if state != helpers.State {
		log.Info("invalid state")
		return helpers.SendResponse(e, 400, "invalid state", nil)
	}

	code := e.QueryParam("code")
	if code == "" {
		log.Error("code not found")
		return helpers.SendResponse(e, 400, "code not found", nil)
	}

	token, err := helpers.GoogleOauth2Config.Exchange(e.Request().Context(), code)
	if err != nil {
		log.Error("failed to exchange token: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	client := helpers.GoogleOauth2Config.Client(e.Request().Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Error("failed to get user info: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}
	defer resp.Body.Close()

	userRequest := &models.GoogleUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(userRequest)
	if err != nil {
		log.Error("failed to decode user info: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	err = api.OAuth2Service.GoogleLogin(e.Request().Context(), userRequest, token)
	if err != nil {
		log.Error("failed to login: ", err)
		return helpers.SendResponse(e, 500, err.Error(), nil)
	}

	return helpers.SendResponse(e, 200, "success", userRequest)
}
