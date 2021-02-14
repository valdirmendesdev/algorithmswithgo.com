package module01

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	//reversed := make([]rune, len(word))
	//for i := len(word) - 1; i > -1; i-- {
	//	reversed[len(word) - i - 1] = rune(word[i])
	//}
	//return string(reversed)

	//Using String Builder
	var sb strings.Builder
	for i := len(word)-1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}
	return sb.String()
}
