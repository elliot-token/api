package service

import (
	"errors"
	"fmt"

	"github.com/elliot-token/api/app/domain"
)

var (
	ErrUserConflict = errors.New("username already exists")
)

type UserService interface {
	SignUp(user domain.UserEntity) error
}

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) SignUp(user domain.UserEntity) error {
	exist, err := u.userRepo.IsUsernameExist(user.Username)
	if err != nil {
		return fmt.Errorf("failed to check username: %w", err)
	}
	if exist {
		return fmt.Errorf("failed to save user: %w", ErrUserConflict)
	}

	if err := u.userRepo.SaveUser(user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}
