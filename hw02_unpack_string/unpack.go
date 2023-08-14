package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder
	var prev rune
	//isEscaping := false

	for _, r := range str {
		/*
			//todo: доработать
			if r == '\\' && !isEscaping {
					isEscaping = true
					continue
				}

				if isEscaping {
					if unicode.IsDigit(r) || r == '\\' {
						result.WriteRune(r)
						isEscaping = false
						continue
					} else {
						return "", ErrInvalidString
					}
				}
		*/
		if unicode.IsDigit(r) {
			if prev == 0 {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(string(r))
			result.WriteString(strings.Repeat(string(prev), count))
			prev = 0
		} else if prev != 0 {
			result.WriteRune(prev)
			prev = r
		} else {
			prev = r
		}
	}

	if prev != 0 {
		result.WriteRune(prev)
	}

	return result.String(), nil
}
