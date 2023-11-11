package arrays

func Sum(numbers []int) (sumation int) {
	for _, number := range numbers {
		sumation += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	numbersLen := len(numbersToSum)
	sums := make([]int, numbersLen)
	for i, currentArray := range numbersToSum {
		sums[i] = Sum(currentArray)
	}

	return sums
}

func SumAllTails(arrs ...[]int) []int {
	var sums []int

	for _, numbers := range arrs {
		sums = append(sums, Sum(numbers[1:]))
	}

	return sums
}
