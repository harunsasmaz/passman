package password

import (
	"errors"

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
