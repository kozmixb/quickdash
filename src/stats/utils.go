package stats

import "math"

func convertTo2Decimals(num float64) float64 {
	multiplied := num * 100
	ceiled := math.Ceil(multiplied)

	return (ceiled / 100)
}
