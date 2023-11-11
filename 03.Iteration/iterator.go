package iterator

func Repeat(char string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated = repeated + char
	}

	return
}
