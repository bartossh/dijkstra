package djikstra

// Locator locates vertex on the graph by unique key
type Locator interface {
	GetKey() int
}

// Positioner describes vertex position in multidimensional space
type Positioner interface {
	GetPosition() []float64
}

// Ruler describes distance between vertexes
type Ruler interface {
	GetDistance(p Positioner) float64
}

// Connector describes vertex connections
type Connector interface {
	GetConnections() []int
}

// Vertexer allows to get all information needed to graph the shortest path
type Vertexer interface {
	Connector
	Ruler
	Positioner
	Locator
}
