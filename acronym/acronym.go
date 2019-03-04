package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	var abbrev string
	isFirstLetterOfWord := true
	s = strings.TrimSpace(s)
	for _, word := range s {
		if isFirstLetterOfWord {
			abbrev += strings.ToUpper(string(word))
			isFirstLetterOfWord = false
		} else if word == ' ' || word == '-' {
			isFirstLetterOfWord = true
		}
	}
	return abbrev
}
