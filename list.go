package main

import (
	"fmt"
)

// ListCmd lists all the allowed signers
func ListCmd() {
	signersFile, err := GetSignersFile()
	if err != nil {
		fmt.Println("Error while getting allowedSignersFile, have you initialized the system?")
		return
	}

	signers, err := ReadConfig(signersFile)
	if err != nil {
		fmt.Println("Error while reading allowedSignersFile")
		return
	}
	for _, signer := range signers {
		fmt.Println(signer)
	}
}
