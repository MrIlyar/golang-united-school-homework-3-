package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, 0, cap(input))

	for index := len(input) - 1; index >= 0; index-- {
		result = append(result, input[index])
	}

	return result
}
