package djikstra

// PathVertex is a vertex on the path with a copy of his parent
type PathVertex struct {
	Parent, Actual Vertexer
}

// Path represents calculated shortest path between to nodes on the  graph
type Path struct {
	Vertexes      []*PathVertex
	TotalDistance float64
}
