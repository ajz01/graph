package main

import "fmt"
import "net"
import "bytes"
import "sync"
import "encoding/gob"
import "github.com/ajz01/graph"
import "github.com/ajz01/graph/gnet"

var wg sync.WaitGroup

func handleInput(conn net.Conn) {
	for {
		fmt.Printf("Enter message\n")
		var a string
		var u, v int
		fmt.Scanln(&a, &u, &v)
		var method gnet.Method
		switch a {
		case "Add":
			method = gnet.Add
		case "Bfs":
			method = gnet.Bfs
		case "Quit":
			wg.Done()
		}
		m := gnet.Message{method, graph.IntEdgeList{{u, v}}}
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		enc.Encode(&m)
		_, err := conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("sent")
	}
}

func handleResponse(conn net.Conn) {
	for {
		b := make([]byte, 1024)
		_, err := conn.Read(b)
		var t graph.Traversal
		buf := bytes.NewBuffer(b)
		dec := gob.NewDecoder(buf)
		dec.Decode(&t)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%v\n", t)
	}
}

func main() {
	gob.Register(graph.IntEdgeList{})
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		panic(fmt.Sprintf("could not dial server: %s\n", err.Error()))
	}
	wg.Add(1)
	go handleInput(conn)
	go handleResponse(conn)
	wg.Wait()
}
