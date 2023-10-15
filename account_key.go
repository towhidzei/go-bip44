package bip44

import (
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
)

type AccountKey struct {
	extendedKey *hdkeychain.ExtendedKey
	startPath   HDStartPath
}

func NewAccountKeyFromXPubKey(value string) (*AccountKey, error) {
	xKey, err := hdkeychain.NewKeyFromString(value)

	if err != nil {
		return nil, err
	}

	return &AccountKey{
		extendedKey: xKey,
	}, nil
}
