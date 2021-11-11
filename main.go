package main

import (
	"fmt"
	"log"

	"github.com/jfamousket/go-kadena/wallet"
)

func main() {
	res, err := wallet.CreateWallet("famousjingo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
