package password

type Level int

const (
	LevelUndefined Level = iota

	// LevelEasy is the options for a weak password.
	LevelEasy = 1

	// LevelMid is the options for a mid-strength password.
	LevelMid = 2

	// LevelHard is the options for a strong password.
	LevelHard = 3
)

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of Uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted digits.
	Digits = "0123456789"

	// Symbols is the list of symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)
