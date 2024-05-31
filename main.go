package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "Configure and download user's SSH Signing keys.")
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "USAGE")
		fmt.Fprintf(flag.CommandLine.Output(), "  %s\n\n", "gh ssh-signing <command> (args)")
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "COMMANDS")
		fmt.Fprintf(flag.CommandLine.Output(), "  %s\n", "  init:              Initialize your system for use with SSH Signing keys")
		fmt.Fprintf(flag.CommandLine.Output(), "  %s\n", "  import (username): Imports a new SSH Signing key for a particular user (must specify GitHub username)")
		fmt.Fprintf(flag.CommandLine.Output(), "  %s\n", "  export:            Export your SSH Signing keys to GitHub")
		fmt.Fprintf(flag.CommandLine.Output(), "  %s\n", "  list:              List all current SSH Signing keys")

		flag.PrintDefaults()
		fmt.Println()
	}
	var sshFile = flag.String("ssh-file", "", "Path to the SSH file")
	var configFile = flag.String("config-file", "", "Path to the allowedSignersFile")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		return
	}
	subcmd, args := args[0], args[1:]
	switch subcmd {
	case "init":
		initArgs := InitArgs{
			ConfigFile: *configFile,
			SshFile:    *sshFile,
		}
		InitCmd(initArgs)
	case "import":
		ImportCmd(args)
	case "list":
		ListCmd()
	case "export":
		ExportCmd()
	default:
		flag.Usage()
	}
}
