package greedy

func SetCover(stations map[string][]string, statesNeeded []string) []string {
	var result []string

	for len(statesNeeded) > 0 {
		var bestStation string
		statesCovered := make(map[string]bool)
		for station, states := range stations {
			covered := make(map[string]bool)
			for _, state := range states {
				if sliceContains[string](state, statesNeeded) {
					covered[state] = true
				}
			}

			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}

		for state := range statesCovered {
			removeElement(state, &statesNeeded)
		}
		result = append(result, bestStation)
	}
	return result
}

func removeElement[T comparable](value T, items *[]T) {
	s := *items
	for i := 0; i < len(s); {
		if s[i] == value {
			s = append(s[:i], s[i+1:]...)
		} else {
			i++
		}
	}
	*items = s
}

func sliceContains[T comparable](value T, items []T) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}
