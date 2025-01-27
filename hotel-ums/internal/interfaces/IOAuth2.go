package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"hotel-ums/internal/models"
)

type IOAuth2Service interface {
	GoogleLogin(ctx context.Context, req *models.GoogleUserInfo, token *oauth2.Token) error
}

type IOAuth2API interface {
	Login(e echo.Context) error
	LoginCallback(e echo.Context) error
}
