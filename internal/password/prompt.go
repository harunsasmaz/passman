package password

import (
	"errors"
	"fmt"
	"golang.org/x/term"
	"os"
	"syscall"
)

func Prompt(label string) (string, error) {
	fmt.Fprint(os.Stderr, label+": ")
	b, err := term.ReadPassword(syscall.Stdin)
	fmt.Println()
	if err != nil {
		return "", errors.New("error on password prompt")
	}
	if string(b) == "" {
		return "", errors.New("password cannot be empty")
	}
	return string(b), nil
}
