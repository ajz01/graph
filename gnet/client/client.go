package client

import (
	"fmt"
	"net"
	"bytes"
	"encoding/gob"
	"github.com/ajz01/graph"
	"github.com/ajz01/graph/gnet"
)

var conn net.Conn

func init() {
	gob.Register(graph.IntEdgeList{})
	var err error
	conn, err = net.Dial("tcp", "localhost:5000")
	if err != nil {
		panic(fmt.Sprintf("could not dial server: %s\n", err.Error()))
	}
	go handleResponse(conn)
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
			return
		}
		fmt.Printf("%v\n", t)
	}
}

func SendMessage(method gnet.Method, u, v int) {
	m := gnet.Message{method, graph.IntEdgeList{{u, v}}}
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	enc.Encode(&m)
	_, err := conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err.Error())
	}
}
