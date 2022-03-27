package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

type credentials struct {
	Source   string
	Password string
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
}
