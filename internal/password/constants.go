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
