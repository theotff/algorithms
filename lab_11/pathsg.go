package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const intInf = ^uint64(0)

type graph struct {
	verts     int
	edges     int
	adjMatrix [][]uint64
	visited   []bool
	dist      []uint64
}

type item struct {
	ind  int
	dist uint64
}

type priorityqueue struct {
	items []*item
	len   int
}

func initPq(len int) *priorityqueue {
	return &priorityqueue{
		items: make([]*item, len),
		len:   0,
	}
}

func (pq *priorityqueue) insert(ind int, dist uint64) {
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

func (pq *priorityqueue) replace(ind int, dist uint64) {
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

func (g *graph) initDist(start int, pq *priorityqueue) {
	for i := 0; i < g.verts; i++ {
		if i != start-1 {
			g.dist[i] = intInf
		} else {
			g.dist[i] = 0
		}
		pq.insert(i, g.dist[i])
	}
}

func dijkstra(graph *graph, start int) {
	pq := initPq(graph.verts)
	graph.initDist(start, pq)
	for !pq.isEmpty() {
		u := pq.removeMin()
		graph.visited[u.ind] = true
		for ind, weight := range graph.adjMatrix[u.ind] {
			if graph.dist[ind] > graph.dist[u.ind]+weight && weight != intInf-1 {
				graph.dist[ind] = graph.dist[u.ind] + weight
				pq.replace(ind, graph.dist[u.ind]+weight)
			}
		}
	}
}

func main() {
	fin, _ := os.Open("pathsg.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	info := strings.Fields(scanner.Text())
	verts, _ := strconv.Atoi(info[0])
	edges, _ := strconv.Atoi(info[1])

	graph := graph{
		verts:     verts,
		edges:     edges,
		adjMatrix: make([][]uint64, verts),
		visited:   make([]bool, verts),
		dist:      make([]uint64, verts),
	}

	for j := 0; j < graph.verts; j++ {
		graph.adjMatrix[j] = make([]uint64, graph.verts)
		for x := 0; x < graph.verts; x++ {
			graph.adjMatrix[j][x] = intInf - 1
		}
	}

	for i := 0; i < graph.edges; i++ {
		scanner.Scan()
		vert := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(vert[0])
		b, _ := strconv.Atoi(vert[1])
		weight, _ := strconv.ParseUint(vert[2], 10, 64)

		graph.adjMatrix[a-1][b-1] = weight
	}

	fout, _ := os.Create("pathsg.out")
	defer fout.Close()

	for start := 1; start <= graph.verts; start++ {
		dijkstra(&graph, start)
		for _, dist := range graph.dist {
			fmt.Fprint(fout, dist, " ")
		}
		fmt.Fprintln(fout)
	}
}
