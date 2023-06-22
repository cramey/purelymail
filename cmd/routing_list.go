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
		fmt.Printf("Usage: %s %s [args]\n",
			exec, cmd)
		flagset.PrintDefaults()
		fmt.Println()
		fmt.Println("Formatting output - format uses the go template format")
		fmt.Println("Variables are enclosed with {{}} and are case sensitive")
		fmt.Println("Available variables:")
		fmt.Println("  .ID         (number) routing rule's unique id")
		fmt.Println("  .Summary    (string) summary of all fields below")
		fmt.Println("  .Domain     (string) routing rule's domain")
		fmt.Println("  .MatchUser  (string) user matched by routing rule")
		fmt.Println("  .Prefix     (bool) is user matched only starting with")
		fmt.Println("  .Addrs      (string array) list of addresses to send to")
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
