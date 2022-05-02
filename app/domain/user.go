package domain

type UserEntity struct {
	WalletAddress string
	Username      string
}

type UserRepository interface {
	IsUsernameExist(username string) (bool, error)
	IsWalletExist(walletAddr string) (bool, error)
	SaveUser(user *UserEntity) error
	GetUser(walletAddr string) (*UserEntity, error)
}
