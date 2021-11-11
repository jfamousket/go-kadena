package helpers

import (
	"encoding/base64"

	"golang.org/x/crypto/blake2b"
)

func CreateBlake2Hash(data []byte) ([32]byte, string) {
	hash := blake2b.Sum256(data)
	return hash, base64.URLEncoding.EncodeToString(hash[:])
}
