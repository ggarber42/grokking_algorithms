package binarysearch

import (
	"fmt"
	"math"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	t.Run("should return the element", func(t *testing.T) {

		tests := []struct {
			input []int
			num   int
		}{
			{
				input: []int{1, 2, 3, 4, 5},
				num:   5,
			},
			{
				input: []int{1, 2, 3, 4, 5},
				num:   4,
			},
			{
				input: []int{1, 2, 3, 4, 5},
				num:   1,
			},
			{
				input: []int{1, 2, 3, 4, 5, 20 ,30 ,40, 50, 100},
				num:   1,
			},
		}

		for i, test := range tests {
			name := fmt.Sprintf("test %d", i)
			t.Run(name, func(t *testing.T) {
				sliceLen := len(test.input)
				maxGuesses := math.Ceil(math.Log2(float64(sliceLen)))
				got, numOfguesses := BinarySearch(test.input, test.num)
				assertCorrectValue(t, got, test.num)

				if numOfguesses > int(maxGuesses) {
					t.Errorf("num of guesses: %d, max guesses: %d", numOfguesses, int(maxGuesses))
				}
			})
		}

	})

	t.Run("sould return -1 if the number do not exist on the input", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		num := 5
		got, _ := BinarySearch(input, num)
		want := -1
		assertCorrectValue(t, got, want)
	})

}

func assertCorrectValue(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
