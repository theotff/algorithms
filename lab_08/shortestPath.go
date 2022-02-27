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
	dist    []int
	visited []bool
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

func main() {
	fin, _ := os.Open("pathbge1.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	info := strings.Fields(scanner.Text())
	verts, _ := strconv.Atoi(info[0])
	edges, _ := strconv.Atoi(info[1])

	graph := &Graph{
		verts:   verts,
		adjList: make([]*Node, verts),
		dist:    make([]int, verts),
		visited: make([]bool, verts),
	}

	for i := 0; i < edges; i++ {
		scanner.Scan()
		verts := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(verts[0])
		b, _ := strconv.Atoi(verts[1])
		graph.adjList[a-1] = &Node{val: b - 1, next: graph.adjList[a-1]}
		graph.adjList[b-1] = &Node{val: a - 1, next: graph.adjList[b-1]}
	}

	q := &Queue{}
	q.push(0)
	graph.visited[0] = true
	for !q.isEmpty() {
		ind := q.pop()
		curDist := graph.dist[ind]
		ptr := graph.adjList[ind]
		for ptr != nil {
			if !graph.visited[ptr.val] {
				q.push(ptr.val)
				graph.dist[ptr.val] = curDist + 1
				graph.visited[ptr.val] = true
			}
			ptr = ptr.next
		}
	}

	fout, _ := os.Create("pathbge1.out")
	for _, elem := range graph.dist {
		fmt.Fprint(fout, elem, " ")
	}
	fout.Close()
}
