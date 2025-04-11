package common

import (
	"reflect"
	"testing"
)

func Test_Set(t *testing.T) {
	t.Run("do not repeat values", func(t *testing.T) {
		valueA := 1
		valueB := 1

		set := NewSet[int]()
		set.Add(valueA)
		set.Add(valueB)

		if set.Size() > 1 {
			t.Errorf("it should not repeat values")
		}

		if !reflect.DeepEqual(set.Values(), []int{1}) {
			t.Errorf("Values should be {1}")
		}

		if !set.Contains(1){
			t.Errorf("should contain 1")
		}
	})
}
