package cli

import (
	"errors"
	"fmt"

	"github.com/harunsasmaz/password-manager/internal/store"
	"github.com/urfave/cli/v2"
)

var deletes = &cli.Command{
	Name:            "delete",
	Usage:           "delete account and password for an alias",
	UsageText:       "passman delete <alias>",
	Category:        "Manager",
	HideHelpCommand: true,
	Action: func(context *cli.Context) (err error) {
		if err = authenticate(); err != nil {
			return err
		}

		if context.IsSet("all") {
			err = store.Clear()
			if err != nil {
				return errors.New("failed to delete all items")
			}

			fmt.Println("Successfully deleted all credentials")
			return nil
		}

		if !context.IsSet("a") {
			return errors.New("alias is not provided")
		}

		err = store.Delete(context.String("a"))
		if err != nil {
			return errors.New("failed to delete credentials for alias: " + context.String("a"))
		}

		fmt.Printf("Successfully deleted credentials for alias: %s\n", context.String("a"))
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "alias",
			Aliases: []string{"a"},
			Usage:   "delete credentials for provided alias",
		},
		&cli.BoolFlag{
			Name:  "all",
			Usage: "delete all stored credentials",
		},
	},
}
