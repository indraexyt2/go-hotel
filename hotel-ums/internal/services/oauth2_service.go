package services

import (
	"context"
	"golang.org/x/oauth2"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/models"
)

type IOAuth2Service struct {
	UserRepository interfaces.IUserRepository
}

func NewOAuth2Service(userRepository interfaces.IUserRepository) *IOAuth2Service {
	return &IOAuth2Service{UserRepository: userRepository}
}

func (s *IOAuth2Service) GoogleLogin(ctx context.Context, req *models.GoogleUserInfo, token *oauth2.Token) error {
	user, _ := s.UserRepository.GetUserByEmail(ctx, req.Email)
	if user == nil {
		userRequest := &models.User{
			PhotoPath:  req.Picture,
			Username:   req.Email,
			Password:   "",
			Email:      req.Email,
			Role:       "guest",
			FullName:   req.Name,
			IsVerified: req.EmailVerified,
			Source:     "google",
		}

		err := s.UserRepository.RegisterNewUser(ctx, userRequest)
		if err != nil {
			return err
		}
	}

	user, err := s.UserRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	userSession := &models.UserSession{
		UserID:       user.ID,
		Token:        token.AccessToken,
		RefreshToken: token.RefreshToken,
		Source:       "google",
		ExpiresAt:    token.Expiry,
		TokenType:    token.TokenType,
	}

	err = s.UserRepository.AddUserSession(ctx, userSession)
	if err != nil {
		return err
	}

	return nil
}
