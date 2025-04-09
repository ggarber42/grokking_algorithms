package breadth_first_search

import (
	"reflect"
	"testing"
)

func Test_Graph(t *testing.T) {

	t.Run("return Graph's correct properties", func(t *testing.T) {
		graph := Graph[string]{Name: "Thomas", Neighbours: []Graph[string]{{Name: "Jade"}}}

		expectedName := "Thomas"
		actualName := graph.Name

		expectedGraph := []Graph[string]{{Name: "Jade"}}
		actualGraph := graph.Neighbours

		if actualName != expectedName {
			t.Errorf("got %s, want  %s", actualName, expectedName)
		}

		if !reflect.DeepEqual(actualGraph, expectedGraph) {
			t.Errorf("got %s, want  %s", actualGraph, expectedGraph)
		}

		if graph.Neighbours[0].Neighbours != nil {
			t.Errorf("got %v, want  %v", graph.Neighbours[0].Neighbours, nil)
		}

	})

}
