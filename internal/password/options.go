package password

type Options struct {
	// Length implies the Length of password, inclusive range is [1, 64].
	Length int

	// NumDigits implies the number of digits contained in a password.
	NumDigits int

	// NumSymbols implies the number of symbols contained in a password.
	NumSymbols int

	// Uppercase implies if password should contain upper case letters.
	Uppercase bool

	// NoRepeat implies if password should not contain repeated characters.
	NoRepeat bool
}

var (

	// easyOptions is a pre-defined options for an easy level password.
	easyOptions = Options{
		Length:     8,
		NumDigits:  2,
		NumSymbols: 2,
		Uppercase:  true,
		NoRepeat:   true,
	}

	// midOptions is a pre-defined options for a mid level password.
	midOptions = Options{
		Length:     16,
		NumDigits:  4,
		NumSymbols: 4,
		Uppercase:  true,
		NoRepeat:   true,
	}

	// hardOptions is a pre-defined options for a hard level password.
	hardOptions = Options{
		Length:     32,
		NumDigits:  8,
		NumSymbols: 8,
		Uppercase:  true,
		NoRepeat:   true,
	}
)
