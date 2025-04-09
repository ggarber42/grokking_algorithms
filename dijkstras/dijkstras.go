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

func DijkstrasSearch(start, finish WeightedGraph) (map[string]int, error) {
	queue := common.Queue[WeightedGraph]{}
	mapPath := make(map[string]Element)

	queue.Enqueue(start)
	for !queue.IsEmpty() {
		n, err := queue.Dequeue()
		if err != nil {
			return nil, err
		}
		for weight, node := range n.Neighbours {
			exist := checkIfNodeExists(mapPath, node.Name)
			if exist {
				element := mapPath[node.Name]
				currentWeight := element.weight
				newWeight := n.Weight + weight
				if currentWeight > newWeight {
					node.Weight = newWeight
					mapPath[node.Name] = Element{name: node.Name, weight: newWeight, parent: &n}
				}
			} else {
				newWeight := n.Weight + weight
				node.Weight = newWeight
				mapPath[node.Name] = Element{name: node.Name, weight: newWeight}
			}
			queue.Enqueue(node) // adds the node with the weight of the current
		}
	}

	elementQueue := common.Queue[Element]{}
	fasterPath := make(map[string]int)

	elementQueue.Enqueue(mapPath[finish.Name])
	for !elementQueue.IsEmpty() {
		el, err := elementQueue.Dequeue()
		if err != nil {
			return nil, err
		}

		fasterPath[el.name] = el.weight

		if el.parent != nil {
			if nextEl, ok := mapPath[el.parent.Name]; ok {
				elementQueue.Enqueue(nextEl)
			}
		}

	}

	// for k, v := range mapPath {
	// 	fasterPath[k] = v.weight
	// }

	return fasterPath, nil
}
