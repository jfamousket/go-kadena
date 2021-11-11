package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jfamousket/go-kadena/common"
)

type InfoResponse struct {
	NodeNumberOfChains uint64      `json:"nodeNumberOfChains"`
	NodeApiVersion     string      `json:"nodeApiVersion"`
	NodeChains         []string    `json:"nodeChains"`
	NodeVersion        string      `json:"nodeVersion"`
	NodeGraphHistory   interface{} `json:"nodeGraphHistory"`
}

func Info() (*InfoResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/info", common.TEST_URL))
	if err != nil {
		log.Fatalf("An error occurred %+v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An error occurred %+v", err)
	}
	if resp.StatusCode == http.StatusOK {
		var ret InfoResponse
		err = json.Unmarshal(body, &ret)
		if err != nil {
			log.Fatalf("An error occurred %+v", err)
		}
		return &ret, nil
	}
	return nil, fmt.Errorf("%v", body)
}
