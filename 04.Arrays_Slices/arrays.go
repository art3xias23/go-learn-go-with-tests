package arrays

func Sum(numbers [5]int) (sumation int) {
	for i := 0; i < 5; i++ {
		sumation +=numbers[i]
	}
	return
}
