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
	for _, r := range str {
		switch {
		case unicode.IsDigit(r):
			if prev == 0 {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(r))
			result.WriteString(strings.Repeat(string(prev), count))
			prev = 0
		case prev != 0:
			result.WriteRune(prev)
			prev = r
		default:
			prev = r
		}
	}
	if prev != 0 {
		result.WriteRune(prev)
	}
	return result.String(), nil
}
