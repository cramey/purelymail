package cmd

import (
	"flag"
	"fmt"
	"os"

	"purelymail/api"
)

func UserDelete(exec, cmd string, args []string) {
	flagset := flag.NewFlagSet(cmd, flag.ExitOnError)
	flagset.SetOutput(os.Stdout)
	flagset.Usage = func() {
		fmt.Printf("Usage: %s %s [args] <user@domain>\n\n", exec, cmd)
		fmt.Printf("Deletes a user in a specific domain.\n\n")
		fmt.Printf("Arguments:\n")
		flagset.PrintDefaults()
	}
	cpath := flagset.String("config", "", "path to configuration")
	flagset.Parse(args)

	if len(flagset.Args()) < 1 {
		fmt.Fprintf(os.Stderr,
			"user@domain required for user delete\n")
		flagset.Usage()
		os.Exit(1)
	}

	email := flagset.Arg(0)

	pm, err := api.Load(*cpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	err = pm.DeleteUser(email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
