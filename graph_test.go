package graph_test

import "github.com/ajz01/graph"
import "testing"
import "fmt"

func CompareTraversal(tr, r graph.Traversal, t *testing.T) {
	if fmt.Sprintf("%v", tr.Color) != fmt.Sprintf("%v", r.Color) {
		t.Errorf("Invalid Color have %v want %v", tr.Color, r.Color)
	}
	if fmt.Sprintf("%v", tr.VertexOrdering) != fmt.Sprintf("%v", r.VertexOrdering) {
		t.Errorf("Invalid Vertex Ordering have %v want %v", tr.VertexOrdering, r.VertexOrdering)
	}
	if fmt.Sprintf("%v", tr.EdgeOrdering) != fmt.Sprintf("%v", r.EdgeOrdering) {
		t.Errorf("Invalid Edge Ordering have %v want %v", tr.EdgeOrdering, r.EdgeOrdering)
	}
	if fmt.Sprintf("%v", tr.Parent) != fmt.Sprintf("%v", r.Parent) {
		t.Errorf("Invalid Parent have %v want %v", tr.Parent, r.Parent)
	}
	if fmt.Sprintf("%v", tr.Distance) != fmt.Sprintf("%v", r.Distance) {
		t.Errorf("Invalid Distance have %v want %v", tr.Distance, r.Distance)
	}
}

func TestConnectedBfs(t *testing.T) {
	data := graph.IntEdgeList{{15, 25}, {15, 50}, {25, 50}}
	g, _, _ := graph.EdgeToAdjacencyList(data)
	tr := graph.Bfs(g, 0, -1)
	r := graph.Traversal{
		[]graph.Color{graph.Black, graph.Black, graph.Black},
		[]int{0, 1, 2},
		[]graph.Edge{{0, 1}, {0, 2}, {1, 2}},
		[]int{-1, 0, 0},
		[]int{0, 1, 1},
	}
	CompareTraversal(tr, r, t)
}

func TestStringConnectedBfs(t *testing.T) {
	data := graph.StringEdgeList{
		{
			graph.StringId{"Test", 15},
			graph.StringId{"Test", 25},
		}, {
			graph.StringId{"Test", 15},
			graph.StringId{"Test", 50},
		}, {
			graph.StringId{"Test", 25},
			graph.StringId{"Test", 50},
		},
	}
	g, _, _ := graph.EdgeToAdjacencyList(data)
	tr := graph.Bfs(g, 0, -1)
	r := graph.Traversal{
		[]graph.Color{graph.Black, graph.Black, graph.Black},
		[]int{0, 1, 2},
		[]graph.Edge{{0, 1}, {0, 2}, {1, 2}},
		[]int{-1, 0, 0},
		[]int{0, 1, 1},
	}
	CompareTraversal(tr, r, t)
}

func TestConnectedDfs(t *testing.T) {
	data := graph.IntEdgeList{{15, 25}, {15, 50}, {25, 50}}
	g, _, _ := graph.EdgeToAdjacencyList(data)
	tr := graph.Dfs(g, 0, -1)
	r := graph.Traversal{
		[]graph.Color{graph.Black, graph.Black, graph.Black},
		[]int{0, 2, 1},
		[]graph.Edge{{0, 1}, {0, 2}, {1, 2}},
		[]int{-1, 0, 0},
		[]int{0, 1, 1},
	}
	CompareTraversal(tr, r, t)
}

func TestStrongConnComponents(t *testing.T) {
	data := graph.IntEdgeList{{15, 25}, {15, 50}, {25, 50}, {35, 75}, {100, 300}}
	g, _, _ := graph.EdgeToAdjacencyList(data)
	tr := graph.StrongConnComponents(g, graph.Dfs)
	r := []graph.Traversal{
		graph.Traversal{
			[]graph.Color{graph.Black, graph.Black, graph.Black, graph.White, graph.White, graph.White, graph.White},
			[]int{0, 2, 1},
			[]graph.Edge{{0, 1}, {0, 2}, {1, 2}},
			[]int{-1, 0, 0, 0, 0, 0, 0},
			[]int{0, 1, 1, 0, 0, 0, 0},
		},
		graph.Traversal{
			[]graph.Color{graph.Black, graph.Black, graph.Black},
			[]int{0, 2, 1},
			[]graph.Edge{{0, 1}, {0, 2}, {1, 2}},
			[]int{-1, 0, 0},
			[]int{0, 1, 1},
		},
		graph.Traversal{
			[]graph.Color{graph.Black, graph.Black, graph.Black},
			[]int{0, 2, 1},
			[]graph.Edge{{0, 1}, {0, 2}, {1, 2}},
			[]int{-1, 0, 0},
			[]int{0, 1, 1},
		},
	}
	for i, _ := range tr {
		CompareTraversal(tr[i], r[i], t)
	}
}

func TestShortestPath(t *testing.T) {
	data := graph.IntEdgeList{{15, 25}, {15, 50}, {25, 50}, {35, 75}, {15, 35}}
	g, m, _ := graph.EdgeToAdjacencyList(data)
	tr := graph.ShortestPath(g, m[15], m[30])
	if tr == nil {
		t.Error("Shortest path failed no path.")
	}
}

func TestWeightedShortestPath(t *testing.T) {
	data := graph.IntWeightedEdgeList{{15, 25, 5}, {25, 30, 8},{50, 75, 10}}
	g, m, _ := graph.EdgeToAdjacencyList(data)
	tr := graph.ShortestPath(g, m[15], m[30])
	if tr == nil {
		t.Error("Shortest path failed no path.")
	}
}
