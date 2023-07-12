package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/example/stringutil"
)

var ErrInvalidString = errors.New("invalid string")

/*
	func Unpacks(_ string) (string, error) {
		// Place your code here.
		return "", nil
	}
*/
func Unpack(input string) (string, error) {
	var result strings.Builder
	var count int
	input = stringutil.Reverse(input)
	for i, r := range input {
		if unicode.IsDigit(r) {
			if i == 0 || !unicode.IsLetter(rune(input[i-1])) {
				return "", ErrInvalidString
			}

			digit, _ := strconv.Atoi(string(r))
			count = count*10 + digit
		} else {
			if count > 0 {
				result.WriteString(strings.Repeat(string(r), count))
				count = 0
			} else {
				result.WriteRune(r)
			}
		}
	}

	return stringutil.Reverse(result.String()), nil
}
