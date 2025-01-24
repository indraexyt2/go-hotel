package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-ums/internal/models"
)

type IUserRegisterService interface {
	RegisterNewUser(ctx context.Context, user *models.User) (*models.User, error)
	EmailVerification(ctx context.Context, tokenVerify string) error
}

type IUserRegisterAPI interface {
	RegisterNewUser(e echo.Context) error
	EmailVerification(e echo.Context) error
}
