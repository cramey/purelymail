package cmd

import (
	"fmt"
	"os"
	"strings"
)

func Domain(exec, cmd string, args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "%s %s: required argument missing\n", exec, cmd)
		printDomainUsage(exec, cmd)
		os.Exit(1)
	}

	switch subcmd := strings.ToLower(args[0]); subcmd {
	case "list", "ls", "l":
		DomainList(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "add", "a":
		DomainAdd(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "delete", "del", "rm", "r":
		DomainDelete(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "update", "up", "u":
		DomainUpdate(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "ownershipcode", "oc", "o":
		DomainOwnershipCode(exec, fmt.Sprintf("%s %s", cmd, subcmd), args[1:])

	case "--help", "help":
		printDomainUsage(exec, cmd)

	default:
		fmt.Fprintf(os.Stderr, "%s %s %s is not a recognized command\n",
			exec, cmd, subcmd)
		printDomainUsage(exec, cmd)
		os.Exit(1)
	}
	os.Exit(0)
}

func printDomainUsage(exec, cmd string) {
	fmt.Printf("Usage: %s %s <command> ...\n", exec, cmd)
	fmt.Printf("See '%s %s <command> --help' for information", exec, cmd)
	fmt.Println(" on a specific command")
	fmt.Println("valid commands:")
	fmt.Println("    add            add domain (aliases: a)")
	fmt.Println("    delete         delete domain (aliases: del, rm, r)")
	fmt.Println("    list           list domains (aliases: ls, l)")
	fmt.Println("    ownershipcode  get domain ownership code (aliases: oc, o)")
	fmt.Println("    update         update domain settings (aliases: up, u)")
}
