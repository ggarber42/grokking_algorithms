package selectio_sort

import (
	"reflect"
	"testing"
)

func Test_SelectionSort(t *testing.T) {
	input := []int{5, 3, 6, 2, 10}
	got := SelectionSort(input)
	want := []int{2, 3, 5, 6, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}
