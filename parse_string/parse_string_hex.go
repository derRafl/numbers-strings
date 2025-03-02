package parse_string

import (
	"math"
	"slices"
)

// ParseStringDecimal erwartet einen String, der eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringHexadecimal(s string) int {
	// TODO
	c := []rune(s)
	slices.Reverse(c)

	result := 0
	for i, n := range c {
		x := ParseDigit(string(n))
		if x == -1 {
			return -1
		} else {
			result += x * int(math.Pow(16, float64(i)))
		}
	}
	return result
}
