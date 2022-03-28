package cli

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/harunsasmaz/passman/internal/store"
	"github.com/urfave/cli/v2"
)

var get = &cli.Command{
	Name:            "get",
	Usage:           "get account and password for an alias",
	UsageText:       "passman get <alias>",
	Category:        "Manager",
	HideHelpCommand: true,
	Action: func(context *cli.Context) error {
		if context.NArg() < 1 {
			return errors.New("alias is not provided")
		}

		if err := authenticate(); err != nil {
			return err
		}

		var creds credentials
		err := store.Get(context.Args().Get(0), &creds)
		if err != nil {
			return errors.New("alias not found")
		}

		fmt.Println("Successfully retrieved password!")
		err = clipboard.WriteAll(creds.Password)
		if err != nil {
			return errors.New("failed to copy password to clipboard")
		}

		fmt.Printf("Password is used for account: %s\n", creds.Source)
		fmt.Println("Copied password to clipboard!")

		return nil
	},
}
