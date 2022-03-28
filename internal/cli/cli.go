package cli

import (
	"errors"
	"fmt"

	"github.com/harunsasmaz/passman/internal/password"
	"github.com/harunsasmaz/passman/internal/store"
	"github.com/urfave/cli/v2"
)

type credentials struct {
	Source   string
	Password string
}

func authenticate() error {
	pass, err := password.Prompt("Password")
	if err != nil {
		return err
	}

	var root string
	err = store.Get("root", &root)
	if err != nil {
		return err
	}

	if pass != root {
		return errors.New("passwords do not match")
	}

	return nil
}

var App = &cli.App{
	Name:      "passman",
	Version:   "1.0.0",
	Usage:     "generate and manage your passwords on your computer",
	UsageText: "passman [COMMANDS] [FLAGS] [ARGS]",
	Copyright: "© 2022 Harun Sasmaz",
	Authors: []*cli.Author{
		{
			Name:  "Harun Sasmaz",
			Email: "me@harunsasmaz.com",
		},
	},
	HideHelpCommand:      true,
	EnableBashCompletion: true,
	CustomAppHelpTemplate: fmt.Sprintf(`%s

WEBSITE: https://harunsasmaz.com

SUPPORT: me@harunsasmaz.com

`, cli.AppHelpTemplate),
	Commands: []*cli.Command{
		generate,
		get,
		create,
		update,
		deletes,
	},
	Before: func(context *cli.Context) error {
		var pass string
		err := store.Get("root", &pass)
		if errors.Is(err, store.ErrNotFound) || pass == "" {
			pass, err = password.Prompt("Enter a password for manager")
			if err != nil {
				return err
			}

			err = store.Put("root", pass)
			if err != nil {
				return errors.New("failed to set manager password")
			}
		}

		return nil
	},
}
