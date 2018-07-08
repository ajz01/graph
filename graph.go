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
	//Set(i, j, v int)
	Size() int
	Len(int) int
}

type Modifiable interface {
	Interface
	Add(int, int)
	Remove(int)
}

type Edge struct {
	U int
	V int
}

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

// InitEdgeList returns an IntGraph re-arranged as an adjacency list of new unique vertex ids built
// from an edge list graph.Interface g.
func InitEdgeList(g Interface) (IntGraph, map[int]int, error) {
	m := make(map[int]int, g.Size())
	for i := 0; i < g.Size(); i++ {
		for j := 0; j < g.Len(i); j++ {
			m[g.Get(i, j)] = -1
		}
	}
	n := 0
	for i := 0; i < g.Size(); i++ {
		for j := 0; j < g.Len(i); j++ {
			if m[g.Get(i, j)] == -1 {
				m[g.Get(i, j)] = n
				n++
			}
		}
	}
	d := make(IntGraph, n)
	for i := 0; i < g.Size(); i++ {
		d[m[g.Get(i, 0)]] = append(d[m[g.Get(i, 0)]], m[g.Get(i, 1)])
	}
	return d, m, nil
}

// Bfs returns a traversal based on graph data
// collected during a breadth first search.
func Bfs(g Interface, s int) Traversal {
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
		t.Color[u] = Black
	}
	return t
}

// Dfs returns a traversal based on graph data
// collected during a depth first search.
func Dfs(g Interface, s int) Traversal {
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
		t.Color[u] = Black
	}
	return t
}

// Convience types for common cases
type IntGraph [][]int

func (g IntGraph) Get(i, j int) int { return g[i][j] }
//func (g IntGraph) Set(i, j, v int)  { g[i][j] = v }
func (g IntGraph) Size() int        { return len(g) }
func (g IntGraph) Len(i int) int    { return len(g[i]) }
func (g *IntGraph) Add(v, u int) {
	if v > len(*g) - 1 {
		a := make([][]int, v - len(*g))
		*g = append(*g, a...)
	}
	(*g)[v - 1] = append((*g)[v - 1], u)
}
func (g *IntGraph) Remove(v int) {
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
type StringId struct {
	S  string
	Id int
}

type StringGraph [][]StringId

func (g StringGraph) Get(i, j int) int { return g[i][j].Id }
func (g StringGraph) Set(i, j, v int)  { g[i][j].Id = v }
func (g StringGraph) Size() int        { return len(g) }
func (g StringGraph) Len(i int) int    { return len(g[i]) }
