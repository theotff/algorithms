package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

type Graph struct {
	verts   int
	adjList []*Node
	depth   []int
}

func (g *Graph) DFS(pos int) {
	adj := g.adjList[pos]
	vertDepth := g.depth[pos]
	for adj != nil {
		if g.depth[adj.val] == 0 {
			g.DFS(adj.val)
		}
		vertDepth = intMax(vertDepth, g.depth[adj.val])
		adj = adj.next
	}
	g.depth[pos] = vertDepth + 1
}

func (g *Graph) isHamiltonian() bool {
	for i := 0; i < g.verts; i++ {
		if g.depth[i] == 0 {
			g.DFS(i)
		}
	}
	for _, elem := range g.depth {
		if elem == g.verts {
			return true
		}
	}
	return false
}

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fin, _ := os.Open("hamiltonian.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	info := strings.Fields(scanner.Text())
	verts, _ := strconv.Atoi(info[0])
	edges, _ := strconv.Atoi(info[1])

	graph := &Graph{
		verts:   verts,
		adjList: make([]*Node, verts),
		depth:   make([]int, verts),
	}

	for i := 0; i < edges; i++ {
		scanner.Scan()
		verts := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(verts[0])
		b, _ := strconv.Atoi(verts[1])
		graph.adjList[a-1] = &Node{val: b - 1, next: graph.adjList[a-1]}
	}

	fout, _ := os.Create("hamiltonian.out")
	hamiltonian := graph.isHamiltonian()
	if hamiltonian {
		fmt.Fprintln(fout, "YES")
	} else {
		fmt.Fprintln(fout, "NO")
	}
	fout.Close()
}
