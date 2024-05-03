package main

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DefaultSignersFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".config/git/allowedSignersFile"), nil
}

func GetSignersFile() (string, error) {
	stdoutStderr, err := exec.Command("git", "config", "--global", "--get", "gpg.ssh.allowedSignersFile").CombinedOutput()
	if err != nil {
		return string(stdoutStderr), err
	}
	return strings.TrimSpace(string(stdoutStderr)), nil
}

func ReadConfig(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	// txtlines now holds all lines from the file
	return txtlines, nil
}

func WriteToConfig(filename string, line string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	datawriter := bufio.NewWriter(file)
	datawriter.WriteString(line + "\n")

	datawriter.Flush()
	file.Close()
	return nil
}
