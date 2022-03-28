package cli

import (
	"errors"
	"fmt"

	"github.com/harunsasmaz/passman/internal/password"
	"github.com/harunsasmaz/passman/internal/store"
	"github.com/urfave/cli/v2"
)

var create = &cli.Command{
	Name:            "create",
	Usage:           "create a new account/password pair with an alias",
	UsageText:       "passman create [FLAGS] [ARGS]",
	Category:        "Manager",
	HideHelpCommand: true,
	Action: func(context *cli.Context) (err error) {
		if !context.IsSet("a") || !context.IsSet("u") {
			return errors.New("alias and account must be provided")
		}

		if !context.IsSet("p") && !context.IsSet("g") {
			return errors.New("either password or generate flag must be provided")
		}

		pass := context.String("p")
		if context.IsSet("g") {
			pass, err = password.Generate(password.LevelHard)
			if err != nil {
				return errors.New("failed to generate password to save")
			}
		}

		creds := credentials{
			Source:   context.String("u"),
			Password: pass,
		}

		err = store.Put(context.String("a"), creds)
		if err != nil {
			return errors.New("failed to save credentials")
		}

		fmt.Printf("Successfully created credentials for alias: %s\n", context.String("a"))
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "alias",
			Aliases: []string{"a"},
			Usage:   "set an alias for the new credentials.",
		},
		&cli.StringFlag{
			Name:    "account",
			Aliases: []string{"u"},
			Usage:   "set account or host that password will be used for.",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			Usage:   "set password.",
		},
		&cli.BoolFlag{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generates a new secure password to save.",
		},
	},
}
