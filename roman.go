package roman

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrOutOfRange = errors.New("Arabic number out of range. Must be 1 to 3,999")
var ErrEmptyString = errors.New("Empty string is invalid Roman numerals")

// error type for encountering non-Roman digits
type errInvalidDigit struct {
	roman string
	i     int
	c     rune
}

func (err *errInvalidDigit) Error() string {
	return fmt.Sprintf(
		"Invalid Roman digit %s (pos %d in \"%s\")",
		strconv.QuoteRune(err.c), err.i, err.roman,
	)
}

// strings of Roman digits and their corresponding Arabic value
type pair struct {
	roman  string
	arabic int
}

// some helpful maps
var arabicFor = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}
var pairs = []pair{
	pair{"M", 1000},
	pair{"CM", 900},
	pair{"D", 500},
	pair{"CD", 400},
	pair{"C", 100},
	pair{"XC", 90},
	pair{"L", 50},
	pair{"XL", 40},
	pair{"X", 10},
	pair{"IX", 9},
	pair{"V", 5},
	pair{"IV", 4},
	pair{"I", 1},
}

func IsValid(r string) bool {
	_, err := Decode(r)
	return err == nil
}

func Encode(arabic int) (string, error) {
	if arabic < 1 || arabic > 3999 {
		return "", ErrOutOfRange
	}

	roman := ""
	for _, p := range pairs {
		for arabic >= p.arabic {
			arabic -= p.arabic
			roman += p.roman
		}

		if arabic == 0 {
			break
		}
	}
	return roman, nil
}

func Decode(roman string) (int, error) {
	if len(roman) == 0 {
		return 0, ErrEmptyString
	}
	roman = strings.ToUpper(roman) // arabicFor uses upper case letters

	previousDigit := 1000
	arabic := 0
	for i, c := range roman {
		digit, ok := arabicFor[c]
		if !ok {
			return 0, &errInvalidDigit{roman, i, c}
		}
		arabic += digit

		if previousDigit < digit {
			arabic -= 2 * previousDigit
		}
		previousDigit = digit
	}

	return arabic, nil
}
