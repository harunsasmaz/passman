package password

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"

	p "github.com/sethvargo/go-password/password"
)

func generate(options *Options) (string, error) {
	return p.Generate(options.Length, options.NumDigits, options.NumSymbols, !options.Uppercase, !options.NoRepeat)
}

func Generate(level Level) (string, error) {
	var options *Options
	switch level {
	case LevelEasy:
		options = &easyOptions
	case LevelMid:
		options = &midOptions
	case LevelHard:
		options = &hardOptions
	default:
		return "", errors.New("undefined password strength level")
	}

	return generate(options)
}

func GenerateWithOptions(options *Options) (string, error) {
	return generate(options)
}

func Prompt(label string) (string, error) {
	fmt.Fprint(os.Stderr, label+": ")
	b, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", errors.New("error on password prompt")
	}
	if string(b) == "" {
		return "", errors.New("password cannot be empty")
	}
	return string(b), nil
}
