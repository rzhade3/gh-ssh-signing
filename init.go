package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func InitCmd(args []string) {
	configFile, err := DefaultSignersFile()
	if err != nil {
		fmt.Println("Something went wrong while getting the default allowedSignersFile")
		return
	}
	if len(args) != 0 {
		configFile = args[0]
	}

	// First check if the Signers file exists
	// If it does not exist, create it
	fileExists, err := checkIfSignersFileExists()
	if err != nil {
		fmt.Println("Something went wrong while checking for the allowedSignersFile")
		return
	}
	if !fileExists {
		fmt.Println("allowedSignersFile not detected, linking existing or creating new file")
		filename, err := createOrLinkSignersFile(configFile)
		fmt.Printf("Created allowedSignersFile at: %s\n", filename)
		if err != nil {
			fmt.Printf("Something went wrong while creating the allowedSignersFile: %v\n", err)
			return
		}
	}
	fmt.Printf("allowedSignersFile has been set at %s\n", configFile)
}

func checkIfSignersFileExists() (bool, error) {
	stdoutStderr, err := exec.Command("git", "config", "--global", "--get", "gpg.ssh.allowedSignersFile").CombinedOutput()
	// This means that the config does not exist, unfortunately it also throws an error so we have a switch case here
	if err != nil && string(stdoutStderr) == "" {
		return false, nil
	}
	if string(stdoutStderr) != "" {
		// Now we have to check whether the file exists
		filename := strings.TrimSpace(string(stdoutStderr))
		_, err := os.Stat(filename)
		return !os.IsNotExist(err), nil
	}
	return false, err
}

func createOrLinkSignersFile(configFile string) (string, error) {
	configFileFolder := filepath.Dir(configFile)
	// Check if the file exists
	// If it does not exist, create it
	// If it exists, link it
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// path/to/whatever does not exist
		// create it
		os.MkdirAll(configFileFolder, os.ModePerm)
		_, err := os.Create(configFile)
		if err != nil {
			return "", err
		}
	}
	err := exec.Command("git", "config", "--global", "gpg.ssh.allowedSignersFile", configFile).Run()
	return configFile, err
}
