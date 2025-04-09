package selectio_sort

func findIndexOfSmallest(input []int) int {
	index := 0
	for i := range input {
		if input[i] < input[index] {
			index = i
		}
	}
	return index
}

func SelectionSort(input []int) []int {
	output := make([]int, 0)
	inputSize := len(input)
	for i := 0; i < inputSize; i++ {
		index := findIndexOfSmallest(input)
		output = append(output, input[index])

		input = append(input[:index], input[index+1:]...)
	}
	return output
}
