package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-ums/internal/models"
)

type IUserLoginService interface {
	Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
}

type IUserLoginAPI interface {
	Login(e echo.Context) error
}
