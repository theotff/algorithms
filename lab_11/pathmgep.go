package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const intInf = ^uint64(0)

var queueLen = 0

type graph struct {
	verts     int
	start     int
	stop      int
	adjMatrix [][]uint64
	visited   []bool
	dist      []uint64
}

type item struct {
	ind  int
	dist uint64
}

type priorityqueue []*item

func (pq *priorityqueue) insert(ind int, dist uint64) {
	(*pq)[queueLen] = &item{
		ind:  ind,
		dist: dist,
	}
	queueLen += 1
	if queueLen == 2 {
		if (*pq)[1].dist < (*pq)[0].dist {
			(*pq)[0], (*pq)[1] = (*pq)[1], (*pq)[0]
		}
	} else if queueLen > 2 {
		index := queueLen - 1
		parent := (index - 1) / 2
		for (*pq)[index].dist < (*pq)[parent].dist {
			(*pq)[index], (*pq)[parent] = (*pq)[parent], (*pq)[index]
			index = parent
			parent = (index - 1) / 2
		}
	}
}

func (pq *priorityqueue) removeMin() *item {
	var result *item
	if queueLen != 0 {
		result = (*pq)[0]
		(*pq)[0] = (*pq)[queueLen-1]
		queueLen -= 1
		if queueLen > 0 {
			index := 0
			count := queueLen - 1
			for {
				child1 := 2*index + 1
				child2 := 2*index + 2

				if child1 > count {
					child1 = index
				}
				if child2 > count {
					child2 = index
				}

				if (*pq)[index].dist <= (*pq)[child1].dist && (*pq)[index].dist <= (*pq)[child2].dist {
					break
				}
				swapChild := child2
				if (*pq)[child1].dist < (*pq)[child2].dist {
					swapChild = child1
				}
				(*pq)[index], (*pq)[swapChild] = (*pq)[swapChild], (*pq)[index]
				index = swapChild
			}
		}
	}
	return result
}

func (pq *priorityqueue) replace(ind int, dist uint64) {
	index := 0
	for i, item := range *pq {
		if item.ind == ind {
			item.dist = dist
			index = i
			break
		}
	}
	if queueLen > 1 {
		parent := (index - 1) / 2
		for (*pq)[index].dist < (*pq)[parent].dist {
			(*pq)[index], (*pq)[parent] = (*pq)[parent], (*pq)[index]
			index = parent
			parent = (index - 1) / 2
		}
	}
}

func (pq priorityqueue) isEmpty() bool {
	return queueLen == 0
}

func dijkstra(graph *graph, pq *priorityqueue) uint64 {
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
	return graph.dist[graph.stop-1]
}

func main() {
	fin, _ := os.Open("pathmgep.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	info := strings.Fields(scanner.Text())
	verts, _ := strconv.Atoi(info[0])
	start, _ := strconv.Atoi(info[1])
	stop, _ := strconv.Atoi(info[2])

	graph := graph{
		verts:     verts,
		start:     start,
		stop:      stop,
		adjMatrix: make([][]uint64, verts),
		visited:   make([]bool, verts),
		dist:      make([]uint64, verts),
	}

	pq := make(priorityqueue, graph.verts)
	for i := 0; i < graph.verts; i++ {
		if i != graph.start-1 {
			graph.dist[i] = intInf
		}
		pq.insert(i, graph.dist[i])
		graph.adjMatrix[i] = make([]uint64, graph.verts)
		scanner.Scan()
		weights := strings.Fields(scanner.Text())
		for index, elem := range weights {
			if elem != "-1" {
				graph.adjMatrix[i][index], _ = strconv.ParseUint(elem, 10, 64)
			} else {
				graph.adjMatrix[i][index] = intInf - 1
			}
		}
	}

	fout, _ := os.Create("pathmgep.out")
	defer fout.Close()
	result := dijkstra(&graph, &pq)
	if result == intInf {
		fmt.Fprint(fout, -1)
		return
	}
	fmt.Fprint(fout, result)
}
