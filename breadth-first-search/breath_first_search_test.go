package breadth_first_search

import (
	"errors"
	"testing"
)

func Test_Breadth_First_Search(t *testing.T) {

	t.Run("returns the target in the shortest path if", func(t *testing.T) {

		t.Run("the shortest path is last neighbour", func(t *testing.T) {

			Target := Graph[string]{Name: "Target"}
			Misha := Graph[string]{Name: "Misha", Neighbours: []Graph[string]{Target}}
			Amy := Graph[string]{Name: "Amy", Neighbours: []Graph[string]{Misha}}
			Tom := Graph[string]{Name: "Tom", Neighbours: []Graph[string]{Amy}}
			Maia := Graph[string]{Name: "Maia", Neighbours: []Graph[string]{Target}}
			Jade := Graph[string]{Name: "Jade", Neighbours: []Graph[string]{Tom, Maia}}
			Alice := Graph[string]{Name: "Alice", Neighbours: []Graph[string]{Jade}}
			Bob := Graph[string]{Name: "Bob", Neighbours: []Graph[string]{Alice}}
			Anuj := Graph[string]{Name: "Anuj", Neighbours: []Graph[string]{Bob}}

			stepsWant := 7
			got, stepsGot, err := BreadthFirstSearch[string](Anuj, Target)
			want := Target

			if got.Name != want.Name {
				t.Errorf("got %q, want %q", got.Name, want.Name)
			}

			if stepsGot != stepsWant {
				t.Errorf("got %d, want %d", stepsGot, stepsWant)
			}

			if err != nil {
				t.Errorf("expected err as nil")
			}
		})

		t.Run("the shortest path is first neighbour", func(t *testing.T) {

			Target := Graph[string]{Name: "Target"}
			Misha := Graph[string]{Name: "Misha", Neighbours: []Graph[string]{Target}}
			Amy := Graph[string]{Name: "Amy", Neighbours: []Graph[string]{Misha}}
			Tom := Graph[string]{Name: "Tom", Neighbours: []Graph[string]{Amy}}
			Maia := Graph[string]{Name: "Maia", Neighbours: []Graph[string]{Target}}
			Jade := Graph[string]{Name: "Jade", Neighbours: []Graph[string]{Maia, Tom}}
			Alice := Graph[string]{Name: "Alice", Neighbours: []Graph[string]{Jade}}
			Bob := Graph[string]{Name: "Bob", Neighbours: []Graph[string]{Alice}}
			Anuj := Graph[string]{Name: "Anuj", Neighbours: []Graph[string]{Bob}}

			stepsWant := 6
			got, stepsGot, err := BreadthFirstSearch[string](Anuj, Target)
			want := Target

			if got.Name != want.Name {
				t.Errorf("got %q, want %q", got.Name, want.Name)
			}

			if stepsGot != stepsWant {
				t.Errorf("got %d, want %d", stepsGot, stepsWant)
			}

			if err != nil {
				t.Errorf("expected err as nil")
			}
		})

	})

	t.Run("return the target", func(t *testing.T) {
		Target := Graph[string]{Name: "Target"}
		Alice := Graph[string]{Name: "Alice", Neighbours: []Graph[string]{Target}}
		Bob := Graph[string]{Name: "Bob", Neighbours: []Graph[string]{Alice}}
		Anuj := Graph[string]{Name: "Anuj", Neighbours: []Graph[string]{Bob}}

		stepsWant := 3
		got, stepsGot, err := BreadthFirstSearch[string](Anuj, Target)
		want := Target

		if got.Name != want.Name {
			t.Errorf("got %q, want %q", got.Name, want.Name)
		}

		if stepsGot != stepsWant {
			t.Errorf("got %d, want %d", stepsGot, stepsWant)
		}

		if err != nil {
			t.Errorf("expected err as nil")
		}

	})

	t.Run("return error if no target is found", func(t *testing.T) {
		Target := Graph[string]{Name: "Target"}
		Alice := Graph[string]{Name: "Alice"}
		Bob := Graph[string]{Name: "Bob", Neighbours: []Graph[string]{Alice}}
		Anuj := Graph[string]{Name: "Anuj", Neighbours: []Graph[string]{Bob}}

		_, _, err := BreadthFirstSearch[string](Anuj, Target)
		expectedError := ErrTargetNotFound

		if !errors.Is(err, expectedError) {
			t.Errorf("got %q, want %q", err.Error(), expectedError.Error())
		}
	})
}
