package cmd

import (
	"fmt"
	"os"
	"strings"
)

func Routing(exec, cmd string, args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "%s %s: required argument missing\n", exec, cmd)
		printMailboxUsage(exec, cmd)
		os.Exit(1)
	}

	switch subcmd := strings.ToLower(args[0]); subcmd {
	case "list", "ls", "l":
		RoutingList(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "add", "a":
		RoutingAdd(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "delete", "del", "rm", "r":
		RoutingDelete(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "--help", "help":
		printMailboxUsage(exec, cmd)

	default:
		fmt.Fprintf(os.Stderr, "%s %s %s is not a recognized command\n",
			exec, cmd, subcmd)
		printMailboxUsage(exec, cmd)
		os.Exit(1)
	}
	os.Exit(0)
}

func printMailboxUsage(exec, cmd string) {
	fmt.Printf("Usage: %s %s <command> ...\n", exec, cmd)
	fmt.Printf("See '%s %s <command> --help' for information", exec, cmd)
	fmt.Println(" on a specific command")
	fmt.Println("valid commands:")
	fmt.Println("    add      add routing rule (aliases: a)")
	fmt.Println("    delete   delete routing rule (aliases: del, rm, r)")
	fmt.Println("    list     list routing rules (aliases: ls, l)")
}
