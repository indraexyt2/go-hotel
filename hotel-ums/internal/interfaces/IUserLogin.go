package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-ums/helpers"
	"hotel-ums/internal/models"
)

type IUserLoginService interface {
	Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
	RefreshToken(ctx context.Context, refreshToken string, claimsToken *helpers.Claims) (*models.RefreshTokenResponse, error)
}

type IUserLoginAPI interface {
	Login(e echo.Context) error
	RefreshToken(e echo.Context) error
}
