package services

import (
	"context"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/models"
	"math/rand"
	"time"
)

type RegisterService struct {
	UserRepository interfaces.IUserRepository
}

func NewRegisterService(userRepository interfaces.IUserRepository) *RegisterService {
	return &RegisterService{
		UserRepository: userRepository,
	}
}

func (s *RegisterService) RegisterNewUser(ctx context.Context, user *models.User) (*models.User, error) {
	var resp *models.User

	userData, _ := s.UserRepository.GetUserByEmail(ctx, user.Email)
	if userData != nil {
		if userData.Email == user.Email && userData.Username == user.Username {
			return nil, errors.New("email or username already exist")
		}
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	tokenByte := make([]byte, 32)
	_, err = rand.Read(tokenByte)
	if err != nil {
		return nil, err
	}

	if user.Role == "admin" {
		user.Role = "admin"
	}

	user.Password = string(hashedPass)
	user.IsVerified = false
	user.EmailVerificationToken.Token = base64.URLEncoding.EncodeToString(tokenByte)
	user.EmailVerificationToken.ExpiresAt = time.Now().Add(time.Hour * 24)

	err = s.UserRepository.RegisterNewUser(ctx, user)
	if err != nil {
		return nil, err
	}

	resp = user
	resp.Password = ""

	return resp, nil
}

func (s *RegisterService) EmailVerification(ctx context.Context, tokenVerify string) error {
	emailVerificationData, err := s.UserRepository.GetEmailVerificationToken(ctx, tokenVerify)
	if err != nil {
		return err
	}

	if emailVerificationData.ExpiresAt.Before(time.Now()) {
		return errors.New("token expired")
	}

	userData, err := s.UserRepository.GetUserById(ctx, emailVerificationData.UserID)
	if err != nil {
		return err
	}

	if userData.IsVerified == true {
		return errors.New("user already verified")
	}

	userData.IsVerified = true
	err = s.UserRepository.UpdateUser(ctx, userData)
	if err != nil {
		return err
	}

	return nil
}
