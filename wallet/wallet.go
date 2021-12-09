package wallet

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jfamousket/go-pact"
)

func CreateWallet(password string) (res *pact.SendResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(pact.Error)
		}
	}()
	keyPair, err := pact.GenKeyPair(password)
	pact.EnforceNoError(err)
	now := time.Now()
	res, err = pact.Send([]pact.PrepareCommand{
		{
			CmdType:   pact.EXEC,
			PactCode:  "",
			EnvData:   "",
			NetworkId: "testnet",
			Meta: &pact.Meta{
				ChainId:      "0",
				Sender:       "",
				GasLimit:     1000,
				GasPrice:     1.0e-2,
				Ttl:          3600,
				CreationTime: uint64(now.Unix()),
			},
			KeyPairs: []pact.KeyPair{keyPair},
		},
	}, pact.TEST_PACT_URL)
	return
}

func SendSigned(signedCommand, apiHost string) (res *pact.SendResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(pact.Error)
		}
	}()
	cmd := struct {
		Cmds []string `json:"cmds"`
	}{
		Cmds: []string{signedCommand},
	}
	body, err := pact.MarshalBody(cmd)
	pact.EnforceNoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/send", apiHost), "application/json", body)
	pact.EnforceNoError(err)
	defer resp.Body.Close()
	err = pact.UnMarshalBody(resp, res)
	pact.EnforceNoError(err)
	return
}
