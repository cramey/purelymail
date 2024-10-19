package cmd

import (
	"fmt"
	"os"
	"strings"
)

func User(exec, cmd string, args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "%s %s: required argument missing\n", exec, cmd)
		printUserUsage(exec, cmd)
		os.Exit(1)
	}

	switch subcmd := strings.ToLower(args[0]); subcmd {
	case "create", "add", "a":
		UserAdd(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "delete", "del", "rm", "r":
		UserDelete(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "list", "ls", "l":
		UserList(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "--help", "help":
		printUserUsage(exec, cmd)

	default:
		fmt.Fprintf(os.Stderr, "%s %s %s is not a recognized command\n",
			exec, cmd, subcmd)
		printUserUsage(exec, cmd)
		os.Exit(1)
	}
	os.Exit(0)
}

func printUserUsage(exec, cmd string) {
	fmt.Printf("Usage: %s %s <command> ...\n", exec, cmd)
	fmt.Printf("See '%s %s <command> --help' for information", exec, cmd)
	fmt.Println(" on a specific command")
	fmt.Println("valid commands:")
	fmt.Println("    create         create a new user (aliases: add, a)")
	fmt.Println("    delete         delete a user (aliases: del, rm, r)")
	fmt.Println("    list           list users (aliases: ls, l)")
}
