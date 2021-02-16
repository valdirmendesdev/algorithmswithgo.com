package module01

import (
	"strings"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var sb strings.Builder
	quocient := dec

	for quocient > 0 {
		sb.WriteByte(charset[quocient%base])
		quocient = quocient / base
	}
	return Reverse(sb.String())
}
