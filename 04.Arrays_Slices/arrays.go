package arrays

func Sum(numbers []int) (sumation int) {
	for _, number := range numbers {
		sumation += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int{
	numbersLen := len(numbersToSum)
	sums := make([]int, numbersLen)
	for i, currentArray := range numbersToSum {
		sums[i] = Sum(currentArray)
	}

	return sums

}
