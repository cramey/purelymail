package cmd

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"
	"purelymail/api"
)

var alphabet []byte = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1',
	'2', '3', '4', '5', '6', '7', '8', '9',
}

func UserAdd(exec, cmd string, args []string) {
	flagset := flag.NewFlagSet(cmd, flag.ExitOnError)
	flagset.SetOutput(os.Stdout)
	flagset.Usage = func() {
		fmt.Printf("Usage: %s %s [args] <user@domain>\n",
			exec, cmd)
		flagset.PrintDefaults()
	}
	reset := flagset.Bool("reset", false, "enable password resetting")
	welcome := flagset.Bool("welcome", false, "send new user welcome email")
	indexing := flagset.Bool("indexing", true, "enable search indexing")
	password := flagset.String(
		"password", "prompt", "password collection scheme",
	)
	reset_email := flagset.String(
		"resetemail", "", "email to use for password reset",
	)
	reset_email_desc := flagset.String(
		"resetemaildescription", "", "description of email used for reset",
	)
	reset_phone := flagset.String(
		"resetphone", "", "phone number to use for password reset",
	)
	reset_phone_desc := flagset.String(
		"resetphonedescription", "", "description of email used for reset",
	)
	cpath := flagset.String("config", "", "path to configuration")
	flagset.Parse(args)

	if len(flagset.Args()) < 1 {
		fmt.Fprintf(os.Stderr,
			"user@domain required for user add\n")
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

	var pass string

	switch strings.ToLower(*password) {
	case "prompt":
		fmt.Printf("Password: ")
		pw, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Printf("\n")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		pass = string(pw)

	case "random":
		// Generate a random password
		rand.Seed(time.Now().UnixNano())
		sz := rand.Intn(15) + 20
		pw := make([]byte, sz)
		for i, _ := range pw {
			pw[i] = alphabet[rand.Intn(len(alphabet))]
		}
		pass = string(pw)

	default:
		fmt.Fprintf(
			os.Stderr, "unsupported password collection method: %s\n", pass,
		)
		os.Exit(1)
	}

	if len(pass) < 1 {
		fmt.Fprintf(os.Stderr, "password required for add\n")
		os.Exit(1)
	}

	pm, err := api.Load(*cpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	err = pm.CreateUser(api.UserConfig{
		User: user, Domain: domain,
		Password:                 pass,
		Reset:                    *reset,
		RecoveryEmail:            *reset_email,
		RecoveryEmailDescription: *reset_email_desc,
		RecoveryPhone:            *reset_phone,
		RecoveryPhoneDescription: *reset_phone_desc,
		Indexing:                 *indexing,
		Welcome:                  *welcome,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	if strings.ToLower(*password) == "random" {
		fmt.Printf("Generated password: %s\n", pass)
	}
}
