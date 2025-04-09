package dijkstras

import (
	common "grokk/common"
)

type Element struct {
	name   string
	weight int
	parent *WeightedGraph
}

func checkIfNodeExists(fasterPath map[string]Element, nodeName string) bool {
	for name := range fasterPath {
		if name == nodeName {
			return true
		}
	}
	return false
}

func populateMapPath(mapPath *map[string]Element, start WeightedGraph) error {
	queue := common.Queue[WeightedGraph]{}
	queue.Enqueue(start)
	for !queue.IsEmpty() {
		n, err := queue.Dequeue()
		if err != nil {
			return err
		}
		for weight, node := range n.Neighbours {
			exist := checkIfNodeExists(*mapPath, node.Name)
			if exist {
				element := (*mapPath)[node.Name]
				currentWeight := element.weight
				newWeight := n.Weight + weight
				if currentWeight > newWeight {
					node.Weight = newWeight
					(*mapPath)[node.Name] = Element{name: node.Name, weight: newWeight, parent: &n}
				}
			} else {
				newWeight := n.Weight + weight
				node.Weight = newWeight
				(*mapPath)[node.Name] = Element{name: node.Name, weight: newWeight}
			}
			queue.Enqueue(node) // adds the node with the weight of the current
		}
	}
	return nil
}

func populateFasterPath(finish WeightedGraph, mapPath map[string]Element, fasterPath *map[string]int) error {
	queue := common.Queue[Element]{}
	queue.Enqueue(mapPath[finish.Name])

	for !queue.IsEmpty() {
		el, err := queue.Dequeue()
		if err != nil {
			return err
		}

		(*fasterPath)[el.name] = el.weight
		if el.parent != nil {
			if nextEl, ok := mapPath[el.parent.Name]; ok {
				queue.Enqueue(nextEl)
			}
		}

	}
	return nil
}

func DijkstrasSearch(start, finish WeightedGraph) (map[string]int, error) {
	mapPath := make(map[string]Element)
	err := populateMapPath(&mapPath, start)
	if err != nil {
		return nil, err
	}

	fasterPath := make(map[string]int)
	err = populateFasterPath(finish, mapPath, &fasterPath)
	if err != nil {
		return nil, err
	}
	return fasterPath, nil
}
