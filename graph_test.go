package graph_test

import "test/graph"
import "testing"
import "fmt"

func CompareTraversal(tr, r graph.Traversal, t *testing.T) {
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
	data := graph.IntGraph {{15, 25},{15, 50},{25, 50}}
	g, _ := graph.InitEdgeList(data)
	tr := graph.Bfs(g, 0)
	r := graph.Traversal{
		[]graph.Color{graph.Black, graph.Black, graph.Black},
		[]int{0, 1, 2},
		[]graph.Edge{graph.Edge{0, 1}, graph.Edge{0, 2}, graph.Edge{1, 2}},
		[]int{-1, 0, 0},
		[]int{0, 1, 1},
	}
	CompareTraversal(tr, r, t)
}

func TestStringConnectedBfs(t *testing.T) {
	data := graph.StringGraph {
		{
			graph.StringId{"Test", 15},
			graph.StringId{"Test", 25},
		},{
			graph.StringId{"Test", 15},
			graph.StringId{"Test", 50},
		},{
			graph.StringId{"Test", 25},
			graph.StringId{"Test", 50},
		},
	}
	g, _ := graph.InitEdgeList(data)
	tr := graph.Bfs(g, 0)
	r := graph.Traversal{
		[]graph.Color{graph.Black, graph.Black, graph.Black},
		[]int{0, 1, 2},
		[]graph.Edge{graph.Edge{0, 1}, graph.Edge{0, 2}, graph.Edge{1, 2}},
		[]int{-1, 0, 0},
		[]int{0, 1, 1},
	}
	CompareTraversal(tr, r, t)
}
