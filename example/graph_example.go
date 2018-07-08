package main

import "github.com/ajz01/graph"
import "fmt"
import "bufio"
import "os"
import "strconv"

func graph_init(add chan graph.Edge, remove chan int, p chan int, quit chan int) {
	data := graph.IntGraph([][]int{{15, 25}, {15, 50}, {25, 50}})
	g, _, _ := graph.InitEdgeList(data)
	for {
		select {
		case v := <-add:
			g.Add(v.V, v.U)
			fmt.Printf("add: %d\n", v)
		case v := <-remove:
			g.Remove(v)
			fmt.Printf("remove: %d\n", v)
		case <-p:
			t := graph.Bfs(g, 0)
			fmt.Printf("%v\n", t)
		case <-quit:
			fmt.Println("quit")
			return
		default:
		}
	}
}

func main() {
	add := make(chan graph.Edge)
	remove := make(chan int)
	p := make(chan int)
	quit := make(chan int)
	go graph_init(add, remove, p, quit)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		v := scanner.Text()
		if v == "q" {
			quit <- 1
			return
		}
		if v[0] == 'a' {
			if len(v) < 5 {
				continue
			}
			u, _ := strconv.Atoi(string(v[2]))
			w, _ := strconv.Atoi(string(v[4]))
			add <- graph.Edge{u, w}
		}
		if v[0] == 'p' {
			p <- 1
		}
	}
}
