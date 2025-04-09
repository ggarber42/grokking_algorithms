package breadth_first_search

import (
	"fmt"
	common "grokk/common"
)

func BreadthFirstSearch[T comparable](graph Graph[T], target Graph[T]) (Graph[T], int, error) {
	var steps int
	q := common.Queue[Graph[T]]{}

	q.Enqueue(graph)
	visited := make(map[T]bool)

	for !q.IsEmpty() {

		value, err := q.Dequeue()
		
		
		if err != nil {
			return Graph[T]{}, steps, err
		}

		fmt.Println(value.Name)
		if value.Name == target.Name {
			return value, steps, nil
		}
		for _, neighbour := range value.Neighbours {
			if !visited[neighbour.Name] {
				q.Enqueue(neighbour)
				visited[neighbour.Name] = true
			}
		}
		steps += 1
	}

	return Graph[T]{}, steps, ErrTargetNotFound
}
