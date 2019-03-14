package luhn

import (
	"strconv"
	"strings"
)

// Valid - Validates that the given string is legal
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	length := len(s)
	if length < 2 {
		return false
	}

	var digit, dupDigit, sum int
	var err error
	for i := 0; i < length; i++ {
		digit, err = strconv.Atoi(string(s[length-1-i]))
		if err != nil {
			return false
		}

		if i%2 == 0 {
			sum += digit
			continue
		}

		dupDigit = 2 * digit
		if dupDigit > 9 {
			dupDigit -= 9
		}
		sum += dupDigit
	}

	return sum%10 == 0
}
