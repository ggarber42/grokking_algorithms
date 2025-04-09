package dijkstras

import (
	common "grokk/common"
)

func checkIfNodeExists(fasterPath map[string]int, nodeName string) bool {
	for name := range fasterPath {
		if name == nodeName {
			return true
		}
	}
	return false
}

func DijkstrasSearch(start, finish WeightedGraph) (map[string]int, error) {
	queue := common.Queue[WeightedGraph]{}
	fasterPath := make(map[string]int)

	queue.Enqueue(start)
	for !queue.IsEmpty() {
		n, err := queue.Dequeue()
		if err != nil {
			return nil, err
		}
		for weight, node := range n.Neighbours {
			exist := checkIfNodeExists(fasterPath, node.Name)
			if exist {
				currentWeight := fasterPath[node.Name]
				newWeight := n.Weight + weight
				if currentWeight > newWeight {
					node.Weight = newWeight
					fasterPath[node.Name] = newWeight
				}
			} else {
				newWeight := n.Weight + weight
				node.Weight = newWeight
				fasterPath[node.Name] = newWeight
			}
			queue.Enqueue(node)
		}
	}

	return fasterPath, nil
}
