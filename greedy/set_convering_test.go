package greedy

import (
	"grokk/common"
	"testing"
)

func Test_SetCovering(t *testing.T) {
	t.Run("return the stations", func(t *testing.T) {
		stations := make(map[string][]string)
		stations["one"] = []string{"id", "nv", "ut"}
		stations["two"] = []string{"wa", "mt", "id"}
		stations["three"] = []string{"or", "nv", "ca"}
		stations["four"] = []string{"nv", "ut"}
		stations["five"] = []string{"ca", "az"}

		statesNeeded := common.NewSet[string]()
		for _, states := range stations {
			for _, state := range states {
				statesNeeded.Add(state)
			}
		}

		expected := []string{"two", "three", "one", "five"}
		actual := SetCover(stations, statesNeeded)
		if len(expected) != len(actual) {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	})
}
