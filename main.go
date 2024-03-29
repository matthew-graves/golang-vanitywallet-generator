package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"

	"time"
)

func GenWallet() {
	account := crypto.GenerateAccount()
	if account.Address.String()[0] == strings.ToUpper(os.Args[1])[0] {
		if strings.HasPrefix(account.Address.String(), strings.ToUpper(os.Args[1])) {
			fmt.Println("\n----------------------------------")
			fmt.Println("Wallet Found!")
			fmt.Println(account.Address.String())
			mnemonic, err := mnemonic.FromPrivateKey(account.PrivateKey)
			if err != nil {
				return
			}
			fmt.Println(mnemonic)
			os.Exit(0)
		}
	}
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println()
		fmt.Println("No parameters provided, please add a word or phrase to search for.")
		fmt.Println("Example Usage: ./vanitywalletgenerator.exe LIBERTY")
		fmt.Println()
		os.Exit(1)
	}
	if strings.ContainsAny(os.Args[1], "0189") {
		fmt.Println("Invalid Input, Characters 0, 1, 8, and 9 are not supported in Algorand Wallets")
		os.Exit(1)
	}
	fmt.Printf("Start Time: %s\n", time.Now())
	x := 1
	starttime := time.Now()
	for {
		go GenWallet()
		x += 1
		elapsed := time.Since(starttime).Seconds()
		if elapsed > 1 {
			starttime = time.Now()
			fmt.Print("\r                                  ")
			fmt.Printf("\r %d wallets per second", x)
			x = 0
		}
	}
}
