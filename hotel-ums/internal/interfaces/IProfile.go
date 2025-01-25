package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"hotel-ums/internal/models"
)

type IProfileService interface {
	UpdateUserProfile(ctx context.Context, user *models.User, photoPath string, userID int) (*models.User, error)
}

type IProfileAPI interface {
	UpdateUserProfile(e echo.Context) error
}
