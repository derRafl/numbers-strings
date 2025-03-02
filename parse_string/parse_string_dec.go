package parse_string

import (
	"math"
	"slices"
)

// ParseStringDecimal erwartet einen String, der eine Dezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringDecimal(s string) int {
	// TODO
	c := []rune(s)
	slices.Reverse(c)

	result := 0
	for i, n := range c {
		x := ParseDigit(string(n))
		if x == -1 || x > 9 {
			return -1
		} else {
			result += x * int(math.Pow(10, float64(i)))
		}
	}
	return result
}
