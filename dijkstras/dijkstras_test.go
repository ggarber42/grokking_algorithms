package dijkstras

import (
	"reflect"
	"testing"
)

func Test_DijkstrasSearch(t *testing.T) {
	finish := WeightedGraph{Name: "finish"}
	A := WeightedGraph{Name: "A", Neighbours: map[int]WeightedGraph{1: finish}}
	B := WeightedGraph{Name: "B", Neighbours: map[int]WeightedGraph{3: A, 5: finish}}
	start := WeightedGraph{Name: "start", Neighbours: map[int]WeightedGraph{6: A, 2: B}}

	want := map[string]int{
		"A":      5,
		"B":      2,
		"finish": 6,
	}
	got, _ := DijkstrasSearch(start, finish)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
