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

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) push(val int) {
	node := &Node{val: val}
	if q.head == nil {
		q.tail = node
		q.head = q.tail
	} else {
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue) pop() int {
	value := q.head.val
	q.head = q.head.next
	return value
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}

type Graph struct {
	verts   int
	adjList []*Node
	color   []int
}

func (g *Graph) isBipartite() bool {
	q := &Queue{}
	color := map[int]int{
		1: 2,
		2: 1,
	}

	for index := range g.adjList {
		if g.color[index] == 0 {
			g.color[index] = 1
			q.push(index)
			for !q.isEmpty() {
				ind := q.pop()
				ptr := g.adjList[ind]
				for ptr != nil {
					if ind == ptr.val {
						return false
					}
					if g.color[ptr.val] == 0 {
						q.push(ptr.val)
						g.color[ptr.val] = color[g.color[ind]]
					}
					if g.color[ptr.val] == g.color[ind] {
						return false
					}
					ptr = ptr.next
				}
			}
		}
	}
	return true
}

func main() {
	fin, _ := os.Open("bipartite.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	info := strings.Fields(scanner.Text())
	verts, _ := strconv.Atoi(info[0])
	edges, _ := strconv.Atoi(info[1])

	graph := &Graph{
		verts:   verts,
		adjList: make([]*Node, verts),
		color:   make([]int, verts),
	}

	for i := 0; i < edges; i++ {
		scanner.Scan()
		verts := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(verts[0])
		b, _ := strconv.Atoi(verts[1])
		graph.adjList[a-1] = &Node{val: b - 1, next: graph.adjList[a-1]}
		graph.adjList[b-1] = &Node{val: a - 1, next: graph.adjList[b-1]}
	}

	bipartite := graph.isBipartite()
	fout, _ := os.Create("bipartite.out")
	if bipartite {
		fmt.Fprintln(fout, "YES")
	} else {
		fmt.Fprintln(fout, "NO")
	}
	fout.Close()
}
