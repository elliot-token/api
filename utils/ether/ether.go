package ether

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetAddressFromSignature(signature, message string) (common.Address, error) {
	sigDecode, err := hexutil.Decode(signature)
	if err != nil {
		return common.Address{}, err
	}
	sigDecode[crypto.RecoveryIDOffset] -= 27 // Magic Mike

	msgHash := accounts.TextHash([]byte(message))
	publicKey, err := crypto.SigToPub(msgHash, sigDecode)
	if err != nil {
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*publicKey), nil
}
