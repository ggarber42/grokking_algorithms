package binarysearch

func BinarySearch(input []int, num int) (int, int) {
	numOfguesses := 0
	low := 0
	high := len(input) - 1
	for {

		mid := (high + low) / 2

		if low > high {
			break
		}

		if input[mid] == num {
			return input[mid], numOfguesses
		}

		if input[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}

		numOfguesses += 1
	}

	return -1, numOfguesses
}
