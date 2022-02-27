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

func (s *Stack) peek() int {
	val := s.head.val
	return val
}

func (s *Stack) isEmpty() bool {
	return s.head == nil
}

type Graph struct {
	verts   int
	edges   int
	adjList []*Node
	visited []int
}

func (g *Graph) construct(verts, edges int) {
	g.verts = verts
	g.edges = edges
	g.adjList = make([]*Node, g.verts)
	g.visited = make([]int, g.verts)
}

func (g *Graph) decompose(comps []int) {
	s := &Stack{}
	component := 0
	for index := range g.adjList {
		if g.visited[index] == 0 {
			g.visited[index] = 1
			s.push(index)
			adj := g.adjList[index]

			for !s.isEmpty() {
				ind := s.pop()
				adj = g.adjList[ind]
				for adj != nil && g.visited[adj.val] == 2 {
					adj = adj.next
				}
				if adj != nil {
					s.push(ind)
					for adj != nil {
						switch g.visited[adj.val] {
						case 0:
							g.visited[adj.val] = 1
							s.push(adj.val)
							adj = g.adjList[adj.val]
						case 1:
							component += 1
							vert := s.pop()
							g.visited[vert] = 2
							comps[vert] = component
							for vert != adj.val {
								vert = s.pop()
								g.visited[vert] = 2
								comps[vert] = component
							}
							adj = nil
						default:
							adj = g.adjList[adj.val]
						}
					}
				} else {
					g.visited[ind] = 2
					component += 1
					comps[ind] = component
				}
			}
		}
	}
}

func main() {
	fin, _ := os.Open("cond.in")
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

	components := make([]int, graph.verts)
	graph.decompose(components)
	fout, _ := os.Create("cond.out")
	for _, elem := range components {
		fmt.Fprint(fout, elem, " ")
	}
	fout.Close()
}
