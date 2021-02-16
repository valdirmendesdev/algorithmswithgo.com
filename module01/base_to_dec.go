package module01

import (
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	multiplier := 1
	result := 0
	for _, character := range Reverse(value) {
		intValue := strings.IndexRune(charset, character)
		result += intValue * multiplier
		multiplier *= base
	}

	return result
}
