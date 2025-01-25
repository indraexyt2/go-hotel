package services

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"hotel-ums/helpers"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/models"
)

type LoginService struct {
	UserRepository interfaces.IUserRepository
}

func NewLoginService(userRepo interfaces.IUserRepository) *LoginService {
	return &LoginService{UserRepository: userRepo}
}

func (s *LoginService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	userData, err := s.UserRepository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	if !userData.IsVerified {
		return nil, errors.New("user not verified")
	}

	token, err := helpers.GenerateToken(ctx, userData.ID, userData.FullName, userData.Email, userData.Role, "token")
	if err != nil {
		return nil, err
	}

	refreshToken, err := helpers.GenerateToken(ctx, userData.ID, userData.FullName, userData.Email, userData.Role, "refresh_token")
	if err != nil {
		return nil, err
	}

	userSession := &models.UserSession{
		UserID:       userData.ID,
		Token:        token,
		RefreshToken: refreshToken,
	}

	err = s.UserRepository.AddUserSession(ctx, userSession)
	if err != nil {
		return nil, err
	}

	resp := &models.LoginResponse{
		UserID:       userData.ID,
		FullName:     userData.FullName,
		Token:        token,
		RefreshToken: refreshToken,
	}

	return resp, nil
}

func (s *LoginService) RefreshToken(ctx context.Context, refreshToken string, claimsToken *helpers.Claims) (*models.RefreshTokenResponse, error) {
	token, err := helpers.GenerateToken(ctx, claimsToken.UserID, claimsToken.FullName, claimsToken.Email, claimsToken.Role, "token")
	if err != nil {
		return nil, err
	}

	err = s.UserRepository.UpdateUserSession(ctx, token, refreshToken)
	if err != nil {
		return nil, err
	}

	resp := &models.RefreshTokenResponse{Token: refreshToken}
	return resp, nil
}
