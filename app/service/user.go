package service

import (
	"errors"
	"fmt"

	"github.com/elliot-token/api/app/domain"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrWalletConflict   = errors.New("wallet address already exists")
	ErrUsernameConflict = errors.New("username already exists")
	ErrWalletNotFound   = errors.New("wallet address not found")
)

type UserService interface {
	SignUp(user *domain.UserEntity) error
	GetUser(walletAddr common.Address) (*domain.UserEntity, error)
}

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) SignUp(user *domain.UserEntity) error {
	if exist, err := u.userRepo.IsWalletExist(user.WalletAddress); err != nil {
		return fmt.Errorf("failed to check wallet address: %w", err)
	} else if exist {
		return fmt.Errorf("failed to save user: %w", ErrWalletConflict)
	}

	if exist, err := u.userRepo.IsUsernameExist(user.Username); err != nil {
		return fmt.Errorf("failed to check username: %w", err)
	} else if exist {
		return fmt.Errorf("failed to save user: %w", ErrUsernameConflict)
	}

	if err := u.userRepo.SaveUser(user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (u *userService) GetUser(walletAddr common.Address) (*domain.UserEntity, error) {
	return u.userRepo.GetUser(walletAddr)
}
