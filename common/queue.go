package common

import "errors"

type Queue[T any] struct {
	values []T
}

func (q *Queue[T]) Enqueue(value T) {
	q.values = append(q.values, value)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if len(q.values) == 0 {
		return zero, errors.New("the queue is empty")
	}

	first := q.values[0]
	q.values = q.values[1:]

	return first, nil
}
