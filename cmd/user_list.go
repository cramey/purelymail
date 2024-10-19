package cmd

import (
	"flag"
	"fmt"
	"os"

	"purelymail/api"
)

func UserList(exec, cmd string, args []string) {
	flagset := flag.NewFlagSet(cmd, flag.ExitOnError)
	flagset.SetOutput(os.Stdout)
	flagset.Usage = func() {
		fmt.Printf("Usage: %s %s [args]\n\n", exec, cmd)
		fmt.Printf("Returns list of users")
		fmt.Printf("Arguments:\n")
		flagset.PrintDefaults()
	}
	cpath := flagset.String("config", "", "path to configuration")
	flagset.Parse(args)

	pm, err := api.Load(*cpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	users, err := pm.ListUsers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	for _, user := range *users {
		fmt.Println(user)
	}
}
