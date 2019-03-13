package grains

import (
	"errors"
	"math"
)

func Square(square int) (uint64, error) {
	// Compute only numbers in the chess board
	if !(square >= 1 && square <= 64) {
		return 0, errors.New("No such square number")
	}
	return uint64(math.Pow(2, float64(square-1))), nil
}

func Total() uint64 {
	// Can be done by looping and calling Square, but much more simple to use math facts here
	return uint64(math.Pow(2, float64(64))) - 1
}
