package wallet

import (
	"crypto/ed25519"
	"fmt"
	"net/http"
	"time"

	"github.com/islishude/bip32"
	"github.com/jfamousket/go-kadena/common"
	"github.com/jfamousket/go-kadena/fetch"
	"github.com/jfamousket/go-kadena/helpers"
	"github.com/tyler-smith/go-bip39"
)

func CreateWallet(password string) (res *fetch.SendResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(common.Error)
		}
	}()

	entropy, err := bip39.NewEntropy(128)
	helpers.EnforceNoError(err)

	mnemonic, err := bip39.NewMnemonic(entropy)
	helpers.EnforceNoError(err)

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	helpers.EnforceNoError(err)

	fmt.Printf("Seed Phrase: %s", mnemonic)

	privKey := bip32.NewRootXPrv(seed)

	now := time.Now()
	res, err = fetch.Send([]helpers.PrepareCommand{
		{
			CmdType:   common.EXEC,
			PactCode:  "",
			EnvData:   "",
			NetworkId: "testnet",
			Meta: &common.Meta{
				ChainId:      "0",
				Sender:       "",
				GasLimit:     1000,
				GasPrice:     1.0e-2,
				Ttl:          3600,
				CreationTime: uint64(now.Unix()),
			},
			KeyPairs: []common.KeyPair{
				{
					Private: ed25519.PrivateKey(privKey.String()),
					Public:  privKey.PublicKey(),
				},
			},
		},
	}, common.TEST_PACT_URL)
	return
}

func SendSigned(signedCommand, apiHost string) (res *fetch.SendResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(common.Error)
		}
	}()
	cmd := struct {
		Cmds []string `json:"cmds"`
	}{
		Cmds: []string{signedCommand},
	}
	body, err := helpers.MarshalBody(cmd)
	helpers.EnforceNoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/send", apiHost), "application/json", body)
	helpers.EnforceNoError(err)
	defer resp.Body.Close()
	err = helpers.UnMarshalBody(resp, res)
	helpers.EnforceNoError(err)
	return
}
