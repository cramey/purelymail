package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"purelymail/api"
)

func UpdateDomain(exec, cmd string, args []string) {
	flagset := flag.NewFlagSet(cmd, flag.ExitOnError)
	flagset.SetOutput(os.Stdout)
	flagset.Usage = func() {
		fmt.Printf("Usage: %s %s [args] <domain>\n",
			exec, cmd)
		flagset.PrintDefaults()
	}
	reset := flagset.Bool("reset", false, "allow account resetting")
	subaddr := flagset.Bool("subaddr", false, "allow symbolic subaddressing")
	dns := flagset.Bool("dns", false, "recheck DNS")
	cpath := flagset.String("config", "", "path to configuration")
	flagset.Parse(args)

	if len(flagset.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "domain name required for domain update\n")
		flagset.Usage()
		os.Exit(1)
	}

	domain := strings.ToLower(flagset.Arg(0))

	pm, err := api.Load(*cpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	if err := pm.UpdateDomain(domain, *reset, *subaddr, *dns); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
