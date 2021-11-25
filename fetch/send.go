package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jfamousket/go-kadena/common"
	"github.com/jfamousket/go-kadena/helpers"
)

// SendResponse is the response for a /send api call
type SendResponse struct {
	ReqKey string `json:"reqKey,omitempty"`
}

// Send sends a Pact command to a running Pact server and retrieves
// the transaction result
func Send(sendCmd []helpers.PrepareCommand, apiHost string) (res *SendResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(common.Error)
		}
	}()
	resp, err := sendRawCmds(sendCmd, apiHost)
	helpers.EnforceNoError(err)
	defer resp.Body.Close()
	err = helpers.UnMarshalBody(resp, res)
	helpers.EnforceNoError(err)
	return
}

func sendRawCmds(sendCmds []helpers.PrepareCommand, apiHost string) (*http.Response, error) {

	if apiHost == "" {
		return nil, fmt.Errorf("apiHost shouldn't be empty")
	}

	cmds := []common.Command{}
	for _, cmd := range sendCmds {
		if cmd.CmdType == common.CONT {
			cmds = append(cmds, helpers.PrepareContCmd(cmd))
		} else if cmd.CmdType == common.EXEC {
			cmds = append(cmds, helpers.PrepareExecCommand(cmd))
		}
	}

	body, err := json.Marshal(common.SendCommand{
		Cmds: cmds,
	})
	fmt.Println(string(body))
	helpers.EnforceNoError(err)

	bodyBytes := bytes.NewBuffer(body)
	return http.Post(fmt.Sprintf("%s/api/v1/send", apiHost), "application/json", bodyBytes)
}
