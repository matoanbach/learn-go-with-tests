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

func SumAllTails(numbersToSumTails ...[]int) []int {
	lengthOfNumbers := len(numbersToSumTails)
	output := make([]int, lengthOfNumbers)
	for i, listOfNumbers := range numbersToSumTails {
		if len(listOfNumbers) == 0 {
			output[i] = 0
		} else {
			output[i] = Sum(listOfNumbers[1:])
		}
	}

	return output
}
