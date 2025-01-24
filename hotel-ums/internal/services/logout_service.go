package services

import (
	"context"
	"hotel-ums/internal/interfaces"
)

type LogoutService struct {
	UserRepository interfaces.IUserRepository
}

func NewLogoutService(userRepository interfaces.IUserRepository) *LogoutService {
	return &LogoutService{UserRepository: userRepository}
}

func (s *LogoutService) Logout(ctx context.Context, token string) error {
	return s.UserRepository.DeleteUserSession(ctx, token)
}
