package djikstra

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestDistanceCalculation(t *testing.T) {
	// given
	cases := []struct {
		title    string
		points   [2][]float64
		distance float64
	}{
		{
			title:    "2D distance calculation 1",
			points:   [2][]float64{{1, 1}, {2, 2}},
			distance: 1.4142135623730951,
		},
		{
			title:    "2D distance calculation 2",
			points:   [2][]float64{{0, 0}, {1, 1}},
			distance: 1.4142135623730951,
		},
		{
			title:    "2D distance calculation 3",
			points:   [2][]float64{{0, 1}, {0, 2}},
			distance: 1,
		},
		{
			title:    "2D distance calculation 4",
			points:   [2][]float64{{1, 0}, {0, 2}},
			distance: 2.23606797749979,
		},
		{
			title:    "2D distance calculation 5",
			points:   [2][]float64{{1, 0}, {1, 0}},
			distance: 0,
		},
		{
			title:    "2D distance calculation 6",
			points:   [2][]float64{{1, 0}, {0, 1}},
			distance: 1.4142135623730951,
		},
		{
			title:    "2D distance calculation 7",
			points:   [2][]float64{{3, 0}, {0, 4}},
			distance: 5,
		},
		{
			title:    "2D distance calculation 8",
			points:   [2][]float64{{0, 8}, {15, 0}},
			distance: 17,
		},
		{
			title:    "3D distance calculation 1",
			points:   [2][]float64{{0, 0, 0}, {1, 0, 1}},
			distance: 1.4142135623730951,
		},
		{
			title:    "3D distance calculation 2",
			points:   [2][]float64{{0, 0, 0}, {1, 1, 1}},
			distance: 1.7320508075688772,
		},
		{
			title:    "3D distance calculation 3",
			points:   [2][]float64{{0, 0, 0}, {0, 0, 1}},
			distance: 1,
		},
		{
			title:    "3D distance calculation 4",
			points:   [2][]float64{{0, 0, 0}, {10, 10, 10}},
			distance: 17.32050807568877,
		},
		{
			title:    "3D distance calculation 5",
			points:   [2][]float64{{2, 3, 6}, {0, 0, 0}},
			distance: 7,
		},
		{
			title:    "4D distance calculation 1",
			points:   [2][]float64{{2, 3, 6, 0}, {0, 0, 0, 0}},
			distance: 7,
		},
		{
			title:    "4D distance calculation 2",
			points:   [2][]float64{{3, 4, 0, 0}, {0, 0, 0, 0}},
			distance: 5,
		},
		{
			title:    "4D distance calculation 3",
			points:   [2][]float64{{1, 1, 1, 1}, {2, 2, 2, 2}},
			distance: 2,
		},
	}

	asr := assert.New(t)

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// when
			p1 := NewPosition(1, mat.NewVecDense(len(c.points[0]), c.points[0]))
			p2 := NewPosition(2, mat.NewVecDense(len(c.points[1]), c.points[1]))

			// then
			dist := p1.getDistance(p2)
			asr.Equal(c.distance, dist, fmt.Sprintf("Expected %v, got %v", c.distance, dist))
		})
	}
}

