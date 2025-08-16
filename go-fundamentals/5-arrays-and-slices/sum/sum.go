package sum

func Sum(numbers []int) int {
	output := 0
	for _, num := range numbers {
		output += num
	}
	return output
}

func SumAll(numbersToSum ...[]int) []int {
	output := []int{}
	for _, listOfNumbers := range numbersToSum {
		output = append(output, Sum(listOfNumbers))
	}

	return output
}
