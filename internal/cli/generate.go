package cli

import (
	"fmt"

	"github.com/harunsasmaz/passman/internal/password"
	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

var generate = &cli.Command{
	Name:            "generate",
	Category:        "Generator",
	Usage:           "generate a random password.",
	UsageText:       "passman generate [FLAGS] [ARGS]\nIf you set --level, other options will be discarded",
	HideHelpCommand: true,
	Action: func(c *cli.Context) error {
		var pass string
		var err error
		if c.IsSet("l") {
			pass, err = password.Generate(password.Level(c.Int("l")))
		} else {
			pass, err = password.GenerateWithOptions(&password.Options{
				Length:     c.Int("n"),
				NumDigits:  c.Int("d"),
				NumSymbols: c.Int("s"),
				Uppercase:  c.Bool("u"),
				NoRepeat:   c.Bool("r"),
			})
		}

		if err != nil {
			return err
		}

		clipboard.Write(clipboard.FmtText, []byte(pass))
		fmt.Println("Copied to clipboard!")

		return nil
	},
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "level",
			Aliases: []string{"l"},
			Usage:   "choose a strength to use built-in options. Easy: 1, Mid: 2, Hard: 3. Example: -l 1",
		},
		&cli.IntFlag{
			Name:    "length",
			Aliases: []string{"n"},
			Value:   16,
			Usage:   "set the length of the password",
		},
		&cli.IntFlag{
			Name:    "digit",
			Aliases: []string{"d"},
			Value:   4,
			Usage:   "set the number of digits included in the password.",
		},
		&cli.IntFlag{
			Name:    "symbol",
			Aliases: []string{"s"},
			Value:   4,
			Usage:   "set the number of symbols included in the password.",
		},
		&cli.BoolFlag{
			Name:    "upper",
			Aliases: []string{"u"},
			Usage:   "set if password can contain uppercase letters.",
		},
		&cli.BoolFlag{
			Name:    "no-repeat",
			Aliases: []string{"r"},
			Usage:   "set if password should not contain repeated characters.",
		},
	},
}