func TestNodesCreation(t *testing.T) {
	type connection struct {
		pos  []float64
		conn []int
	}

	// given
	cases := []struct {
		title string
		graph []connection
	}{
		{
			title: "small graph",
			graph: []connection{
				{
					pos:  []float64{0, 0, 0},
					conn: []int{1},
				},
				{
					pos:  []float64{0, 1, 0},
					conn: []int{0, 2},
				},
				{
					pos:  []float64{0, 0, 1},
					conn: []int{1, 3},
				},
				{
					pos:  []float64{1, 1, 0},
					conn: []int{2},
				},
			},
		},
		{
			title: "big graph",
			graph: []connection{
				{
					pos:  []float64{0, 0, 0},
					conn: []int{1},
				},
				{
					pos:  []float64{0, 1, 0},
					conn: []int{0, 2},
				},
				{
					pos:  []float64{0, 0, 1},
					conn: []int{1, 3},
				},
				{
					pos:  []float64{1, 1, 0},
					conn: []int{2, 4},
				},
				{
					pos:  []float64{2, 0, 0},
					conn: []int{3, 5},
				},
				{
					pos:  []float64{0, 2, 0},
					conn: []int{4, 6},
				},
				{
					pos:  []float64{0, 2, 1},
					conn: []int{5, 7},
				},
				{
					pos:  []float64{1, 2, 0},
					conn: []int{6, 8},
				},
				{
					pos:  []float64{3, 0, 0},
					conn: []int{7, 9},
				},
				{
					pos:  []float64{0, 3, 0},
					conn: []int{8, 10},
				},
				{
					pos:  []float64{0, 3, 1},
					conn: []int{9, 11},
				},
				{
					pos:  []float64{1, 3, 0},
					conn: []int{10},
				},
			},
		},
		{
			title: "complex graph",
			graph: []connection{
				{
					pos:  []float64{0, 0, 0},
					conn: []int{1, 10},
				},
				{
					pos:  []float64{0, 1, 0},
					conn: []int{0, 2, 11},
				},
				{
					pos:  []float64{0, 0, 1},
					conn: []int{1, 3, 5},
				},
				{
					pos:  []float64{1, 1, 0},
					conn: []int{2, 4, 7},
				},
				{
					pos:  []float64{2, 0, 0},
					conn: []int{3, 5, 9},
				},
				{
					pos:  []float64{0, 2, 0},
					conn: []int{4, 6, 5, 8},
				},
				{
					pos:  []float64{0, 2, 1},
					conn: []int{5, 7, 1, 11, 9, 4},
				},
				{
					pos:  []float64{1, 2, 0},
					conn: []int{6, 8, 9, 10},
				},
				{
					pos:  []float64{3, 0, 0},
					conn: []int{7, 9, 8, 6, 4, 2},
				},
				{
					pos:  []float64{0, 3, 0},
					conn: []int{8, 10, 9, 7, 6, 5, 4, 3, 2},
				},
				{
					pos:  []float64{0, 3, 1},
					conn: []int{9, 11, 0},
				},
				{
					pos:  []float64{1, 3, 0},
					conn: []int{10},
				},
			},
		},
	}

	asr := assert.New(t)

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// when
			nodes := make([]*node, 0, len(c.graph))
			for i, g := range c.graph {
				ps := NewPosition(i, mat.NewVecDense(len(g.pos), g.pos))
				n := newStartNode(ps)
				nodes = append(nodes, n)
			}
			for i, g := range c.graph {
				for _, idx := range g.conn {
					nodes[i].neighbours[nodes[idx]] = struct{}{}
				}
			}
			// then
			for i, n := range nodes {
				poss := make([]int, 0)
				for nn := range n.neighbours {
					poss = append(poss, nn.vertex.key)
				}
				sort.Ints(poss)
				conn := c.graph[i].conn
				sort.Ints(conn)
				asr.Equal(conn, poss, "should have the same number od nodes as neighbours", conn, poss)
			}
		})
	}
}

