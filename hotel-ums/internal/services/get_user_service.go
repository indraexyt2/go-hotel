package services

import (
	"context"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/models"
)

type GetUserService struct {
	UserRepository interfaces.IUserRepository
}

func NewGetUserService(userRepo interfaces.IUserRepository) *GetUserService {
	return &GetUserService{UserRepository: userRepo}
}

func (s *GetUserService) GetUser(ctx context.Context, id int) (*models.User, error) {
	resp, err := s.UserRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	resp.Password = ""
	return resp, nil
}

func (s *GetUserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.UserRepository.GetAllUsers(ctx)
}
