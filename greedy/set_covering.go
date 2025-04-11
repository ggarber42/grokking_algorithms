package greedy

import "grokk/common"

func SetCover(stations map[string][]string, statesNeeded *common.Set[string]) []string {
	statesCovered := common.NewSet[string]()
	stationsNedded := common.NewSet[string]()

	for station, states := range stations {
		for _, state := range states {
			if !statesCovered.Contains(state) && statesNeeded.Contains(state) {
				stationsNedded.Add(station)
				statesCovered.Add(state)
			}
		}

	}
	return stationsNedded.Values()
}
