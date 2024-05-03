package main

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

func AddCmd(args []string) {
	// Add a new SSH Signing key for a particular user (must specify GitHub username)
	if len(args) == 0 {
		fmt.Println("You must specify a GitHub username")
		return
	}
	username := args[0]
	fmt.Printf("Adding SSH Signing key for user: %s\n", username)

	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get user ID and email
	userResponse := struct {
		Login string
		Id    int
		Email string
	}{}
	userEndpoint := fmt.Sprintf("users/%s", username)
	err = client.Get(userEndpoint, &userResponse)
	if err != nil {
		fmt.Println("User not found")
		return
	}
	// This is when a user has set their email privacy to... private
	if userResponse.Email == "" {
		userResponse.Email = fmt.Sprintf("%s@users.noreply.github.com", userResponse.Login)
	}
	fmt.Println("Found user details")

	// Get user's keys
	keyResponse := []struct {
		Key   string
		Title string
	}{}
	keysEndpoint := fmt.Sprintf("user/%d/ssh_signing_keys", userResponse.Id)
	err = client.Get(keysEndpoint, &keyResponse)
	if err != nil {
		fmt.Printf("Error while making request to /users/%d/ssh_signing_keys\n", userResponse.Id)
		fmt.Println(err)
		return
	}
	if len(keyResponse) == 0 {
		fmt.Println("No SSH signing keys found for user")
		return
	}
	fmt.Println("Found user keys")

	// Now we simply write to file: `~/.config/git/allowedSignersFile`
	signersFile, err := GetSignersFile()
	if err != nil {
		fmt.Println("Error while getting allowedSignersFile, have you initialized the system?")
		fmt.Println(err)
		return
	}
	for _, key := range keyResponse {
		line := fmt.Sprintf("%s %s", userResponse.Email, key.Key)
		err = WriteToConfig(signersFile, line)
	}
	fmt.Println("Added SSH Signing key for user!")
}
