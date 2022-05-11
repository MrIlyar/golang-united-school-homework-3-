package homework

func average(input [15]float32) (result float32) {
	var sum float32

	for _, value := range input {
		sum += value
	}

	result = sum / float32(15)

	return result
}
