package repositories

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"hotel-ums/internal/models"
)

type UserRepository struct {
	DB    *gorm.DB
	Redis *redis.ClusterClient
}

func NewUserRepository(db *gorm.DB, redis *redis.ClusterClient) *UserRepository {
	return &UserRepository{
		DB:    db,
		Redis: redis,
	}
}

func (r *UserRepository) RegisterNewUser(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetEmailVerificationToken(ctx context.Context, tokenVerify string) (*models.EmailVerificationToken, error) {
	var emailVerificationToken *models.EmailVerificationToken
	err := r.DB.WithContext(ctx).Where("token = ?", tokenVerify).First(&emailVerificationToken).Error
	if err != nil {
		return nil, err
	}
	return emailVerificationToken, nil
}

func (r *UserRepository) GetEmailVerificationTokenById(ctx context.Context, userID int) (*models.EmailVerificationToken, error) {
	var emailVerificationToken *models.EmailVerificationToken
	err := r.DB.WithContext(ctx).Where("user_id = ?", userID).First(&emailVerificationToken).Error
	if err != nil {
		return nil, err
	}
	return emailVerificationToken, nil
}

func (r *UserRepository) UpdateEmailVerificationToken(ctx context.Context, emailVerificationToken *models.EmailVerificationToken) error {
	return r.DB.WithContext(ctx).Save(emailVerificationToken).Error
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) AddUserSession(ctx context.Context, userSession *models.UserSession) error {
	return r.DB.WithContext(ctx).Create(userSession).Error
}

func (r *UserRepository) GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (*models.UserSession, error) {
	var userSession models.UserSession
	err := r.DB.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&userSession).Error
	if err != nil {
		return nil, err
	}
	return &userSession, nil
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error) {
	var userSession models.UserSession
	err := r.DB.WithContext(ctx).Where("token = ?", token).First(&userSession).Error
	if err != nil {
		return nil, err
	}
	return &userSession, nil
}

func (r *UserRepository) UpdateUserSession(ctx context.Context, token, refreshToken string) error {
	return r.DB.WithContext(ctx).Model(&models.UserSession{}).Where("refresh_token = ?", refreshToken).Update("token", token).Error
}
