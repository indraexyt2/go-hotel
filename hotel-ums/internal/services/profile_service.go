package services

import (
	"context"
	"hotel-ums/internal/interfaces"
	"hotel-ums/internal/models"
)

type ProfileService struct {
	UserRepository interfaces.IUserRepository
}

func NewProfileService(userRepository interfaces.IUserRepository) *ProfileService {
	return &ProfileService{UserRepository: userRepository}
}

func (s *ProfileService) UpdateUserProfile(ctx context.Context, user *models.User, photoPath string, userID int) (*models.User, error) {
	userData, err := s.UserRepository.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}

	if photoPath != "" {
		userData.PhotoPath = photoPath
	}

	if user.FullName != "" {
		userData.FullName = user.FullName
	}

	if user.Email != "" {
		userData.Email = user.Email
	}

	if user.Phone != "" {
		userData.Phone = user.Phone
	}

	if user.Address != "" {
		userData.Address = user.Address
	}

	err = s.UserRepository.UpdateUser(ctx, userData)
	if err != nil {
		return nil, err
	}

	resp := userData

	return resp, nil
}
