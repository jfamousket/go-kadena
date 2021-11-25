package fetch

import (
	"fmt"
	"net/http"

	"github.com/jfamousket/go-kadena/common"
	"github.com/jfamousket/go-kadena/helpers"
)

type LocalResponse struct {
	Gas          uint64        `json:"gas,omitempty"`
	Result       common.Result `json:"result,omitempty"`
	ReqKey       string        `json:"reqKey,omitempty"`
	Logs         string        `json:"logs,omitempty"`
	MetaData     interface{}   `json:"metaData,omitempty"`
	Continuation interface{}   `json:"continuation,omitempty"`
	TxId         string        `json:"txId,omitempty"`
}

func Local(localCmd helpers.PrepareCommand, apiHost string) (res *LocalResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(common.Error)
		}
	}()
	resp, err := localRawCmd(localCmd, apiHost)
	helpers.EnforceNoError(err)
	defer resp.Body.Close()
	err = helpers.UnMarshalBody(resp, res)
	helpers.EnforceNoError(err)
	return
}

func localRawCmd(localCmd helpers.PrepareCommand, apiHost string) (res *http.Response, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(common.Error)
		}
	}()
	helpers.EnforceType(apiHost, "string", "apiHost")
	helpers.EnforceValid(apiHost != "", fmt.Errorf("No api host provided"))
	cmd := helpers.PrepareExecCommand(localCmd)
	body, err := helpers.MarshalBody(cmd)
	helpers.EnforceNoError(err)
	return http.Post(fmt.Sprintf("%s/api/v1/local", apiHost), "application/json", body)
}
