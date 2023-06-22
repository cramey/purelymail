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
		fmt.Printf("Usage: %s %s [args]\n",
			exec, cmd)
		flagset.PrintDefaults()
		fmt.Println()
		fmt.Println("Formatting output - format uses the go template format")
		fmt.Println("Variables are enclosed with {{}} and are case sensitive")
		fmt.Println("Available variables:")
		fmt.Println("  .Name       (string) domain name")
		fmt.Println("  .Summary    (string) summary of all attributes below")
		fmt.Println("  .Reset      (bool) is account reset allowed")
		fmt.Println("  .SubAddr    (bool) does domain use subaddressing")
		fmt.Println("  .Shared     (bool) is domain shared")
		fmt.Println("  .DNS.MX     (bool) does DNS pass MX check")
		fmt.Println("  .DNS.SPF    (bool) does DNS pass SPF check")
		fmt.Println("  .DNS.DKIM   (bool) does DNS pass DKIM check")
		fmt.Println("  .DNS.DMARC  (bool) does DNS pass DMARC check")
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
