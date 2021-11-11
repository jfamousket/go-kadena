package wallet

import (
	"fmt"

	"github.com/islishude/bip32"
	"github.com/jfamousket/go-kadena/fetch"
	"github.com/tyler-smith/go-bip39"
)

func CreateWallet(password string) (*fetch.LocalResponse, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Seed Phrase: %s", mnemonic)

	privKey := bip32.NewRootXPrv(seed)

	return fetch.Local(fmt.Sprintf(`{"%s":(try false (coin.details "%s"))}`, privKey.XPub().String(), privKey.XPub().String()))
}
