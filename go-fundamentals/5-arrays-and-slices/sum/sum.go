package sum

func Sum(numbers []int) int {
	output := 0
	for _, num := range numbers {
		output += num
	}
	return output
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)

	output := make([]int, lengthOfNumbers)
	for i, listOfNumbers := range numbersToSum {
		output[i] = Sum(listOfNumbers)
	}

	return output
}
