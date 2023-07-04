package cmd

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"purelymail/api"
)

func DomainList(exec, cmd string, args []string) {
	flagset := flag.NewFlagSet(cmd, flag.ExitOnError)
	flagset.SetOutput(os.Stdout)
	flagset.Usage = func() {
		fmt.Printf("Usage: %s %s [args]\n\n", exec, cmd)
		fmt.Printf("Lists domains whose emails are managed by Purelymail.\n\n")
		fmt.Printf("Arguments:\n")
		flagset.PrintDefaults()
		fmt.Println()
		fmt.Printf("Formatting output - format uses the go template format\n")
		fmt.Printf("Variables are enclosed with {{}} and are case sensitive\n")
		fmt.Printf("Available variables:\n")
		fmt.Printf("  .Name       (string) domain name\n")
		fmt.Printf("  .Summary    (string) summary of all attributes below\n")
		fmt.Printf("  .Reset      (bool) is account reset allowed\n")
		fmt.Printf("  .SubAddr    (bool) does domain use subaddressing\n")
		fmt.Printf("  .Shared     (bool) is domain shared\n")
		fmt.Printf("  .DNS.MX     (bool) does DNS pass MX check\n")
		fmt.Printf("  .DNS.SPF    (bool) does DNS pass SPF check\n")
		fmt.Printf("  .DNS.DKIM   (bool) does DNS pass DKIM check\n")
		fmt.Printf("  .DNS.DMARC  (bool) does DNS pass DMARC check\n")
	}
	cpath := flagset.String("config", "", "path to configuration")
	shared := flagset.Bool("shared", false, "show shared domains")
	format := flagset.String(
		"format",
		"{{.Name}} {{.Summary}}",
		"output format",
	)
	flagset.Parse(args)

	tmpl, err := template.New(cmd).Parse(*format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	pm, err := api.Load(*cpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	domains, err := pm.ListDomains(*shared)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	for _, domain := range *domains {
		if err := tmpl.Execute(os.Stdout, domain); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Println()
	}
}
