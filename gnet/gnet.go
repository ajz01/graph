package gnet

import "github.com/ajz01/graph"

type Method int

const (
	Add Method = iota
	Remove
)

type Message struct {
	Method Method
	E graph.EdgeList
}
