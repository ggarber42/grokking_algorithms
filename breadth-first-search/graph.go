package breadth_first_search

type Graph[T any] struct {
	Name T
	Neighbours []Graph[T]
}
