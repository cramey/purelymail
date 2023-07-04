package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"purelymail/api"
)

func RoutingDelete(exec, cmd string, args []string) {
	flagset := flag.NewFlagSet(cmd, flag.ExitOnError)
	flagset.SetOutput(os.Stdout)
	flagset.Usage = func() {
		fmt.Printf("Usage: %s %s [args] <id>\n\n", exec, cmd)
		fmt.Printf("Deletes an existing routing rule by ID.\n\n")
		fmt.Printf("Arguments:\n")
		flagset.PrintDefaults()
	}
	cpath := flagset.String("config", "", "path to configuration")
	flagset.Parse(args)

	if len(flagset.Args()) < 1 {
		fmt.Fprintf(os.Stderr,
			"id required for routing rule deletion\n")
		flagset.Usage()
		os.Exit(1)
	}

	id, err := strconv.ParseInt(flagset.Arg(0), 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid id\n")
		flagset.Usage()
		os.Exit(1)
	}

	pm, err := api.Load(*cpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	err = pm.DeleteRoutingRule(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
