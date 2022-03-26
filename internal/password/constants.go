package password

import "errors"

type Level int

const (
	LevelUndefined Level = iota

	// LevelEasy is the options for a weak password.
	LevelEasy = 1

	// LevelMid is the options for a mid strength password.
	LevelMid = 2

	// LevelHard is the options for a strong password.
	LevelHard = 3
)

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted digits.
	Digits = "0123456789"

	// Symbols is the list of symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var (
	// ErrExceedsTotalLength is the error returned with the number of digits and
	// symbols is greater than the total length.
	ErrExceedsTotalLength = errors.New("number of digits and symbols must be less than total length")

	// ErrLettersExceedsAvailable is the error returned with the number of letters
	// exceeds the number of available letters and repeats are not allowed.
	ErrLettersExceedsAvailable = errors.New("number of letters exceeds available letters and repeats are not allowed")

	// ErrDigitsExceedsAvailable is the error returned with the number of digits
	// exceeds the number of available digits and repeats are not allowed.
	ErrDigitsExceedsAvailable = errors.New("number of digits exceeds available digits and repeats are not allowed")

	// ErrSymbolsExceedsAvailable is the error returned with the number of symbols
	// exceeds the number of available symbols and repeats are not allowed.
	ErrSymbolsExceedsAvailable = errors.New("number of symbols exceeds available symbols and repeats are not allowed")
)
