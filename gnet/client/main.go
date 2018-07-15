package main

import "fmt"
import "net"
import "bytes"
import "encoding/gob"
import "github.com/ajz01/graph"
import "github.com/ajz01/graph/gnet"

func main() {
	gob.Register(graph.IntEdgeList{})
	m := gnet.Message{gnet.Add, graph.IntEdgeList{{0, 1}}}
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		panic(fmt.Sprintf("could not dial server: %s\n", err.Error()))
	}
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	enc.Encode(&m)
	conn.Write(buf.Bytes())
	fmt.Println("sent")
}
