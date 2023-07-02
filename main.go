package main

import (
	"fmt"
	"os"
	"strings"

	"purelymail/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: required argument missing\n", os.Args[0])
		printDefaultUsage()
		os.Exit(1)
	}

	switch cn := strings.ToLower(os.Args[1]); cn {
	case "routing", "route", "r":
		cmd.Routing(os.Args[0], cn, os.Args[2:])

	case "domains", "domain", "d":
		cmd.Domain(os.Args[0], cn, os.Args[2:])

	case "users", "user", "u":
		cmd.User(os.Args[0], cn, os.Args[2:])

	case "--help", "help":
		printDefaultUsage()

	default:
		fmt.Fprintf(os.Stderr, "%s: '%s' is not a recognized command\n",
			os.Args[0], os.Args[1])
		printDefaultUsage()
		os.Exit(1)
	}

	os.Exit(0)
}

func printDefaultUsage() {
	fmt.Printf("Usage: %s <command> ...\n", os.Args[0])
	fmt.Printf("See '%s <command> --help' for information", os.Args[0])
	fmt.Println(" on a specific command")
	fmt.Println("valid commands:")
	fmt.Println("    domain       domain management")
	fmt.Println("    d            alias for domain")
	fmt.Println("    routing      routing rules management")
	fmt.Println("    r            alias for routing")
	fmt.Println("    user         user management")
	fmt.Println("    u            alias for user")
}
