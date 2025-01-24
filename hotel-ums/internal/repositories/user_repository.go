package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"hotel-ums/constants"
	"hotel-ums/helpers"
	"hotel-ums/internal/models"
	"time"
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

	result, err := r.Redis.Get(ctx, fmt.Sprintf(constants.RedisKeyUserEmail, email)).Result()
	if err == nil && result != "" {
		if err = json.Unmarshal([]byte(result), &user); err != nil {
			return nil, err
		}
		helpers.Logger.Info("success get user from redis")
		return &user, nil
	}

	err = r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	go func() {
		ctx = context.Background()
		userJson, err := json.Marshal(user)
		if err != nil {
			helpers.Logger.Warn("failed to marshal user: ", err)
			return
		}

		err = r.Redis.Set(ctx, fmt.Sprintf(constants.RedisKeyUserEmail, email), string(userJson), time.Hour*24).Err()
		if err != nil {
			helpers.Logger.Warn("failed to set user to redis: ", err)
			return
		}

		helpers.Logger.Info("successfully set user to redis")
	}()

	return &user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	result, err := r.Redis.Get(ctx, fmt.Sprintf(constants.RedisKeyUserUsername, username)).Result()
	if err == nil && result != "" {
		if err = json.Unmarshal([]byte(result), &user); err != nil {
			return nil, err
		}
		helpers.Logger.Info("success get user from redis")
		return &user, nil
	}

	err = r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	go func() {
		ctx = context.Background()
		userJson, err := json.Marshal(user)
		if err != nil {
			helpers.Logger.Warn("failed to marshal user: ", err)
			return
		}

		err = r.Redis.Set(ctx, fmt.Sprintf(constants.RedisKeyUserUsername, username), string(userJson), time.Hour*24).Err()
		if err != nil {
			helpers.Logger.Warn("failed to set user to redis: ", err)
			return
		}

		helpers.Logger.Info("successfully set user to redis")
	}()

	return &user, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id int) (*models.User, error) {
	var user models.User

	result, err := r.Redis.Get(ctx, fmt.Sprintf(constants.RedisKeyUserID, string(rune(id)))).Result()
	if err == nil && result != "" {
		if err = json.Unmarshal([]byte(result), &user); err != nil {
			return nil, err
		}
		helpers.Logger.WithField("user", user).Info("success get user from redis")
		return &user, nil
	}

	err = r.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	go func() {
		ctx = context.Background()
		userJson, err := json.Marshal(user)
		if err != nil {
			helpers.Logger.Warn("failed to marshal user: ", err)
			return
		}

		err = r.Redis.Set(ctx, fmt.Sprintf(constants.RedisKeyUserID, string(rune(id))), string(userJson), time.Hour*24).Err()
		if err != nil {
			helpers.Logger.Warn("failed to set user to redis: ", err)
			return
		}

		helpers.Logger.Info("successfully set user to redis")
	}()

	return &user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	result, err := r.Redis.Get(ctx, constants.RedisKeyAllUsers).Result()
	if err == nil && result != "" {
		if err = json.Unmarshal([]byte(result), &users); err != nil {
			helpers.Logger.Warn("failed to unmarshal data")
			return nil, err
		}

		helpers.Logger.Info("success get data all users from redis")
		return users, nil
	}

	err = r.DB.WithContext(ctx).Where("role != ?", "admin").Omit("password").Find(&users).Error
	if err != nil {
		return nil, err
	}

	go func() {
		ctx = context.Background()
		usersJson, err := json.Marshal(users)
		if err != nil {
			helpers.Logger.Warn("error marshal user data to json")
			return
		}

		err = r.Redis.Set(ctx, constants.RedisKeyAllUsers, string(usersJson), time.Hour*24).Err()
		if err != nil {
			helpers.Logger.Warn("failed to set all user to redis: ", err)
			return
		}

		helpers.Logger.Info("success set data all users to redis")
	}()

	return users, nil
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

	result, err := r.Redis.Get(ctx, fmt.Sprintf(constants.RedisKeyUserSessionByRefreshToken, refreshToken)).Result()
	if err == nil && result != "" {
		if err = json.Unmarshal([]byte(result), &userSession); err != nil {
			helpers.Logger.Warn("failed to unmarshal data")
			return nil, err
		}

		helpers.Logger.Info("success get user session from redis")
		return &userSession, nil
	}

	err = r.DB.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&userSession).Error
	if err != nil {
		return nil, err
	}

	go func() {
		ctx = context.Background()
		userSessionJson, err := json.Marshal(userSession)
		if err != nil {
			helpers.Logger.Warn("error marshal user data to json")
			return
		}

		err = r.Redis.Set(ctx, fmt.Sprintf(constants.RedisKeyUserSessionByRefreshToken, refreshToken), string(userSessionJson), time.Hour*24).Err()
		if err != nil {
			helpers.Logger.Warn("failed to set user session to redis: ", err)
			return
		}

		helpers.Logger.Info("success set user session to redis")
	}()
	return &userSession, nil
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error) {
	var userSession models.UserSession

	result, err := r.Redis.Get(ctx, fmt.Sprintf(constants.RedisKeyUserSessionByToken, token)).Result()
	if err == nil && result != "" {
		if err = json.Unmarshal([]byte(result), &userSession); err != nil {
			helpers.Logger.Warn("failed to unmarshal data")
			return nil, err
		}

		helpers.Logger.Info("success get user session from redis")
		return &userSession, nil
	}

	err = r.DB.WithContext(ctx).Where("token = ?", token).First(&userSession).Error
	if err != nil {
		return nil, err
	}

	go func() {
		ctx = context.Background()
		userSessionJson, err := json.Marshal(userSession)
		if err != nil {
			helpers.Logger.Warn("error marshal user data to json")
			return
		}

		err = r.Redis.Set(ctx, fmt.Sprintf(constants.RedisKeyUserSessionByToken, token), string(userSessionJson), time.Hour*24).Err()
		if err != nil {
			helpers.Logger.Warn("failed to set user session to redis: ", err)
			return
		}

		helpers.Logger.Info("success set user session to redis")
	}()

	return &userSession, nil
}

func (r *UserRepository) UpdateUserSession(ctx context.Context, token, refreshToken string) error {
	return r.DB.WithContext(ctx).Model(&models.UserSession{}).Where("refresh_token = ?", refreshToken).Update("token", token).Error
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string) error {
	return r.DB.WithContext(ctx).Where("token = ?", token).Delete(&models.UserSession{}).Error
}
