package mathutil

import "math"

// Round float to 2 decimal places
func RoundFloatToTwoDecimalPlaces(f float64) float64 {
	return math.Floor(f*100) / 100
}
