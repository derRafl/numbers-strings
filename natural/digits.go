package natural

import (
	"slices"
)

func Digits(n int) []int {
	result := []int{}

	for n != 0 {
		r := n % 10
		n = n / 10

		result = append(result, r)
	}

	slices.Reverse(result)
	return result
}
