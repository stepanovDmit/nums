package nums

func Add(numbers ...int) int {
	sum := 0
	for _, value :=range numbers {
		sum += value
	}

	return sum
}

func Sub(i, j int) int {
	return i - j
}