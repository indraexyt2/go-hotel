package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
)

type IUserLogoutService interface {
	Logout(ctx context.Context, token string) error
}

type IUserLogoutAPI interface {
	Logout(e echo.Context) error
}
