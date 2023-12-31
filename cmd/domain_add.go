package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"purelymail/api"
)

func DomainAdd(exec, cmd string, args []string) {
	flagset := flag.NewFlagSet(cmd, flag.ExitOnError)
	flagset.SetOutput(os.Stdout)
	flagset.Usage = func() {
		fmt.Printf("Usage: %s %s [args] <domain>\n\n", exec, cmd)
		fmt.Printf("Adds a domain to be managed by Purelymail. ")
		fmt.Printf("To be succesful, the domain must have an ownership ")
		fmt.Printf("code set in a DNS TXT record.\n\n")
		fmt.Printf("Arguments:\n")
		flagset.PrintDefaults()
	}
	cpath := flagset.String("config", "", "path to configuration")
	flagset.Parse(args)

	if len(flagset.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "domain name required for domain add\n")
		flagset.Usage()
		os.Exit(1)
	}

	domain := strings.ToLower(flagset.Arg(0))

	pm, err := api.Load(*cpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	if err := pm.AddDomain(domain); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
