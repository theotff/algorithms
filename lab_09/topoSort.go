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
	adjList []*Node
	visited []int
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func (g *Graph) topoSort(sorted []int) bool {
	s := new(Stack)
	sortInd := 0
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
							s.push(adj.val)
							g.visited[adj.val] = 1
						case 1:
							return true
						}
						adj = g.adjList[adj.val]
					}
				} else {
					g.visited[ind] = 2
					sorted[sortInd] = ind
					sortInd += 1
				}
			}
		}
	}
	reverse(sorted)
	return false
}

func main() {
	fin, _ := os.Open("topsort.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	info := strings.Fields(scanner.Text())
	verts, _ := strconv.Atoi(info[0])
	edges, _ := strconv.Atoi(info[1])

	graph := &Graph{
		verts:   verts,
		adjList: make([]*Node, verts),
		visited: make([]int, verts),
	}

	for i := 0; i < edges; i++ {
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
		for i := 0; i < len(sorted); i++ {
			fmt.Fprint(fout, sorted[i]+1, " ")
		}
	}
	fout.Close()
}
