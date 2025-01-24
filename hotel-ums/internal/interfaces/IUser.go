package interfaces

import (
	"context"
	"hotel-ums/internal/models"
)

type IUserRepository interface {
	RegisterNewUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
	GetEmailVerificationToken(ctx context.Context, tokenVerify string) (*models.EmailVerificationToken, error)
	UpdateUser(ctx context.Context, user *models.User) error
	GetEmailVerificationTokenById(ctx context.Context, userID int) (*models.EmailVerificationToken, error)
	UpdateEmailVerificationToken(ctx context.Context, emailVerificationToken *models.EmailVerificationToken) error
}
