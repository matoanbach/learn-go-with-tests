package sum

func Sum(numbers []int) int {
	output := 0
	for _, num := range numbers {
		output += num
	}
	return output
}

func SumAll() any {
	return ""
}
