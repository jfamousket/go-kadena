package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jfamousket/go-kadena/common"
)

func MarshalBody(value interface{}) (b *bytes.Buffer, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(common.Error)
		}
	}()
	body, err := json.Marshal(value)
	EnforceNoError(err)
	b = bytes.NewBuffer(body)
	return
}

func UnMarshalBody(resp *http.Response, returnType interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(common.Error)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	EnforceNoError(err)
	EnforceValid(resp.StatusCode == http.StatusOK, fmt.Errorf("%v", string(body)))
	err = json.Unmarshal(body, &returnType)
	EnforceNoError(err)
	return
}
