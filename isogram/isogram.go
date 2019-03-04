package isogram

import (
	"strings"
)

func IsIsogram(input string) bool {
	var table [26]bool
	chars := []rune(strings.ToUpper(input))
	for i := 0; i < len(chars); i++ {
		c := string(input[i])
		if c == "-" || c == " " {
			continue
		}
		// Use ASCII code to get position in table array
		if table[chars[i]-65] {
			return false
		}
		table[chars[i]-65] = true
	}
	return true
}
