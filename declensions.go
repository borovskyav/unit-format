package unit

import "math"

func declensions(value float64, nominative string, genitive string, plural string) string {
	absValue := math.Abs(value)
	intPart, fracPart := math.Modf(absValue)
	if fracPart != 0 {
		return plural
	}
	tenRemainder := int(intPart) % 10
	hundredRemainder := int(intPart) % 100
	if tenRemainder == 0 || tenRemainder >= 5 || (hundredRemainder >= 10 && hundredRemainder <= 20) {
		return plural
	}
	if tenRemainder == 1 {
		return nominative
	}
	return genitive
}
