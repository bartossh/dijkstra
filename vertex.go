package djikstra

import "gonum.org/v1/gonum/mat"

// vertex describes position in multidimensional space
type vertex struct {
	key         int
	position    []float64
	connections []int
}

// NewVertex creates instance of vertex implementing vertexer interface
func NewVertex(key int, position []float64, connections []int) *vertex {
	return &vertex{key, position, connections}
}

func (v vertex) GetKey() int {
	return v.key
}

func (v vertex) GetPosition() []float64 {
	return v.position
}

func (v vertex) GetDistance(np Positioner) float64 {
	p := np.GetPosition()
	vec1 := mat.NewVecDense(len(v.position), v.position)
	vec2 := mat.NewVecDense(len(p), p)
	rv := new(mat.VecDense)
	rv.SubVec(vec1.TVec(), vec2.TVec())
	return rv.Norm(2)
}

func (v vertex) GetConnections() []int {
	return v.connections
}
