package main

import (
	"strings"
	"fmt"
	"os"
	"encoding/json"

	"github.com/cli/go-gh/v2/pkg/api"
)

func ExportCmd() {
	fmt.Println("Exporting all SSH Signing keys to GitHub")
	sshFile, err := GetSshFile()
	if err != nil {
		fmt.Println("Something went wrong while getting the SSH file")
		return
	}
	// Read file from the config File
	content, err := os.ReadFile(sshFile)
	if err != nil {
		fmt.Println("Something went wrong while reading the SSH file")
		return
	}
	sshKey := string(content)
	// Get the client
	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	keyRequest := struct {
		Key string `json:"key"`
		Title string `json:"title"`
	}{
		Key: sshKey,
		Title: "CLI",
	}
	keyEndpoint := "user/ssh_signing_keys"
	keyRequestBytes, err := json.Marshal(keyRequest)
	if err != nil {
		fmt.Println("Something went wrong while marshalling the key request")
		return
	}
	err = client.Post(keyEndpoint, strings.NewReader(string(keyRequestBytes)), nil)
	if err != nil {
		fmt.Println("Something went wrong while making the request to /user/ssh_signing_keys, %v", err)
		return
	}
	fmt.Println("Successfully exported SSH Signing key to GitHub")
}
