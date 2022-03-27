package cli

import (
	"errors"
	"fmt"

	"github.com/harunsasmaz/password-manager/internal/store"
	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

var get = &cli.Command{
	Name:      "get",
	Usage:     "get username and password for an alias",
	UsageText: "passman get <alias>",
	Category:  "Manager",
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
		fmt.Printf("Password is used for account: %s\n", creds.Source)
		clipboard.Write(clipboard.FmtText, []byte(creds.Password))
		fmt.Println("Copied password to clipboard!")

		return nil
	},
}
