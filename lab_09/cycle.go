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
	visited []int
}

func (g *Graph) construct(verts, edges int) {
	g.verts = verts
	g.edges = edges
	g.adjList = make([]*Node, g.verts)
	g.visited = make([]int, g.verts)
}

func (g *Graph) containsCycle() (bool, []int) {
	s := &Stack{}
	for index := range g.adjList {
		if g.visited[index] == 0 {
			g.visited[index] = 1
			s.push(index)
			adj := g.adjList[index]

			for !s.isEmpty() {
				ind := s.pop()
				adj = g.adjList[ind]
				for adj != nil && g.visited[adj.val] != 0 {
					adj = adj.next
				}
				if adj != nil {
					s.push(ind)
					for adj != nil {
						switch g.visited[adj.val] {
						case 0:
							s.push(adj.val)
							g.visited[adj.val] = 1
						case 1:
							val := s.pop()
							arr := make([]int, 0)
							arr = append(arr, adj.val)
							for val != adj.val {
								arr = append(arr, val)
								val = s.pop()
							}
							return true, arr
						}
						adj = g.adjList[adj.val]
					}
				} else {
					g.visited[ind] = 2
				}
			}
		}
	}
	return false, []int{}
}

func main() {
	fin, _ := os.Open("cycle.in")
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

	fout, _ := os.Create("cycle.out")
	cycle, cycleArr := graph.containsCycle()
	if cycle {
		fmt.Fprintln(fout, "YES")
		for _, elem := range cycleArr {
			fmt.Fprint(fout, elem+1, " ")
		}
	} else {
		fmt.Fprintln(fout, "NO")
	}
	fout.Close()
}
