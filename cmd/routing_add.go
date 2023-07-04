package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"purelymail/api"
)

func RoutingAdd(exec, cmd string, args []string) {
	flagset := flag.NewFlagSet(cmd, flag.ExitOnError)
	flagset.SetOutput(os.Stdout)
	flagset.Usage = func() {
		fmt.Printf("Usage: %s %s [args] <user@domain> <to ...>\n\n", exec, cmd)
		fmt.Printf("Adds an email routing rule. Used for aliases, catch-all ")
		fmt.Printf("addresses, and prefix-matched addresses.\n\n")
		fmt.Printf("Arguments:\n")
		flagset.PrintDefaults()
	}
	prefix := flagset.Bool("prefix", false, "is added rule a prefix")
	cpath := flagset.String("config", "", "path to configuration")
	flagset.Parse(args)

	if len(flagset.Args()) < 2 {
		fmt.Fprintf(os.Stderr,
			"user@domain and to address required for routing rule add\n")
		flagset.Usage()
		os.Exit(1)
	}

	sp := strings.Split(flagset.Arg(0), "@")
	if len(sp) != 2 {
		fmt.Fprintf(os.Stderr, "invalid format for user@domain\n")
		flagset.Usage()
		os.Exit(1)
	}
	domain := sp[1]
	user := sp[0]
	addresses := flagset.Args()[1:]

	pm, err := api.Load(*cpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	err = pm.CreateRoutingRule(domain, user, *prefix, addresses)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
