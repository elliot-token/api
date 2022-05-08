package domain

import "github.com/ethereum/go-ethereum/common"

type UserEntity struct {
	WalletAddress common.Address
	Username      string
}

type UserRepository interface {
	IsUsernameExist(username string) (bool, error)
	IsWalletExist(walletAddr common.Address) (bool, error)
	SaveUser(user *UserEntity) error
	GetUser(walletAddr common.Address) (*UserEntity, error)
}
