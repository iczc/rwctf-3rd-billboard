package internal

import (
	"errors"
	"strings"

	"github.com/iczc/billboard/playground/common"
)

type Verifier struct {
	lcd         string
	checkWinner bool
}

func NewFlagVerifier(lcd string, checkMode string) *Verifier {
	var checkWinner bool
	if checkMode == "0" {
		checkWinner = false
	} else {
		checkWinner = true
	}
	return &Verifier{
		lcd:         lcd,
		checkWinner: checkWinner,
	}
}

func (v *Verifier) ValidateTx(txHash, token string) error {
	result, err := common.QueryTx(v.lcd, strings.ToUpper(txHash))
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
		if address, err := getAddressByToken(token); err != nil || ctfMsg.Value.Winner != address {
			return errors.New("invalid winner")
		}
	}

	return nil
}
