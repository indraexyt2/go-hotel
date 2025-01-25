package interfaces

import (
	"context"
	"hotel-ums/internal/models"
)

type IUserRepository interface {
	RegisterNewUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetEmailVerificationToken(ctx context.Context, tokenVerify string) (*models.EmailVerificationToken, error)
	UpdateUser(ctx context.Context, user *models.User) error
	GetEmailVerificationTokenById(ctx context.Context, userID int) (*models.EmailVerificationToken, error)
	UpdateEmailVerificationToken(ctx context.Context, emailVerificationToken *models.EmailVerificationToken) error

	AddUserSession(ctx context.Context, userSession *models.UserSession) error
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (*models.UserSession, error)
	GetUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error)
	UpdateUserSession(ctx context.Context, token, refreshToken string) error

	DeleteUserSession(ctx context.Context, token string) error
}