func TestVertexesGetPositionByKey(t *testing.T) {
	cases := []struct {
		title      string
		vertexes   *Vertexes
		key1, key2 int
		p1, p2     *Position
	}{
		{
			title: "small simple one way graph",
			vertexes: NewVertexes(
				[]*Vertex{
					NewVertex(NewPosition(5, mat.NewVecDense(2, []float64{5, 0})), []int{4}),
					NewVertex(NewPosition(0, mat.NewVecDense(2, []float64{0, 0})), []int{1}),
					NewVertex(NewPosition(1, mat.NewVecDense(2, []float64{1, 0})), []int{0, 2}),
					NewVertex(NewPosition(2, mat.NewVecDense(2, []float64{2, 0})), []int{1, 3}),
					NewVertex(NewPosition(4, mat.NewVecDense(2, []float64{4, 0})), []int{3, 5}),
					NewVertex(NewPosition(3, mat.NewVecDense(2, []float64{3, 0})), []int{2, 4}),
				}),
			key1: 0,
			key2: 5,
			p1:   NewPosition(0, mat.NewVecDense(2, []float64{0, 0})),
			p2:   NewPosition(5, mat.NewVecDense(2, []float64{5, 0})),
		},
		{
			title: "small simple graph shuffle",
			vertexes: NewVertexes(
				[]*Vertex{
					{
						position:    NewPosition(3, mat.NewVecDense(2, []float64{3, 0})),
						connections: []int{2, 4},
					},
					{
						position:    NewPosition(2, mat.NewVecDense(2, []float64{2, 0})),
						connections: []int{1, 3},
					},
					{
						position:    NewPosition(5, mat.NewVecDense(2, []float64{5, 0})),
						connections: []int{4},
					},
					{
						position:    NewPosition(0, mat.NewVecDense(2, []float64{0, 0})),
						connections: []int{1},
					},
					{
						position:    NewPosition(1, mat.NewVecDense(2, []float64{1, 0})),
						connections: []int{0, 2},
					},
					{
						position:    NewPosition(4, mat.NewVecDense(2, []float64{4, 0})),
						connections: []int{3, 5},
					},
				}),
			key1: 0,
			key2: 5,
			p1:   NewPosition(0, mat.NewVecDense(2, []float64{0, 0})),
			p2:   NewPosition(5, mat.NewVecDense(2, []float64{5, 0})),
		},
		{
			title: "bigger simple shuffle graph",
			vertexes: NewVertexes(
				[]*Vertex{
					{
						position:    NewPosition(3, mat.NewVecDense(2, []float64{3, 0})),
						connections: []int{2, 4},
					},
					{
						position:    NewPosition(2, mat.NewVecDense(2, []float64{2, 0})),
						connections: []int{1, 3},
					},
					{
						position:    NewPosition(5, mat.NewVecDense(2, []float64{5, 0})),
						connections: []int{4},
					},
					{
						position:    NewPosition(0, mat.NewVecDense(2, []float64{0, 0})),
						connections: []int{1},
					},
					{
						position:    NewPosition(1, mat.NewVecDense(2, []float64{1, 0})),
						connections: []int{0, 2},
					},
					{
						position:    NewPosition(4, mat.NewVecDense(2, []float64{4, 0})),
						connections: []int{3, 5},
					},
					{
						position:    NewPosition(7, mat.NewVecDense(2, []float64{9, 0})),
						connections: []int{3, 5},
					},
					{
						position:    NewPosition(6, mat.NewVecDense(2, []float64{2000, 0})),
						connections: []int{3, 5},
					},
					{
						position:    NewPosition(8, mat.NewVecDense(2, []float64{121, 0})),
						connections: []int{3, 5},
					},
				}),
			key1: 5,
			key2: 8,
			p1:   NewPosition(5, mat.NewVecDense(2, []float64{5, 0})),
			p2:   NewPosition(8, mat.NewVecDense(2, []float64{121, 0})),
		},
	}

	asr := assert.New(t)

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			p1 := c.vertexes.GetPositionByKey(c.key1)
			p2 := c.vertexes.GetPositionByKey(c.key2)
			asr.Equal(c.p1, p1, "wrong position selected")
			asr.Equal(c.p2, p2, "wrong position selected")
		})
	}
}

