package common

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.data[value] = struct{}{}
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Contains(value T) bool {
	if _, exists := s.data[value]; exists {
		return true
	}
	return false
}

func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.data))
	for value := range s.data {
		values = append(values, value)
	}
	return values
}
