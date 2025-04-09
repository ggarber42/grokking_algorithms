package dijkstras

import (
	"reflect"
	"testing"
)

func Test_DijkstrasSearch(t *testing.T) {

	t.Run("find the fastest path", func(t *testing.T) {
		t.Run("fist case", func(t *testing.T) {
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
		})

		t.Run("piano case", func(t *testing.T) {
			piano := WeightedGraph{Name: "piano"}
			drumSet := WeightedGraph{Name: "drumSet", Neighbours: map[int]WeightedGraph{10: piano}}
			bassGuitar := WeightedGraph{Name: "bassGuitar", Neighbours: map[int]WeightedGraph{20: piano}}
			rareLP := WeightedGraph{Name: "rareLP", Neighbours: map[int]WeightedGraph{15: bassGuitar, 20: drumSet}}
			poster := WeightedGraph{Name: "poster", Neighbours: map[int]WeightedGraph{30: bassGuitar, 35: drumSet}}
			book := WeightedGraph{Name: "book", Neighbours: map[int]WeightedGraph{0: poster, 5: rareLP}}

			want := map[string]int{
				"rareLP":     5,
				"drumSet": 25,
				"piano":      35,
			}
			got, _ := DijkstrasSearch(book, piano)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	})

}
