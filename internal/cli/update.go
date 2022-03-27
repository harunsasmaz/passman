package cli

import (
	"errors"
	"github.com/harunsasmaz/password-manager/internal/password"
	"github.com/harunsasmaz/password-manager/internal/store"
	"github.com/urfave/cli/v2"
)

var update = &cli.Command{
	Name:      "update",
	Usage:     "update an existing username/password pair with an alias",
	UsageText: "passman update [FLAGS] [ARGS]",
	Category:  "Manager",
	Action: func(context *cli.Context) (err error) {
		err = store.Get(context.String("a"), nil)
		if err != nil {
			return errors.New("there is no credentials with provided alias")
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

		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "alias",
			Aliases: []string{"a"},
			Usage:   "set an alias for the new credentials.",
		},
		&cli.StringFlag{
			Name:    "username",
			Aliases: []string{"u"},
			Usage:   "set username or host that password will be used for.",
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
