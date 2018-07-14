// Pacakge graph provides graph operations for any type that
// implements graph.Interface. Traversal data is provided by
// the graph.Traversal type.
package graph

import "github.com/ajz01/graph/queue"
import "github.com/ajz01/graph/stack"

type Color int

const (
	White Color = iota
	Gray
	Black
)

// Any type that implements graph.Interface can be used as a graph.
// The type must be able to provide some adjacency list like behavior
// through the graph.Interface as follows:
//
// Get(i, j) shall provide a unique id of a vertex at the jth adjacency position of the vertex with id i
// Set(i, j) shall set the id of the vertex at the jth adjacency position of the vertex having id i
// Size() shall return the total number of vertices in the graph
// Len() shall return the number of vertices adjacent to the vertex v having id i
//
// Edge lists consisting of a [][]graph.Interface can be converted using the InitEdgeList graph.Interface method.
type Interface interface {
	Get(int, int) int
	Size() int
	Len(int) int
}

type Weighted interface {
	Interface
	Weight(int, int) int
}

type Modifiable interface {
	Interface
	Add(int, int)
	Remove(int)
}

type WeightedModifiable interface {
	Interface
	Add(int, int, int)
	Remove(int)
}

type EdgeList interface {
	Len() int
	Get(int) (int, int)
}

type WeightedEdgeList interface {
	EdgeList
	Weight(int) int
}

type Edge struct {
	U int
	V int
}

type WeightedEdge struct {
	U int
	V int
	Weight int
}

// Traversal stores the nformation discovered during a graph traversal operation.
// for methods such as StrongConnComponents that run multiple passes of a traversal
// a slice of Traversals is returned so that each unique component can be analyzed.
// The index of each Traversal in such a Traversal slice can be used as a component id
// and the vertices belonging to that component are easily identified by inspecting the
// VertexOrdering field of each components Traversal
type Traversal struct {
	Color          []Color
	VertexOrdering []int
	EdgeOrdering   []Edge
	Parent         []int
	Distance       []int
}

func NewTraversal(n int) Traversal {
	var t Traversal
	t.Color = make([]Color, n)
	t.Parent = make([]int, n)
	t.Distance = make([]int, n)
	return t
}

// EdgeToAdjacencyList converts and edge list to adjacency list.
// adjacency lists are more efficient, but edge lists are likely
// input formats
func EdgeToAdjacencyList(g EdgeList) (Interface, map[int]int, error) {
	m := make(map[int]int, g.Len())
	for i := 0; i < g.Len(); i++ {
		u, v := g.Get(i)
		m[u] = -1
		m[v] = -1
	}
	n := 0
	for i := 0; i < g.Len(); i++ {
		u, v := g.Get(i)
		if m[u] == -1 {
			m[u] = n
			n++
		}
		if m[v] == -1 {
			m[v] = n
			n++
		}
	}
	if w, ok := g.(WeightedEdgeList); ok {
		d := make(IntWeightedAdjacencyList, n)
		for i := 0; i < g.Len(); i++ {
			u, v := g.Get(i)
			weight := w.Weight(i)
			d.Add(m[u], m[v], weight)
		}
		return d, m, nil
	}
	d := make(IntAdjacencyList, n)
	for i := 0; i < g.Len(); i++ {
		u, v := g.Get(i)
		d.Add(m[u], m[v])
	}
	return d, m, nil
}

// Bfs returns a Traversal based on graph data
// collected during a breadth first search.
func Bfs(g Interface, s, d int) Traversal {
	t := NewTraversal(g.Size())
	sz := g.Size()
	for v := 0; v < sz; v++ {
		t.Color[v] = White
	}
	t.Color[s] = Gray
	t.Parent[s] = -1
	t.Distance[s] = 0
	var q queue.IntSlice
	q.Add(s)
	for !q.Empty() {
		u := q.Get()
		t.VertexOrdering = append(t.VertexOrdering, u)
		t.Color[u] = Black
		if u == d {
			return t
		}
		for i := 0; i < g.Len(u); i++ {
			v := g.Get(u, i)
			t.EdgeOrdering = append(t.EdgeOrdering, Edge{u, v})
			if t.Color[v] == White {
				t.Color[v] = Gray
				t.Parent[v] = u
				t.Distance[v] = t.Distance[u] + 1
				q.Add(v)
			}
		}
		q.Remove()
	}
	return t
}

