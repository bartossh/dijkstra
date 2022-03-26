package djikstra

import (
	"errors"
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// Vertex represents graph node position in multidimensional euclidean space and all its connections
type Vertex struct {
	position    *Position
	connections []int
}

// NewVertex is constructor creating pointer to new instance of Vertex
func NewVertex(p *Position, c []int) *Vertex {
	return &Vertex{
		position:    p,
		connections: c,
	}
}

// Position returns vertex position
func (v Vertex) Position() *Position {
	return v.position
}

// Connections returns vertex connections
func (v Vertex) Connections() []int {
	return v.connections
}

// Vertexes is a collection of Vertex entities
type Vertexes struct {
	inner []*Vertex
}

// NewVertexes creates pointer to Vertexes that has an inner slice of pointer to future graph vertex points
func NewVertexes(vertexes []*Vertex) *Vertexes {
	return &Vertexes{inner: vertexes}
}

// GetPositionByIdx returns vertex position by idx in Vertexes slice
func (v Vertexes) GetPositionByIdx(idx int) *Position { return v.inner[idx].position }

// GetPositionByKey returns vertex position by its position key
func (v Vertexes) GetPositionByKey(key int) *Position {
	for _, v := range v.inner {
		if v.position.key == key {
			return v.position
		}
	}
	return nil
}

// PathNode represents node with his parent node
type PathNode struct {
	parent, actual *node
}

// Path represents calculated shortest path between to nodes on the  graph
type Path struct {
	Nodes         []*PathNode
	TotalDistance float64
}

// Position describes position in multidimensional space
type Position struct {
	key    int
	vector *mat.VecDense
}

// NewPosition creates a pointer receiver instance describing vertex Position on connections graph
func NewPosition(key int, vector *mat.VecDense) *Position {
	return &Position{key, vector}
}

func (p Position) Key() int {
	return p.key
}

func (p Position) Vector() *mat.VecDense {
	return p.vector
}

func (p Position) getDistance(np *Position) float64 {
	rv := new(mat.VecDense)
	rv.SubVec(p.vector.TVec(), np.vector.TVec())
	return rv.Norm(2)
}

type node struct {
	total      float64
	vertex     *Position
	neighbours map[*node]struct{}
}

func (n *node) GetNodeKeyAndPosition() *Position {
	return n.vertex
}

func newStartNode(vertex *Position) *node {
	neighbours := make(map[*node]struct{})
	return &node{vertex: vertex, neighbours: neighbours}
}

func (n *node) getNeighboursDistance(nn *node) float64 {
	_, ok := n.neighbours[nn]
	if !ok {
		return math.MaxFloat64
	}
	return n.vertex.getDistance(nn.vertex)
}

type graph struct {
	vertexes  []*Vertex
	unvisited map[int]*node
}

// NewGraph creates new graph for dijkstra the shortest path calculation
func NewGraph(vertexes *Vertexes) *graph {
	graphVertexes := make([]*Vertex, 0, len(vertexes.inner))
	for _, v := range vertexes.inner {
		vv := *v
		graphVertexes = append(graphVertexes, &vv)
	}
	return &graph{unvisited: createUnvisited(vertexes.inner), vertexes: graphVertexes}
}

// CalculateResultGraphFromPosition provides Path of the shortest path calculation and error if path has no solution
func (g *graph) CalculateResultGraphFromPosition(a, b *Position) (Path, error) {
	st, fn := g.findStarFinishNodeByPosition(a, b)
	if st == nil {
		return Path{}, errors.New("cannot find start node")
	}
	if fn == nil {
		return Path{}, errors.New("cannot find finish node")
	}
	return g.calcResultGraph(st, fn)
}

// CalculateResultGraphFromKeys provides Path of the shortest path calculation and error if path has no solution
func (g *graph) CalculateResultGraphFromKeys(a, b int) (Path, error) {
	st, fn := g.findStartFinnishNodeByTheKey(a, b)
	if st == nil {
		return Path{}, errors.New("cannot find start node")
	}
	if fn == nil {
		return Path{}, errors.New("cannot find finish node")
	}
	return g.calcResultGraph(st, fn)
}

func (g *graph) calcResultGraph(st, fn *node) (Path, error) {
	if st.getNeighboursDistance(fn) == 0 {
		return Path{}, nil
	}
	st.total = 0
	res := Path{}

	act := st
Loop:
	for {
		for n := range act.neighbours {
			dist := act.getNeighboursDistance(n)
			dist = dist + act.total
			if dist < n.total {
				n.total = dist
			}
		}
		delete(g.unvisited, act.vertex.key)
		minDist := math.MaxFloat64
		for _, n := range g.unvisited {
			if n.total < minDist {
				minDist = n.total
				act = n
			}
		}
		if act.vertex.key == fn.vertex.key {
			delete(g.unvisited, act.vertex.key)
			res.TotalDistance = act.total
			break Loop
		}
		if minDist == math.MaxFloat64 {
			return res, fmt.Errorf("there is no connection between nodes of key %#v and %#v", st.vertex.key, fn.vertex.key)
		}
	}
	act = fn
	var parent *node
	for {
		if act.vertex.key == st.vertex.key {
			break
		}
		minDist := math.MaxFloat64
		parent = act
		for n := range act.neighbours {
			if n.total < minDist {
				minDist = n.total
				act = n
			}
		}
		pn := &PathNode{
			parent: parent,
			actual: act,
		}
		res.Nodes = append(res.Nodes, pn)
	}
	g.unvisited = createUnvisited(g.vertexes)

	return res, nil
}

func (g *graph) findStarFinishNodeByPosition(a, b *Position) (*node, *node) {
	var start, finish *node
Loop:
	for _, n := range g.unvisited {
		n.total = math.MaxFloat64
		if n.vertex.getDistance(a) == 0 {
			n.total = 0
			start = n
		}
		if n.vertex.getDistance(b) == 0 {
			finish = n
		}
		if start != nil && finish != nil {
			break Loop
		}
	}
	return start, finish
}

func (g *graph) findStartFinnishNodeByTheKey(a, b int) (*node, *node) {
	na := g.unvisited[a]
	nb := g.unvisited[b]
	return na, nb
}

func createUnvisited(vrx []*Vertex) map[int]*node {
	kv := make(map[int]*Vertex)
	unvisited := make(map[int]*node)

	for _, v := range vrx {
		kv[v.position.key] = v
		nd := newStartNode(v.position)
		nd.total = math.MaxFloat64
		unvisited[v.position.key] = nd
	}

	for k, v := range kv {
		for _, nk := range v.connections {
			nd := unvisited[k]
			nd.neighbours[unvisited[nk]] = struct{}{}
			unvisited[k] = nd
		}
	}
	return unvisited
}
