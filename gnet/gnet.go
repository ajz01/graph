package gnet

import "github.com/ajz01/graph"

type Method int

const (
	Add Method = iota
	Remove
	Bfs
	Dfs
	StrongConnected
	Print
)

type Message struct {
	Method Method
	E graph.EdgeList
}
