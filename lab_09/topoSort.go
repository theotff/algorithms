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
	cycle   []bool
}

func (g *Graph) construct(verts, edges int) {
	g.verts = verts
	g.edges = edges
	g.adjList = make([]*Node, g.verts)
	g.visited = make([]bool, g.verts)
}

func (g *Graph) topoSort(sorted []int) bool {
	s := &Stack{}
	sortInd := 0
	for index := range g.adjList {
		if !g.visited[index] {
			g.visited[index] = true
			s.push(index)
			adj := g.adjList[index]
			g.cycle = make([]bool, g.verts)
			for adj != nil {
				if !g.visited[adj.val] {
					s.push(adj.val)
					g.visited[adj.val] = true
				}
				if g.cycle[adj.val] {
					return true
				}
				g.cycle[adj.val] = true
				adj = g.adjList[adj.val]
			}
			for !s.isEmpty() {
				ind := s.pop()
				sorted[sortInd] = ind + 1
				sortInd += 1
				node := g.adjList[ind]
				for node != nil {
					if !g.visited[node.val] {
						s.push(node.val)
						g.visited[node.val] = true
					}
					node = node.next
				}
			}
		}
	}
	return false
}

func main() {
	fin, _ := os.Open("topsort.in")
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

	fout, _ := os.Create("topsort.out")
	sorted := make([]int, graph.verts)
	cycle := graph.topoSort(sorted)
	if cycle {
		fmt.Fprintln(fout, -1)
	} else {
		for i := len(sorted) - 1; i >= 0; i-- {
			fmt.Fprint(fout, sorted[i], " ")
		}
	}
	fout.Close()
}
