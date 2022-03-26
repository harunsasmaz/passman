package main

import (
	"github.com/harunsasmaz/password-manager/internal/cli"
	"log"
	"os"
)

func main() {

	err := cli.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
