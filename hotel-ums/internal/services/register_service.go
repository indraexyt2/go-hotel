package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/models"
	"math/rand"
	"time"
)

type RegisterService struct {
	UserRepository interfaces.IUserRepository
	External       interfaces.IExternal
}

func NewRegisterService(userRepository interfaces.IUserRepository, ext interfaces.IExternal) *RegisterService {
	return &RegisterService{
		UserRepository: userRepository,
		External:       ext,
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

	notifyReq := &models.NotificationRequest{
		TemplateName: "account_activation",
		Recipient:    user.Email,
		Placeholder: map[string]string{
			"FullName":       user.FullName,
			"ActivationLink": "http://localhost:8080/api/ums/v1/email-verification/" + user.EmailVerificationToken.Token,
		},
	}

	jsonNotifyReq, _ := json.Marshal(notifyReq)
	err = s.External.SendMessageNotification(ctx, jsonNotifyReq)
	if err != nil {
		return nil, err
	}

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

func (s *RegisterService) ResendEmailVerification(ctx context.Context, email string) error {
	userData, err := s.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if userData.IsVerified == true {
		return errors.New("user already verified")
	}

	verificationTokenData, err := s.UserRepository.GetEmailVerificationTokenById(ctx, userData.ID)
	if err != nil {
		return err
	}

	tokenByte := make([]byte, 32)
	_, err = rand.Read(tokenByte)
	if err != nil {
		return err
	}

	verificationTokenData.Token = base64.URLEncoding.EncodeToString(tokenByte)
	verificationTokenData.ExpiresAt = time.Now().Add(time.Hour * 24)

	err = s.UserRepository.UpdateEmailVerificationToken(ctx, verificationTokenData)
	if err != nil {
		return err
	}

	return nil
}
