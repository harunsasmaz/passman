package password

import "errors"

func generate(options *Options) (string, error) {
	return "", nil
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
