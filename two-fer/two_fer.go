package twofer

import (
	"fmt"
)

func ShareWith(s string) string {
	result := "you"
	if len(s) > 0 {
		result = s
	}
	return fmt.Sprintf("One for %s, one for me.", result)
}
