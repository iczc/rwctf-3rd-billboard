package internal

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/iczc/billboard/playground/common"
)

type result struct {
	Height    string       `json:"height"`
	TxHash    string       `json:"txhash"`
	CodeSpace string       `json:"codespace"`
	Code      int64        `json:"code"`
	RawLog    string       `json:"raw_log"`
	GasWanted string       `json:"gas_wanted"`
	GasUsed   string       `json:"gas_used"`
	Tx        *transaction `json:"tx"`
	Timestamp string       `json:"timestamp"`
	Error     string       `json:"error"`
}

type transaction struct {
	Type  string `json:"type"`
	Value struct {
		Msg []struct {
			Type  string `json:"type"`
			Value struct {
				Winner string `json:"winner"`
				ID     string `json:"id"`
			} `json:"value"`
		} `json:"msg"`
	} `json:"value"`
}

func queryTxInfo(lcd, txHash string) (*result, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	url := fmt.Sprintf("%s/txs/%s", lcd, txHash)
	resp := &result{}
	statusCode, err := common.HTTPGet(client, url, resp)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(resp.Error)
	}

	return resp, nil
}

func calcAddressByToken(token string) (string, error) {
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
