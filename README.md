# dijkstra

Go implementation of dijkstra algorithm.

### Key features

- Multidimensional vertex position.
- Allows calculation of different paths based on once created graph, no need to reset/reload graph.
- Focuses on speed, this is first iteration so future improvements are possible.
- Uses abstraction to calculate distance, position and connection relations, so you may use your own implementation of the vertex
- Provides implementation of Vertexer based on [gonum matrix](gonum.org/v1/gonum/mat)

### Abstraction 

`Vertexer` is a collective interface required to calculate path. Look in to the documentation to see how to implement it or use
implementation provided by the package calling `NewVertex(key, possition, connections)`

### Usage example

```go
package main

import "github.com/bartossh/dijkstra"

vertexes := []dijkstra.Vertexer{
    dijkstra.NewVertex(0, []float64{0, 0}, []int{1}),
    dijkstra.NewVertex(1, []float64{1, 0}, []int{0, 2}),
    dijkstra.NewVertex(2, []float64{2, 0}, []int{1, 3}),
    dijkstra.NewVertex(3, []float64{3, 0}, []int{2, 4}),
    dijkstra.NewVertex(4, []float64{4, 0}, []int{3, 5}),
    dijkstra.NewVertex(5, []float64{5, 0}, []int{4}),
}

graph := dijkstra.NewGraph(vertexes)

path, err := graph.CalculateResultGraph(vertexes[0], vertexes[5])
if err != nil {
	// do something with error
}
// use path

```

