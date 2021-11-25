package fetch

import (
	"fmt"
	"net/http"

	"github.com/jfamousket/go-kadena/common"
	"github.com/jfamousket/go-kadena/helpers"
)

func SPV(spvCmd common.SPVCommand, apiHost string) (res interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(common.Error)
		}
	}()
	helpers.EnforceType(spvCmd.TargetChainId, "string", "targetChainId")
	helpers.EnforceType(spvCmd.RequestKey, "string", "requestKey")
	req, err := helpers.MarshalBody(spvCmd)
	helpers.EnforceNoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/spv", apiHost), "application/json", req)
	helpers.EnforceNoError(err)
	err = helpers.UnMarshalBody(resp, res)
	helpers.EnforceNoError(err)
	return
}
