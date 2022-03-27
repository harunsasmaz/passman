package cli

import (
	"errors"
	"github.com/harunsasmaz/password-manager/internal/store"
	"github.com/urfave/cli/v2"
)

var deletes = &cli.Command{
	Name:      "delete",
	Usage:     "delete username and password for an alias",
	UsageText: "passman delete <alias>",
	Category:  "Manager",
	Action: func(context *cli.Context) (err error) {
		if context.IsSet("all") {
			err = store.Clear()
			if err != nil {
				return errors.New("failed to delete all items")
			}

			return nil
		}

		err = store.Delete(context.String("a"))
		if err != nil {
			return errors.New("failed to delete credentials for alias: " + context.String("a"))
		}

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
			Usage: "delete all stored items",
		},
	},
}
