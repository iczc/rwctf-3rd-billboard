package internal

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/iczc/billboard/playground/common"
)

func getAddressByToken(token string) (string, error) {
	mnemonic, err := common.GenerateMnemonic(token)
	if err != nil {
		return "", err
	}

	var account uint32
	var index uint32
	hdPath := keys.CreateHDPath(account, index).String()

	var bip39Passphrase string
	privKey, err := common.GetPrivKeyFromMnemonic(mnemonic, bip39Passphrase, hdPath)
	if err != nil {
		return "", err
	}

	var address types.AccAddress = privKey.PubKey().Address().Bytes()

	return address.String(), nil
}
