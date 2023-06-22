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
		fmt.Println("  .Name                   domain name")
		fmt.Println("  .AccountReset           is account reset allowed")
		fmt.Println("  .SymbolicSubaddressing  domain uses symbolic subaddressing")
		fmt.Println("  .Shared                 is domain shared")
		fmt.Println("  .DNSSummary             dns summary object")
		fmt.Println("    .MX                   does DNS pass MX check")
		fmt.Println("    .SPF                  does DNS pass SPF check")
		fmt.Println("    .DKIM                 does DNS pass DKIM check")
		fmt.Println("    .DMARC                does DNS pass DMARC check")
	}
	cpath := flagset.String("config", "", "path to configuration")
	shared := flagset.Bool("shared", false, "show shared domains")
	format := flagset.String(
		"format",
		"{{.Name}} {{if .Shared}}[shared]{{end}}",
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
