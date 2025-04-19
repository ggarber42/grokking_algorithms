package greedy

import (
	"testing"
)

func Test_SetCovering(t *testing.T) {
	t.Run("return the stations", func(t *testing.T) {
		statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

		stations := make(map[string][]string)
		stations["kone"] = []string{"id", "nv", "ut"}
		stations["ktwo"] = []string{"wa", "mt", "id"}
		stations["kthree"] = []string{"or", "nv", "ca"}
		stations["kfour"] = []string{"nv", "ut"}
		stations["kfive"] = []string{"ca", "az"}


		expected := []string{"ktwo", "kthree", "kone", "kfive"}
		actual := SetCover(stations, statesNeeded)
		if len(expected) != len(actual) {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	})
}
