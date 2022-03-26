package djikstra

import (
	"errors"
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// Vertex represents graph node position in multidimensional euclidean space and all its connections
type Vertex struct {
	Position    *Position
	Connections []int
}

// Vertexes is a collection of Vertex entities
type Vertexes struct {
	inner []*Vertex
}

func NewVertexes(vertexes []*Vertex) *Vertexes {
	return &Vertexes{inner: vertexes}
}

// GetPositionByIdx returns vertex position by idx in Vertexes slice
func (v Vertexes) GetPositionByIdx(idx int) *Position { return v.inner[idx].Position }

// GetPositionByKey returns vertex position by its position key
func (v Vertexes) GetPositionByKey(key int) *Position {
	for _, v := range v.inner {
		if v.Position.key == key {
			return v.Position
		}
	}
	return nil
}

// PathNode represents next node in path with his parent
type PathNode struct {
	parent, actual *node
}

type Path struct {
	Visited       map[int]*node
	Path          map[*PathNode]struct{}
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

func newStartNode(vertex *Position) *node {
	neighbours := make(map[*node]struct{})
	return &node{vertex: vertex, neighbours: neighbours}
}

func (n *node) setStart() {
	n.total = 0
}

func (n *node) putNeighbour(nn *node) bool {
	_, ok := n.neighbours[nn]
	n.neighbours[nn] = struct{}{}
	return ok
}

func (n *node) getNeighboursDistance(nn *node) float64 {
	_, ok := n.neighbours[nn]
	if !ok {
		return math.MaxFloat64
	}
	return n.vertex.getDistance(nn.vertex)
}

type graph struct {
	visited, unvisited map[int]*node
}

// NewGraph creates new graph for dijkstra the shortest path calculation
func NewGraph(vertexes *Vertexes) *graph {
	kv := make(map[int]*Vertex)
	unvisited := make(map[int]*node)

	for _, v := range vertexes.inner {
		kv[v.Position.key] = v
		nd := newStartNode(v.Position)
		nd.total = math.MaxFloat64
		unvisited[v.Position.key] = nd
	}

	for k, v := range kv {
		for _, nk := range v.Connections {
			nd := unvisited[k]
			nd.neighbours[unvisited[nk]] = struct{}{}
			unvisited[k] = nd
		}
	}

	return &graph{unvisited: unvisited, visited: make(map[int]*node)}
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
		return Path{
			Visited:       map[int]*node{st.vertex.key: st},
			Path:          make(map[*PathNode]struct{}),
			TotalDistance: 0,
		}, nil
	}
	st.total = 0
	res := Path{
		Path: make(map[*PathNode]struct{}),
	}

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
		g.visited[act.vertex.key] = act
		minDist := math.MaxFloat64
		for _, n := range g.unvisited {
			if n.total < minDist {
				minDist = n.total
				act = n
			}
		}
		if act.vertex.key == fn.vertex.key {
			delete(g.unvisited, act.vertex.key)
			g.visited[act.vertex.key] = act
			res.TotalDistance = act.total
			res.Visited = g.visited
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
		res.Path[pn] = struct{}{}
	}

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
