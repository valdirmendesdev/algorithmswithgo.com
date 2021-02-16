package module01

import "strconv"

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

	result := ""
	quocient := dec

	for quocient > 0 {
		switch quocient % base {
		case 10:
			result = "A" + result
		case 11:
			result = "B" + result
		case 12:
			result = "C" + result
		case 13:
			result = "D" + result
		case 14:
			result = "E" + result
		case 15:
			result = "F" + result
		default:
			result = strconv.Itoa(quocient%base) + result
		}
		quocient = quocient / base
	}
	return result
}
