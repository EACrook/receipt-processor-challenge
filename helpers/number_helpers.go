package helpers

import (
	"math"
)

func isRoundNumber(total float64) bool {
	return total == math.Floor(total)
}

func isMultipleFloat(dividend float64, divisor float64) bool {
	remainder := math.Mod(dividend, divisor)
	return math.Abs(remainder) < 1e-9
}

func isMultipleInt(dividend int, divisor int) bool {
	return dividend%divisor == 0
}