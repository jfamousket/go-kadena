package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jfamousket/go-kadena/common"
	"github.com/jfamousket/go-kadena/helpers"
)

type LocalResult struct {
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type LocalResponse struct {
	Gas          uint64      `json:"gas,omitempty"`
	Result       LocalResult `json:"result,omitempty"`
	ReqKey       string      `json:"reqKey,omitempty"`
	Logs         string      `json:"logs,omitempty"`
	MetaData     interface{} `json:"metaData,omitempty"`
	Continuation interface{} `json:"continuation,omitempty"`
	TxId         string      `json:"txId,omitempty"`
}

func Local(code string) (*LocalResponse, error) {
	now := time.Now()
	cmd := common.CommandField{
		Nonce: now.Format("2006-01-02"),
		Meta: common.Meta{
			ChainId:      "0",
			Sender:       "",
			GasLimit:     1000,
			GasPrice:     1.0e-2,
			Ttl:          3600,
			CreationTime: uint64(now.Unix()),
		},
		Signers: []common.Signer{},
		Payload: common.Payload{
			Exec: common.Exec{
				Data: "",
				Code: code,
			},
			Cont: common.Cont{
				PactId:   "",
				Rollback: true,
				Step:     1,
				Data:     nil,
				Proof:    "",
			},
		},
		NetworkId: "testnet",
	}
	cmdString, err := json.Marshal(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	_, hash := helpers.CreateBlake2Hash(cmdString)
	postBody, err := json.Marshal(common.Command{
		Cmd:  string(cmdString),
		Hash: hash,
		Sigs: []common.Sig{},
	})
	if err != nil {
		log.Fatalln(err)
	}
	resBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/local", common.TEST_PACT_URL), "application/json", resBody)
	if err != nil {
		log.Fatalf("An error occurred %+v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An error occurred %+v", err)
	}
	if resp.StatusCode == http.StatusOK {
		var ret LocalResponse
		err = json.Unmarshal(body, &ret)
		if err != nil {
			log.Fatalf("An error occurred %+v", err)
		}
		return &ret, nil
	}
	return nil, fmt.Errorf("%v", body)
}
