package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		// BFS din nodul src
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}
		for _, d := range dist {
			if d == -1 {
				return -1 // avem un nod la care nu putem ajunge din src
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

// Info reprezinta o pereche (id , valoare ) retinuta in memoria unui nod.
type Info struct {
	ID    int
	Value int
}

// Memory este o multime de token -uri ( reuniunea se face prin inserare ).
type Memory map[Info]struct{}
type Node struct {
	ID        int
	Neighbors []int
	M         Memory
	InChan    chan Memory
	Ctrl      chan int
	Done      chan bool
	Crashed   bool
}

func main() {
	nodes, adj, s := citesteDate()
	diam := diametru(adj)

	if diam == -1 {
		fmt.Println("Graf neconex")
		return
	}

	for _, n := range nodes {
		go runNode(n, nodes)
	}

	simulare(nodes, s, diam)
	afisareRezultat(nodes)
}

func runNode(node *Node, nodes []*Node) {
	for {
		cmd := <-node.Ctrl
		if cmd == -1 || cmd == -2 {
			if cmd == -2 {
				node.Done <- true
			}
			return
		}

		for _, vec := range node.Neighbors {
			if !nodes[vec].Crashed {
				mCopy := make(Memory)
				for k, v := range node.M {
					mCopy[k] = v
				}
				nodes[vec].InChan <- mCopy
			}
		}
		node.Done <- true

		<-node.Ctrl
		for len(node.InChan) > 0 {
			for info := range <-node.InChan {
				node.M[info] = struct{}{}
			}
		}
		node.Done <- true
	}
}

func simulare(nodes []*Node, s, diam int) {
	maxIter := (s + 1) * diam
	crashIndex := 0

	for t := 0; t <= maxIter; t++ {
		if t > 0 && crashIndex < s && crashIndex < len(nodes) {
			victima := nodes[crashIndex]
			victima.Crashed = true
			victima.Ctrl <- -1

			fmt.Printf("T=%d: picat %d. Mem:\n", t, victima.ID)

			for _, n := range nodes {
				if !n.Crashed {
					fmt.Printf("Nod %d: %d\n", n.ID, len(n.M))
				}
			}
			crashIndex++
		}

		semnalizare(nodes, 1)
		semnalizare(nodes, 2)
	}

	for _, n := range nodes {
		if !n.Crashed {
			n.Ctrl <- -2
			<-n.Done
		}
	}
}

func semnalizare(nodes []*Node, cmd int) {
	for _, n := range nodes {
		if !n.Crashed {
			n.Ctrl <- cmd
		}
	}
	for _, n := range nodes {
		if !n.Crashed {
			<-n.Done
		}
	}
}

func citesteDate() ([]*Node, [][]int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Eroare fisier")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	s, _ := strconv.Atoi(parts[1])

	scanner.Scan()
	parts = strings.Fields(scanner.Text())
	tokens := make([]int, n)
	for i := 0; i < n; i++ {
		tokens[i], _ = strconv.Atoi(parts[i])
	}

	nodes := make([]*Node, n)
	adj := make([][]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())

		startIdx := 0
		if len(parts) > 1 && parts[0] == strconv.Itoa(len(parts)-1) {
			startIdx = 1
		}

		for j := startIdx; j < len(parts); j++ {
			val, _ := strconv.Atoi(parts[j])
			adj[i] = append(adj[i], val)
		}

		nodes[i] = &Node{
			ID:        i,
			Neighbors: adj[i],
			M:         Memory{Info{ID: i, Value: tokens[i]}: struct{}{}},
			InChan:    make(chan Memory, n*n),
			Ctrl:      make(chan int),
			Done:      make(chan bool),
		}
	}
	return nodes, adj, s
}

func afisareRezultat(nodes []*Node) {
	for _, n := range nodes {
		if !n.Crashed {
			var finalMem []Info
			for info := range n.M {
				finalMem = append(finalMem, info)
			}
			sort.Slice(finalMem, func(a, b int) bool {
				return finalMem[a].ID < finalMem[b].ID
			})

			fmt.Print("Final: ")
			for _, info := range finalMem {
				fmt.Printf("(%d, %d) ", info.ID, info.Value)
			}
			fmt.Println()
			return
		}
	}
}
