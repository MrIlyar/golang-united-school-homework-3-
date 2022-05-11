package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	result = make([]string, 0, len(input))
	var keys = make([]int, 0, len(input))

	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, value := range keys {
		result = append(result, input[value])
	}

	return result
}
