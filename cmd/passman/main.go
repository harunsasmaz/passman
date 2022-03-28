package main

import (
	"fmt"
	"log"
	"os"

	"github.com/harunsasmaz/passman/internal/cli"
	"github.com/harunsasmaz/passman/internal/store"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("failed to read home directory of user")
	}

	path := fmt.Sprintf("%s/.passman", home)

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatalln("failed to initialize passman path")
	}

	dbPath := fmt.Sprintf("%s/store.db", path)
	err = store.Open(dbPath)
	if err != nil {
		log.Fatalln("failed to initialize local key-value store")
	}
	defer store.Close()

	err = cli.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
