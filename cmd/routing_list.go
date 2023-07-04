package cmd

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"purelymail/api"
)

func RoutingList(exec, cmd string, args []string) {
	flagset := flag.NewFlagSet(cmd, flag.ExitOnError)
	flagset.SetOutput(os.Stdout)
	flagset.Usage = func() {
		fmt.Printf("Usage: %s %s [args]\n\n", exec, cmd)
		fmt.Printf("Lists routing rules.\n\n")
		fmt.Printf("Arguments:\n")
		flagset.PrintDefaults()
		fmt.Println()
		fmt.Printf("Formatting output - format uses the go template format\n")
		fmt.Printf("Variables are enclosed with {{}} and are case sensitive\n")
		fmt.Printf("Available variables:\n")
		fmt.Printf("  .ID         (number) routing rule's unique id\n")
		fmt.Printf("  .Summary    (string) summary of all fields below\n")
		fmt.Printf("  .Domain     (string) routing rule's domain\n")
		fmt.Printf("  .MatchUser  (string) user matched by routing rule\n")
		fmt.Printf("  .Prefix     (bool) is user matched only starting with\n")
		fmt.Printf("  .Addrs      (string array) list of addresses to send to\n")
	}
	cpath := flagset.String("config", "", "path to configuration")
	format := flagset.String("format", "{{.ID}} {{.Summary}}", "output format")
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

	rules, err := pm.ListRoutingRules()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	for _, rule := range *rules {
		if err := tmpl.Execute(os.Stdout, rule); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Println()
	}
}
