package main

import "fmt"
import "os"
import "strconv"
import "github.com/ajz01/graph/gnet"
import "github.com/ajz01/graph/gnet/client"

func handleInput(a string, u, v int) {
	var method gnet.Method
	switch a {
	case "Add":
		method = gnet.Add
	case "Remove":
		method = gnet.Remove
	case "Bfs":
		method = gnet.Bfs
	case "Dfs":
		method = gnet.Dfs
	}
	client.SendMessage(method, u, v)
	fmt.Println("sent")
}

func main() {
	if len(os.Args) < 4 {
		return
	}
	u, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error %s not integer type.\n", os.Args[2])
	}
	v, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Error %s not integer type.\n", os.Args[3])
	}
	handleInput(os.Args[1], u, v)
}
