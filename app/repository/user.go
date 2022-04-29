package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/elliot-token/api/app/domain"
	"gorm.io/gorm"
)

type UserDBModel struct {
	WalletAddress string    `gorm:"column:wallet_address"`
	Username      string    `gorm:"column:username"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"updated_at"`
}

func (UserDBModel) TableName() string {
	return "users"
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) SaveUser(user domain.UserEntity) error {
	if err := u.db.Create(&UserDBModel{
		WalletAddress: user.WalletAddress,
		Username:      user.Username,
	}).Error; err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (u *userRepo) IsUsernameExist(username string) (bool, error) {
	if err := u.db.Where("username = ?", username).First(&UserDBModel{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to query username: %w", err)
	}
	return true, nil
}

func (u *userRepo) IsWalletExist(walletAddr string) (bool, error) {
	if err := u.db.Where("wallet_address = ?", walletAddr).First(&UserDBModel{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to query wallet address: %w", err)
	}
	return true, nil
}
