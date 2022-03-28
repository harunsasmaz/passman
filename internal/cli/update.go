package cli

import (
	"errors"
	"fmt"

	"github.com/harunsasmaz/password-manager/internal/password"
	"github.com/harunsasmaz/password-manager/internal/store"
	"github.com/urfave/cli/v2"
)

var update = &cli.Command{
	Name:            "update",
	Usage:           "update an existing account or password with an alias",
	UsageText:       "passman update [FLAGS] [ARGS]",
	Category:        "Manager",
	HideHelpCommand: true,
	Action: func(context *cli.Context) (err error) {
		if err = authenticate(); err != nil {
			return err
		}

		var creds credentials
		err = store.Get(context.String("a"), &creds)
		if err != nil {
			return errors.New("there is no credentials with provided alias")
		}

		if context.IsSet("g") {
			creds.Password, err = password.Generate(password.LevelHard)
			if err != nil {
				return errors.New("failed to generate password to update")
			}
			fmt.Println("Password renewed!")
		}

		if context.IsSet("p") {
			creds.Password = context.String("p")
			fmt.Println("Password renewed!")
		}

		if context.IsSet("u") {
			creds.Source = context.String("u")
			fmt.Println("Account name renewed!")
		}

		err = store.Put(context.String("a"), creds)
		if err != nil {
			return errors.New("failed to save credentials")
		}

		fmt.Println("Successfully saved changes!")
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "alias",
			Aliases: []string{"a"},
			Usage:   "set alias that you want to update.",
		},
		&cli.StringFlag{
			Name:    "account",
			Aliases: []string{"u"},
			Usage:   "set account or host if you want to update.",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			Usage:   "set password if you want to update.",
		},
		&cli.BoolFlag{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generates a new secure password to update.",
		},
	},
}
