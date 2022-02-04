package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	verts     int
	edges     int
	adjMatrix [][]int
}

func main() {
	var fin, fout *os.File
	fin, err := os.Open("input.txt")

	if errors.Is(err, os.ErrNotExist) {
		fin, _ = os.Open("parallel_edges.in")
		fout, _ = os.Create("parallel_edges.out")
	} else {
		fout, _ = os.Create("output.txt")
	}

	graph := &Graph{}
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	info := strings.Fields(scanner.Text())
	graph.verts, _ = strconv.Atoi(info[0])
	graph.edges, _ = strconv.Atoi(info[1])

	graph.adjMatrix = make([][]int, graph.verts)
	for i := 0; i < graph.verts; i++ {
		graph.adjMatrix[i] = make([]int, graph.verts)
	}

	result := "NO"
	for i := 0; i < graph.edges; i++ {
		scanner.Scan()
		verts := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(verts[0])
		b, _ := strconv.Atoi(verts[1])

		if graph.adjMatrix[a-1][b-1] == 1 {
			result = "YES"
			break
		}
		graph.adjMatrix[a-1][b-1] = 1
		graph.adjMatrix[b-1][a-1] = 1
	}

	fmt.Fprintln(fout, result)
}
