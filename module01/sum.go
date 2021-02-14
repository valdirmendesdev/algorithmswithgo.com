package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	//Structured programming solution
	//total := 0
	//for _, num := range numbers {
	//	total += num
	//}
	//return total

	//Solution using recursion
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}
