package mathutil

import "math"

// Round float to 2 decimal places
func RoundNearest(x float64) float64 {
	return math.Round(x*100) / 100
}
