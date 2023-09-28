package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var res strings.Builder

	for i, val := range s {
		if unicode.IsDigit(val) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(val) && unicode.IsDigit(rune(s[i+1])) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(val) && int(val-'0') != 0 {
			res.WriteString(strings.Repeat(string(s[i-1]), int(val-'0')-1))
			continue
		}

		if unicode.IsDigit(val) && int(val-'0') == 0 {
			temp := res.String()
			res.Reset()
			res.WriteString(temp[0 : len(temp)-1])
			continue
		}

		if unicode.IsLetter(val) {
			res.WriteString(string(val))
		}
	}

	return res.String(), nil
}
