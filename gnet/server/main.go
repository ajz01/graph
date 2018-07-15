package main

import "fmt"
import "net"
import "bytes"
import "sync"
import "encoding/gob"
import "github.com/ajz01/graph"
import "github.com/ajz01/graph/gnet"

var g graph.Modifiable
var wg sync.WaitGroup

func handleRequest(conn net.Conn) {
	for {
		b := make([]byte, 1024)
		_, err := conn.Read(b)
		if err != nil {
			panic(fmt.Sprintf("%s\n", err.Error()))
		}
		buf := bytes.NewBuffer(b)
		var m gnet.Message
		dec := gob.NewDecoder(buf)
		err = dec.Decode(&m)
		if err != nil {
			panic(fmt.Sprintf("could not decode message: %s", err.Error()))
		}
		fmt.Printf("received: %v\n", m)
		if m.Method == gnet.Add {
			graph.Add(g, m.E)
		}
		if m.Method == gnet.Bfs {
			u, v := m.E.Get(0)
			t := graph.Bfs(g, u, v)
			buf = new(bytes.Buffer)
			enc := gob.NewEncoder(buf)
			enc.Encode(&t)
			conn.Write(buf.Bytes())
		}
		if m.Method == gnet.Dfs {
			t := graph.Dfs(g, 0, 0)
			buf = new(bytes.Buffer)
			enc := gob.NewEncoder(buf)
			enc.Encode(&t)
			conn.Write(buf.Bytes())
		}

		if m.Method == gnet.Print {
			fmt.Println("%v\n", g)
		}
	}
}

func main() {
	g = &graph.IntAdjacencyList{{}}
	gob.Register(graph.IntEdgeList{})
	server, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		panic(fmt.Sprintf("%s\n", err.Error()))
	}
	defer server.Close()
	conn, err := server.Accept()
	if err != nil {
		panic(fmt.Sprintf("%s\n", err.Error()))
	}
	wg.Add(1)
	go handleRequest(conn)
	wg.Wait()
}
