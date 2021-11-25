package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jfamousket/go-kadena/common"
	"github.com/jfamousket/go-kadena/helpers"
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

func Poll(requestKeys []string, apiHost string) (res *PollResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(common.Error)
		}
	}()
	postBody, err := json.Marshal(RequestKeys{RequestKeys: requestKeys})
	helpers.EnforceNoError(err)
	req := bytes.NewBuffer(postBody)
	resp, err := http.Post(
		fmt.Sprintf("%s/api/v1/poll", apiHost),
		"application/json",
		req,
	)
	helpers.EnforceNoError(err)
	defer resp.Body.Close()
	err = helpers.UnMarshalBody(resp, res)
	helpers.EnforceNoError(err)
	return
}
