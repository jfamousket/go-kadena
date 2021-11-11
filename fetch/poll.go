package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jfamousket/go-kadena/common"
)

type Request struct {
	RequestKeys string `json:"requestKeys,omitempty"`
}

type PollResponse struct {
	Gas          int64             `json:"gas,omitempty"`
	ReqKey       string            `json:"reqKey,omitempty"`
	TxId         string            `json:"txId,omitempty"`
	Logs         string            `json:"logs,omitempty"`
	MetaData     interface{}       `json:"metaData,omitempty"`
	Continuation interface{}       `json:"continuation,omitempty"`
	Events       common.PactEvents `json:"events,omitempty"`
}

type RequestKeys struct {
	RequestKeys []string `json:"requestKeys,omitempty"`
}

func Poll(requestKeys []string) (*PollResponse, error) {
	postBody, err := json.Marshal(RequestKeys{RequestKeys: requestKeys})
	if err != nil {
		log.Fatalln(err)
	}
	resBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/poll", common.TEST_PACT_URL), "application/json", resBody)
	if err != nil {
		log.Fatalf("An error occurred %+v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An error occurred %+v", err)
	}
	if resp.StatusCode == http.StatusOK {
		var ret PollResponse
		err = json.Unmarshal(body, &ret)
		if err != nil {
			log.Fatalf("An error occurred %+v", err)
		}
		return &ret, nil
	}
	return nil, fmt.Errorf("%v", body)
}
