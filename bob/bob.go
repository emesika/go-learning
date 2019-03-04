package bob

import (
	"strings"
	"unicode"
)

func Hey(stmt string) string {
	stmt = strings.TrimSpace(stmt)

	switch {
	case stmt == "":
		return "Fine. Be that way!"
	case IsAllUpper(stmt) && IsQuestion(stmt):
		return "Calm down, I know what I'm doing!"
	case IsAllUpper(stmt):
		return "Whoa, chill out!"
	case IsQuestion(stmt):
		return "Sure."
	default:
		return "Whatever."
	}
}

func IsQuestion(stmt string) bool {
	return (stmt[len(stmt)-1] == '?')
}
func IsAllUpper(stmt string) bool {
	return IsAny(stmt, unicode.IsUpper) && !IsAny(stmt, unicode.IsLower)
}

func IsAny(items string, test func(rune) bool) bool {
	for _, item := range items {
		if test(item) {
			return true
		}
	}
	return false
}
