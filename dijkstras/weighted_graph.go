package dijkstras

type WeightedGraph struct {
	Name       string
	Neighbours map[int]WeightedGraph
	Weight     int
}
