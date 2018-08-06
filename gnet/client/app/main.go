package main

import "fmt"
import "sync"
import "github.com/ajz01/graph/gnet"
import "github.com/ajz01/graph/gnet/client"

var wg sync.WaitGroup

func handleInput() {
	for {
		fmt.Printf("Enter message\n")
		var a string
		var u, v int
		fmt.Scanln(&a, &u, &v)
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
		case "Quit":
			wg.Done()
		}
		client.SendMessage(method, u, v)
		fmt.Println("sent")
	}
}

func main() {
	wg.Add(1)
	go handleInput()
	wg.Wait()
}
