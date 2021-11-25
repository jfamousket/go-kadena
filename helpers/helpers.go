package helpers

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/jfamousket/go-kadena/common"
	"golang.org/x/crypto/blake2b"
)

func CreateBlake2Hash(data []byte) ([32]byte, string) {
	hash := blake2b.Sum256(data)
	return hash, base64.URLEncoding.EncodeToString(hash[:])
}

func EnforceType(value, valueType interface{}, msg string) {
	if ok := reflect.TypeOf(value) == reflect.TypeOf(valueType); !ok {
		panic(common.Error(fmt.Sprintf("%s must be a %t: %s", msg, valueType, value)))
	}
}

func EnforceNoError(err error) {
	if ok := err == nil; !ok {
		panic(common.Error(err.Error()))
	}
}

func EnforceValid(valid bool, err error) {
	if !valid {
		panic(common.Error(err.Error()))
	}
}

func MakeMeta(
	sender, chainId string,
	gasPrice float64,
	gasLimit uint64,
	creationTime uint64,
	ttl float64,
) *common.Meta {
	EnforceType(sender, "string", "sender")
	EnforceType(chainId, "string", "chainId")
	EnforceType(gasLimit, uint64(10), "gasLimit")
	EnforceType(creationTime, uint64(10), "creationTime")
	EnforceType(gasPrice, float64(10), "gasPrice")
	EnforceType(ttl, float64(10), "ttl")
	return &common.Meta{
		ChainId:      chainId,
		Sender:       sender,
		GasLimit:     gasLimit,
		GasPrice:     gasPrice,
		Ttl:          ttl,
		CreationTime: creationTime,
	}
}

func MakeSigner(keyPair common.KeyPair) common.Signer {
	return common.Signer{
		PubKey: hex.EncodeToString(keyPair.Public),
		Scheme: common.ED25519,
	}
}
