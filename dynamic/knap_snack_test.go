package dynamic

import (
	"reflect"
	"testing"
)

func Test_KnpaSnack(t *testing.T) {

	knapSnackConstrain := 6

	items := []Item{
		{name: "Water", weight: 3, value: 10},
		{name: "Book", weight: 1, value: 3},
		{name: "Food", weight: 2, value: 9},
		{name: "Jacket", weight: 2, value: 5},
		{name: "Camera", weight: 1, value: 6},
	}

	expected := []Item{
		{name: "Water", weight: 3, value: 10},
		{name: "Food", weight: 2, value: 9},
		{name: "Camera", weight: 1, value: 6},
	}
	actual := KnapSnackSolve(items, knapSnackConstrain)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

}