func TestResultGraph(t *testing.T) {
	cases := []struct {
		title    string
		vertexes *Vertexes
		result   Path
		st, fn   int
	}{
		{
			title: "small simple one way graph",
			vertexes: NewVertexes(
				[]*Vertex{
					{
						position:    NewPosition(5, mat.NewVecDense(2, []float64{5, 0})),
						connections: []int{4},
					},
					{
						position:    NewPosition(0, mat.NewVecDense(2, []float64{0, 0})),
						connections: []int{1},
					},
					{
						position:    NewPosition(1, mat.NewVecDense(2, []float64{1, 0})),
						connections: []int{0, 2},
					},
					{
						position:    NewPosition(2, mat.NewVecDense(2, []float64{2, 0})),
						connections: []int{1, 3},
					},
					{
						position:    NewPosition(4, mat.NewVecDense(2, []float64{4, 0})),
						connections: []int{3, 5},
					},
					{
						position:    NewPosition(3, mat.NewVecDense(2, []float64{3, 0})),
						connections: []int{2, 4},
					},
				}),
			result: Path{TotalDistance: 5},
			st:     0,
			fn:     5,
		},
		{
			title: "small circular graph",
			vertexes: NewVertexes(
				[]*Vertex{
					{
						position:    NewPosition(5, mat.NewVecDense(2, []float64{0, 1})),
						connections: []int{4, 0},
					},
					{
						position:    NewPosition(0, mat.NewVecDense(2, []float64{0, 0})),
						connections: []int{1, 5},
					},
					{
						position:    NewPosition(1, mat.NewVecDense(2, []float64{1, 0})),
						connections: []int{0, 2},
					},
					{
						position:    NewPosition(2, mat.NewVecDense(2, []float64{2, 0})),
						connections: []int{1, 3},
					},
					{
						position:    NewPosition(4, mat.NewVecDense(2, []float64{4, 0})),
						connections: []int{3, 5},
					},
					{
						position:    NewPosition(3, mat.NewVecDense(2, []float64{3, 0})),
						connections: []int{2, 4},
					},
				}),
			result: Path{TotalDistance: 1},
			st:     0,
			fn:     5,
		},
		{
			title: "large graph",
			vertexes: NewVertexes(
				[]*Vertex{
					{
						position:    NewPosition(0, mat.NewVecDense(2, []float64{0, 0})),
						connections: []int{1, 5, 11},
					},
					{
						position:    NewPosition(1, mat.NewVecDense(2, []float64{10, 0})),
						connections: []int{0, 2},
					},
					{
						position:    NewPosition(2, mat.NewVecDense(2, []float64{20, 0})),
						connections: []int{1, 3},
					},
					{
						position:    NewPosition(3, mat.NewVecDense(2, []float64{30, 0})),
						connections: []int{2, 4},
					},
					{
						position:    NewPosition(4, mat.NewVecDense(2, []float64{40, 0})),
						connections: []int{3, 5},
					},
					{
						position:    NewPosition(5, mat.NewVecDense(2, []float64{5, 0})),
						connections: []int{4, 6, 11, 0},
					},
					{
						position:    NewPosition(6, mat.NewVecDense(2, []float64{60, 0})),
						connections: []int{5, 7},
					},
					{
						position:    NewPosition(7, mat.NewVecDense(2, []float64{70, 0})),
						connections: []int{6, 8},
					},
					{
						position:    NewPosition(8, mat.NewVecDense(2, []float64{80, 0})),
						connections: []int{7, 9},
					},
					{
						position:    NewPosition(9, mat.NewVecDense(2, []float64{90, 0})),
						connections: []int{7, 10, 12, 14},
					},
					{
						position:    NewPosition(10, mat.NewVecDense(2, []float64{5, 5})),
						connections: []int{9, 14, 11},
					},
					{
						position:    NewPosition(11, mat.NewVecDense(2, []float64{0, 5})),
						connections: []int{0, 5, 12, 10},
					},
					{
						position:    NewPosition(12, mat.NewVecDense(2, []float64{0, 20})),
						connections: []int{11, 9},
					},
					{
						position:    NewPosition(13, mat.NewVecDense(2, []float64{0, 30})),
						connections: []int{12, 14},
					},
					{
						position:    NewPosition(14, mat.NewVecDense(2, []float64{0, 40})),
						connections: []int{13, 9, 10},
					},
				}),
			result: Path{TotalDistance: 10},
			st:     0,
			fn:     10,
		},
	}

	asr := assert.New(t)

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			g := NewGraph(c.vertexes)
			result, err := g.CalculateResultGraphFromPosition(c.vertexes.GetPositionByKey(c.st), c.vertexes.GetPositionByKey(c.fn))
			asr.Equal(nil, err, "error should be nil")
			asr.Equal(c.result.TotalDistance, result.TotalDistance, "total distance isn't correct")
			g = NewGraph(c.vertexes)
			result, err = g.CalculateResultGraphFromKeys(c.st, c.fn)
			asr.Equal(nil, err, "error should be nil")
			asr.Equal(c.result.TotalDistance, result.TotalDistance, "total distance isn't correct")
		})
	}
}

