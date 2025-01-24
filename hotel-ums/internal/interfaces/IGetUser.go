package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-ums/internal/models"
)

type IGetUserService interface {
	GetUser(ctx context.Context, id int) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type IGetUserAPI interface {
	GetUser(e echo.Context) error
	GetAllUsers(e echo.Context) error
}
