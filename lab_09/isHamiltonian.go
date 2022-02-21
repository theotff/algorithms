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

type Stack struct {
	head *Node
}

func (s *Stack) push(val int) {
	s.head = &Node{val: val, next: s.head}
}

func (s *Stack) pop() int {
	value := s.head.val
	s.head = s.head.next
	return value
}

func (s *Stack) isEmpty() bool {
	return s.head == nil
}

type Graph struct {
	verts   int
	edges   int
	adjList []*Node
	visited []bool
}

func (g *Graph) construct(verts, edges int) {
	g.verts = verts
	g.edges = edges
	g.adjList = make([]*Node, g.verts)
	g.visited = make([]bool, g.verts)
}

func (g *Graph) topoSort() []int {
	sorted := make([]int, g.verts)
	s := &Stack{}
	sortInd := 0
	for index := range g.adjList {
		if !g.visited[index] {
			g.visited[index] = true
			s.push(index)
			adj := g.adjList[index]

			for !s.isEmpty() {
				ind := s.pop()
				adj = g.adjList[ind]
				for adj != nil && g.visited[adj.val] {
					adj = adj.next
				}
				if adj != nil {
					s.push(ind)
					for adj != nil {
						if !g.visited[adj.val] {
							s.push(adj.val)
							g.visited[adj.val] = true
						}
						adj = g.adjList[adj.val]
					}
				} else {
					sorted[sortInd] = ind
					sortInd += 1
				}
			}
		}
	}
	return sorted
}

func (g *Graph) isHamiltonian() bool {
	sorted := g.topoSort()
	for i := 0; i < len(sorted)-1; i++ {
		ptr := g.adjList[i]
		for ptr != nil {
			if ptr.val == i+1 {
				break
			}
			ptr = ptr.next
		}
		if ptr == nil {
			return false
		}
	}
	return true
}

func main() {
	fin, _ := os.Open("hamiltonian.in")
	graph := &Graph{}
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	info := strings.Fields(scanner.Text())
	verts, _ := strconv.Atoi(info[0])
	edges, _ := strconv.Atoi(info[1])

	graph.construct(verts, edges)

	for i := 0; i < graph.edges; i++ {
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
