package raindrops

import (
	"math"
	"strconv"
)

func Convert(num int) string {
	numf := float64(num)
	mod3 := math.Mod(numf, 3.0)
	mod5 := math.Mod(numf, 5.0)
	mod7 := math.Mod(numf, 7.0)
	var result = ""
	if (mod3 == 0) {
		result += "Pling"
	}
	if (mod5 == 0)  {
		result += "Plang"
	}
	if (mod7 == 0) {
		result += "Plong"
	}
	if (len(result) == 0) {
		return strconv.Itoa(num)
	} else {
		return result
	}
}


