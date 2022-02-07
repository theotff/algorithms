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
	comps   []int
}

func main() {
	fin, _ := os.Open("components.in")
	graph := &Graph{}
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	info := strings.Fields(scanner.Text())
	graph.verts, _ = strconv.Atoi(info[0])
	graph.edges, _ = strconv.Atoi(info[1])

	graph.adjList = make([]*Node, graph.verts)
	graph.comps = make([]int, graph.verts)

	for i := 0; i < graph.edges; i++ {
		scanner.Scan()
		verts := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(verts[0])
		b, _ := strconv.Atoi(verts[1])
		graph.adjList[a-1] = &Node{val: b - 1, next: graph.adjList[a-1]}
		graph.adjList[b-1] = &Node{val: a - 1, next: graph.adjList[b-1]}
	}

	s := &Stack{}
	compCount := 1
	for index := range graph.adjList {
		if graph.comps[index] == 0 {
			graph.comps[index] = compCount
			s.push(index)
			for !s.isEmpty() {
				ind := s.pop()
				ptr := graph.adjList[ind]
				for ptr != nil {
					if graph.comps[ptr.val] == 0 {
						s.push(ptr.val)
						graph.comps[ptr.val] = compCount
					}
					ptr = ptr.next
				}
			}
			compCount += 1
		}
	}

	fout, _ := os.Create("components.out")
	fmt.Fprintln(fout, compCount-1)
	for _, elem := range graph.comps {
		fmt.Fprint(fout, elem, " ")
	}
}
