package main

import "fmt"
import "net"
import "bytes"
import "encoding/gob"
import "github.com/ajz01/graph"
import "github.com/ajz01/graph/gnet"

func handleRequest(conn net.Conn) {
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
}

func main() {
	gob.Register(graph.IntEdgeList{})
	server, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		panic(fmt.Sprintf("%s\n", err.Error()))
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(fmt.Sprintf("%s\n", err.Error()))
		}
		go handleRequest(conn)
	}
}