// Dfs returns a Traversal based on graph data
// collected during a depth first search.
func Dfs(g Interface, s, d int) Traversal {
	t := NewTraversal(g.Size())
	sz := g.Size()
	for v := 0; v < sz; v++ {
		t.Color[v] = White
	}
	t.Color[s] = Gray
	t.Parent[s] = -1
	t.Distance[s] = 0
	var st stack.IntSlice
	st.Push(s)
	for !st.Empty() {
		u := st.Pop()
		t.VertexOrdering = append(t.VertexOrdering, u)
		t.Color[u] = Black
		if u == d {
			return t
		}
		for i := 0; i < g.Len(u); i++ {
			v := g.Get(u, i)
			t.EdgeOrdering = append(t.EdgeOrdering, Edge{u, v})
			if t.Color[v] == White {
				t.Color[v] = Gray
				t.Parent[v] = u
				t.Distance[v] = t.Distance[u] + 1
				st.Push(v)
			}
		}
	}
	return t
}

// StrongConnComponents returns a slice of Traversals one for each strongly connected component.
// the index of each Traversal can be used as a unique component id and the verices belonging
// to that component can be obtained from the VertexOrdering field of the Traversal.
func StrongConnComponents(g Interface, tm func(Interface, int, int) Traversal) []Traversal {
	t := make([]Color, g.Size())
	var m []Traversal
	for i := 0; i < g.Size(); i++ {
		if t[i] == White {
			m = append(m, tm(g, i, -1))
		}
		for j, v := range m[len(m)-1].Color {
			t[j] = v
		}
	}
	return m
}

func ShortestPath(g Interface, s, d int) *Traversal {
	if s < 0 || s > g.Size() || d < 0 || d > g.Size() {
		return nil
	}
	switch g.(type) {
	case Weighted:
		// replace with weighted graph shortest path algorithm
		tr := Bfs(g, s, d)
		if tr.Color[d] == White {
			return nil
		}
		return &tr
	case Interface:
		tr := Bfs(g, s, d)
		if tr.Color[d] == White {
			return nil
		}
		return &tr
	}
	return nil
}

// Convience types for common cases
type IntAdjacencyList [][]int

func (g IntAdjacencyList) Get(i, j int) int { return g[i][j] }

//func (g IntAdjacencyList) Set(i, j, v int)  { g[i][j] = v }
func (g IntAdjacencyList) Size() int     { return len(g) }
func (g IntAdjacencyList) Len(i int) int { return len(g[i]) }
func (g *IntAdjacencyList) Add(v, u int) {
	if v > len(*g)-1 {
		a := make([][]int, v-len(*g))
		*g = append(*g, a...)
	}
	(*g)[v] = append((*g)[v], u)
}
func (g *IntAdjacencyList) Remove(v int) {
	for i := range *g {
		for j := range (*g)[i] {
			if (*g)[i][j] == v {
				(*g)[i] = append((*g)[i][:j-1], (*g)[i][j:]...)
			}
		}
		if i == v {
			*g = append((*g)[:i-1], (*g)[i:]...)
		}
	}
}

type IntEdgeList [][2]int
func (g IntEdgeList) Len() int { return len(g) }
func (g IntEdgeList) Get(i int) (int, int) { return g[i][0], g[i][1] }

type IntWeightedEdgeList [][3]int
func (g IntWeightedEdgeList) Len() int { return len(g) }
func (g IntWeightedEdgeList) Get(i int) (int, int) { return g[i][0], g[i][1] }
func (g IntWeightedEdgeList) Weight(i int) int { return g[i][2] }

type IntWeightedAdjacencyList [][][]int

func (g IntWeightedAdjacencyList) Get(i, j int) int { return g[i][j][0] }

func (g IntWeightedAdjacencyList) Size() int     { return len(g) }
func (g IntWeightedAdjacencyList) Len(i int) int { return len(g[i]) }
func (g *IntWeightedAdjacencyList) Add(v, u, w int) {
	if v > len(*g)-1 {
		a := make([][][]int, v-len(*g))
		*g = append(*g, a...)
	}
	(*g)[v] = append((*g)[v], []int{u, w})
}
func (g *IntWeightedAdjacencyList) Remove(v int) {
	for i := range *g {
		for j := range (*g)[i] {
			if (*g)[i][j][0] == v {
				(*g)[i] = append((*g)[i][:j-1], (*g)[i][j:]...)
			}
		}
		if i == v {
			*g = append((*g)[:i-1], (*g)[i:]...)
		}
	}
}

func (g IntWeightedAdjacencyList) Weight(i, j int) int { return g[i][j][1] }

type StringId struct {
	S  string
	Id int
}

type StringEdgeList[][2]StringId
func (g StringEdgeList) Len() int { return len(g) }
func (g StringEdgeList) Get(i int) (int, int) { return g[i][0].Id, g[i][1].Id }

type StringAdjacencyList [][]StringId

func (g StringAdjacencyList) Get(i, j int) int { return g[i][j].Id }
func (g StringAdjacencyList) Size() int        { return len(g) }
func (g StringAdjacencyList) Len(i int) int    { return len(g[i]) }