var vertex = struct {
	vertexes *Vertexes
}{
	vertexes: NewVertexes(
		[]*Vertex{
			{
				position:    NewPosition(0, mat.NewVecDense(2, []float64{0, 0})),
				connections: []int{1, 5, 11},
			},
			{
				position:    NewPosition(1, mat.NewVecDense(2, []float64{10, 0})),
				connections: []int{0, 2},
			},
			{
				position:    NewPosition(2, mat.NewVecDense(2, []float64{20, 0})),
				connections: []int{1, 3},
			},
			{
				position:    NewPosition(3, mat.NewVecDense(2, []float64{30, 0})),
				connections: []int{2, 4},
			},
			{
				position:    NewPosition(4, mat.NewVecDense(2, []float64{40, 0})),
				connections: []int{3, 5},
			},
			{
				position:    NewPosition(5, mat.NewVecDense(2, []float64{5, 0})),
				connections: []int{4, 6, 11, 0},
			},
			{
				position:    NewPosition(6, mat.NewVecDense(2, []float64{60, 0})),
				connections: []int{5, 7},
			},
			{
				position:    NewPosition(7, mat.NewVecDense(2, []float64{70, 0})),
				connections: []int{6, 8},
			},
			{
				position:    NewPosition(8, mat.NewVecDense(2, []float64{80, 0})),
				connections: []int{7, 9},
			},
			{
				position:    NewPosition(9, mat.NewVecDense(2, []float64{90, 0})),
				connections: []int{7, 10, 12, 14},
			},
			{
				position:    NewPosition(10, mat.NewVecDense(2, []float64{5, 5})),
				connections: []int{9, 14, 11},
			},
			{
				position:    NewPosition(11, mat.NewVecDense(2, []float64{0, 5})),
				connections: []int{0, 5, 12, 10},
			},
			{
				position:    NewPosition(12, mat.NewVecDense(2, []float64{0, 20})),
				connections: []int{11, 9},
			},
			{
				position:    NewPosition(13, mat.NewVecDense(2, []float64{0, 30})),
				connections: []int{12, 14},
			},
			{
				position:    NewPosition(14, mat.NewVecDense(2, []float64{0, 40})),
				connections: []int{13, 9, 10},
			},
		}),
}

var benchcases = []struct {
	st, fn int
	dst    float64
}{
	{
		st:  0,
		fn:  10,
		dst: 10,
	},
	{
		st:  0,
		fn:  14,
		dst: 45.35533905932738,
	},
	{
		st:  0,
		fn:  7,
		dst: 70,
	},
	{
		st:  0,
		fn:  13,
		dst: 55.35533905932738,
	},
	{
		st:  0,
		fn:  11,
		dst: 5,
	},
	{
		st:  0,
		fn:  12,
		dst: 20,
	},
	{
		st:  0,
		fn:  13,
		dst: 55.35533905932738,
	},
	{
		st:  0,
		fn:  6,
		dst: 60,
	},
}

func BenchmarkGraph_CalculateResultGraphFromPosition(b *testing.B) {
	asr := assert.New(b)
	g := NewGraph(vertex.vertexes)
	for n := 0; n < b.N; n++ {
		for _, c := range benchcases {
			_, err := g.CalculateResultGraphFromPosition(vertex.vertexes.GetPositionByKey(c.st), vertex.vertexes.GetPositionByKey(c.fn))
			asr.Equal(err, nil, "error should be nil")
		}
	}
}

func BenchmarkGraph_CalculateResultGraphFromKeys(b *testing.B) {
	asr := assert.New(b)
	g := NewGraph(vertex.vertexes)
	for n := 0; n < b.N; n++ {
		for _, c := range benchcases {
			d, err := g.CalculateResultGraphFromKeys(c.st, c.fn)
			asr.Equal(d.TotalDistance, c.dst, "total distance is incorrect")
			asr.Equal(err, nil, "error should be nil")
		}
	}
}
