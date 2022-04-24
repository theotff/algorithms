package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const intInf = ^uint32(0)

type graph struct {
	verts   uint16
	edges   uint32
	adjList []*node
	dist    []*dist
}

type node struct {
	val    uint16
	weight uint32
	next   *node
}

type dist struct {
	dist    uint32
	visited bool
}

type item struct {
	ind  uint16
	dist uint32
}

type priorityqueue struct {
	items []*item
	len   int
}

func initPq(len uint16) *priorityqueue {
	return &priorityqueue{
		items: make([]*item, len),
		len:   0,
	}
}

func (pq *priorityqueue) insert(ind uint16, dist uint32) {
	pq.items[pq.len] = &item{
		ind:  ind,
		dist: dist,
	}
	pq.len += 1
	if pq.len == 2 {
		if pq.items[1].dist < pq.items[0].dist {
			pq.items[0], pq.items[1] = pq.items[1], pq.items[0]
		}
	} else if pq.len > 2 {
		index := pq.len - 1
		parent := (index - 1) / 2
		for pq.items[index].dist < pq.items[parent].dist {
			pq.items[index], pq.items[parent] = pq.items[parent], pq.items[index]
			index = parent
			parent = (index - 1) / 2
		}
	}
}

func (pq *priorityqueue) removeMin() *item {
	var result *item
	if pq.len != 0 {
		result = pq.items[0]
		pq.items[0] = pq.items[pq.len-1]
		pq.len -= 1
		if pq.len > 0 {
			index := 0
			count := pq.len - 1
			for {
				child1 := 2*index + 1
				child2 := 2*index + 2

				if child1 > count {
					child1 = index
				}
				if child2 > count {
					child2 = index
				}

				if pq.items[index].dist <= pq.items[child1].dist && pq.items[index].dist <= pq.items[child2].dist {
					break
				}
				swapChild := child2
				if pq.items[child1].dist < pq.items[child2].dist {
					swapChild = child1
				}
				pq.items[index], pq.items[swapChild] = pq.items[swapChild], pq.items[index]
				index = swapChild
			}
		}
	}
	return result
}

func (pq *priorityqueue) replace(ind uint16, dist uint32) {
	index := 0
	for i, item := range pq.items {
		if item.ind == ind {
			item.dist = dist
			index = i
			break
		}
	}
	if pq.len > 1 {
		parent := (index - 1) / 2
		for pq.items[index].dist < pq.items[parent].dist {
			pq.items[index], pq.items[parent] = pq.items[parent], pq.items[index]
			index = parent
			parent = (index - 1) / 2
		}
	}
}

func (pq priorityqueue) isEmpty() bool {
	return pq.len == 0
}

func (g *graph) initDist(start uint16, pq *priorityqueue) {
	var i uint16
	for i = 0; i < g.verts; i++ {
		if i != start-1 {
			g.dist[i] = &dist{
				dist:    intInf,
				visited: false,
			}
		} else {
			g.dist[i] = &dist{
				dist:    0,
				visited: false,
			}
		}
		pq.insert(i, g.dist[i].dist)
	}
}

func dijkstra(graph *graph, start uint16) {
	pq := initPq(graph.verts)
	graph.initDist(start, pq)
	for !pq.isEmpty() {
		u := pq.removeMin()
		graph.dist[u.ind].visited = true
		ptr := graph.adjList[u.ind]
		for ptr != nil {
			if graph.dist[ptr.val].dist > graph.dist[u.ind].dist+ptr.weight {
				graph.dist[ptr.val].dist = graph.dist[u.ind].dist + ptr.weight
				pq.replace(ptr.val, graph.dist[u.ind].dist+ptr.weight)
			}
			ptr = ptr.next
		}
	}
}

func main() {
	fin, _ := os.Open("pathbgep.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	info := strings.Fields(scanner.Text())
	verts, _ := strconv.ParseUint(info[0], 10, 16)
	edges, _ := strconv.ParseUint(info[1], 10, 32)

	graph := graph{
		verts:   uint16(verts),
		edges:   uint32(edges),
		adjList: make([]*node, verts),
		dist:    make([]*dist, verts),
	}

	var i uint32
	for i = 0; i < graph.edges; i++ {
		scanner.Scan()
		vert := strings.Fields(scanner.Text())
		a, _ := strconv.ParseUint(vert[0], 10, 16)
		b, _ := strconv.ParseUint(vert[1], 10, 16)
		weight, _ := strconv.ParseUint(vert[2], 10, 32)

		graph.adjList[a-1] = &node{val: uint16(b - 1), weight: uint32(weight), next: graph.adjList[a-1]}
		graph.adjList[b-1] = &node{val: uint16(a - 1), weight: uint32(weight), next: graph.adjList[b-1]}
	}

	fout, _ := os.Create("pathbgep.out")
	defer fout.Close()
	dijkstra(&graph, 1)
	for _, dist := range graph.dist {
		fmt.Fprint(fout, dist.dist, " ")
	}
}
