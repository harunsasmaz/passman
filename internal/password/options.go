package password

type Options struct {
	// length indices the length of password, inclusive range is [1, 64].
	length int

	// numDigits indices the number of digits contained in a password.
	numDigits int

	// numSymbols indices the number of symbols contained in a password.
	numSymbols int

	// lowercase indices if password should contain lower case letters.
	lowercase bool

	// uppercase indices if password should contain upper case letters.
	uppercase bool

	// digits indices if password should contain digits.
	digits bool

	// symbols indices if password should contain symbols.
	symbols bool

	// repeat indices if password should contain repeated characters.
	repeat bool
}

var (

	// easyOptions is a pre-defined options for an easy level password.
	easyOptions = Options{
		length:     8,
		numDigits:  2,
		numSymbols: 0,
		lowercase:  true,
		uppercase:  false,
		digits:     true,
		symbols:    false,
		repeat:     false,
	}

	// midOptions is a pre-defined options for a mid level password.
	midOptions = Options{
		length:     16,
		numDigits:  4,
		numSymbols: 4,
		lowercase:  true,
		uppercase:  true,
		digits:     true,
		symbols:    true,
		repeat:     false,
	}

	// hardOptions is a pre-defined options for a hard level password.
	hardOptions = Options{
		length:     32,
		numDigits:  8,
		numSymbols: 8,
		lowercase:  true,
		uppercase:  true,
		digits:     true,
		symbols:    true,
		repeat:     false,
	}
)
