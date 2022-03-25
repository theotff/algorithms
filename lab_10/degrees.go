package main

import (
	"bufio"
	"errors"
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
}

func (g *Graph) degrees(degrees []int) {
	for i := 0; i < g.verts; i++ {
		length := 0
		adj := g.adjList[i]
		for adj != nil {
			length += 1
			adj = adj.next
		}
		degrees[i] = length
	}
}

func main() {
	var fin, fout *os.File
	fin, err := os.Open("input.txt")

	if errors.Is(err, os.ErrNotExist) {
		fin, _ = os.Open("degrees.in")
		fout, _ = os.Create("degrees.out")
	} else {
		fout, _ = os.Create("output.txt")
	}
	defer fout.Close()

	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	info := strings.Fields(scanner.Text())
	verts, _ := strconv.Atoi(info[0])
	edges, _ := strconv.Atoi(info[1])

	graph := &Graph{
		verts:   verts,
		adjList: make([]*Node, verts),
	}

	for i := 0; i < edges; i++ {
		scanner.Scan()
		verts := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(verts[0])
		b, _ := strconv.Atoi(verts[1])
		graph.adjList[a-1] = &Node{val: b - 1, next: graph.adjList[a-1]}
		graph.adjList[b-1] = &Node{val: a - 1, next: graph.adjList[b-1]}
	}

	degrees := make([]int, graph.verts)
	graph.degrees(degrees)
	for _, elem := range degrees {
		fmt.Fprintf(fout, "%d ", elem)
	}
}
