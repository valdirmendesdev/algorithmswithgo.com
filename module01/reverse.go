package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	reversed := ""
	for _, r := range word {
		reversed = string(r) + reversed
	}
	return reversed

	//Using String Builder
	//var sb strings.Builder
	//for i := len(word)-1; i >= 0; i-- {
	//	sb.WriteByte(word[i])
	//}
	//return sb.String()
}
