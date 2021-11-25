package fetch

import (
	"fmt"
	"net/http"

	"github.com/jfamousket/go-kadena/common"
	"github.com/jfamousket/go-kadena/helpers"
)

type ListenResponse struct {
	Gas          uint64        `json:"gas"`
	Result       common.Result `json:"result"`
	ReqKey       string        `json:"reqKey"`
	Logs         string        `json:"logs"`
	MetaData     interface{}   `json:"metaData"`
	Continuation interface{}   `json:"continuation"`
	TxId         interface{}   `json:"txId"`
}

func Listen(requestKey string, apiHost string) (res *ListenResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(common.Error)
		}
	}()
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/listen", apiHost))
	helpers.EnforceNoError(err)
	defer resp.Body.Close()
	err = helpers.UnMarshalBody(resp, res)
	helpers.EnforceNoError(err)
	return
}
