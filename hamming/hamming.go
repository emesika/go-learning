package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	var err error
	chain1 := []rune(a)
	chain2 := []rune(b)
	dist := 0
	err = nil
	if (len(a) != len(b)) {
		err = errors.New("Parameters are not in the same length")
		return dist, err
	}
	for i := 0; i < len(a) ; i++ {
		if (chain1[i] != chain2[i]) {
			dist++
		}
	}
	return dist, err
}
