package arrays

func Sum(numbers []int) (sumation int) {
	for _, number := range numbers {
		sumation += number
	}
	return
}
