package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/bartekn/go-bip39"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/types"
	tmcrypto "github.com/tendermint/tendermint/crypto"
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

type Verifier struct {
	lcd         string
	checkWinner bool
}

func NewFlagVerifier(lcd string, checkMode string) *Verifier {
	var checkWinner bool
	if checkMode != "0" {
		checkWinner = true
	}
	return &Verifier{
		lcd:         lcd,
		checkWinner: checkWinner,
	}
}

func (v *Verifier) ValidateTx(token, txHash string) error {
	result, err := v.queryTxInfo(strings.ToUpper(txHash))
	if err != nil {
		return err
	}

	if result.Code != 0 {
		return errors.New("failed tx")
	}

	ctfMsg := result.Tx.Value.Msg[0]
	if ctfMsg.Type != "billboard/CaptureTheFlag" {
		return errors.New("invalid tx type")
	}

	if v.checkWinner {
		if address, err := calcAddressByToken(token); err != nil || ctfMsg.Value.Winner != address {
			return errors.New("invalid winner")
		}
	}

	return nil
}

func (v *Verifier) queryTxInfo(txHash string) (*result, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	url := fmt.Sprintf("%s/txs/%s", v.lcd, txHash)
	resp := &result{}
	statusCode, err := httpGet(client, url, resp)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(resp.Error)
	}

	return resp, nil
}

func httpGet(client *http.Client, url string, result interface{}) (int, error) {
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, json.Unmarshal(body, result)
}

func calcAddressByToken(token string) (string, error) {
	mnemonic, err := generateMnemonic(token)
	if err != nil {
		return "", err
	}

	var account uint32
	var index uint32
	hdPath := keys.CreateHDPath(account, index).String()

	var bip39Passphrase string
	privKey, err := getPrivKeyFromMnemonic(mnemonic, bip39Passphrase, hdPath)
	if err != nil {
		return "", err
	}

	var address types.AccAddress = privKey.PubKey().Address().Bytes()

	return address.String(), nil
}

func generateMnemonic(inputEntropy string) (string, error) {
	hashedEntropy := sha256.Sum256([]byte(inputEntropy))
	entropySeed := hashedEntropy[:]

	return bip39.NewMnemonic(entropySeed)
}

func getPrivKeyFromMnemonic(mnemonic string, bip39Passphrase, hdPath string) (tmcrypto.PrivKey, error) {
	derivedPriv, err := keys.SecpDeriveKey(mnemonic, bip39Passphrase, hdPath)
	if err != nil {
		return nil, err
	}

	privKey := keys.SecpPrivKeyGen(derivedPriv)
	return privKey, nil
}
