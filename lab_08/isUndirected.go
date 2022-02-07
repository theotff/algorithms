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
		fin, _ = os.Open("is_undirected.in")
		fout, _ = os.Create("is_undirected.out")
	} else {
		fout, _ = os.Create("output.txt")
	}

	graph := &Graph{}
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	graph.verts, _ = strconv.Atoi(scanner.Text())

	graph.adjMatrix = make([][]int, graph.verts)
	for i := 0; i < graph.verts; i++ {
		graph.adjMatrix[i] = make([]int, graph.verts)
	}

	result := "YES"

	for i := 0; i < graph.verts; i++ {
		scanner.Scan()
		data := strings.Fields(scanner.Text())
		for j := 0; j < graph.verts; j++ {
			graph.adjMatrix[i][j], _ = strconv.Atoi(data[j])
		}
	}

	for i := 0; i < graph.verts; i++ {
		if result == "NO" {
			break
		}
		for j := 0; j < graph.verts; j++ {
			if graph.adjMatrix[i][j] != graph.adjMatrix[j][i] || i == j && graph.adjMatrix[i][j] == 1 {
				result = "NO"
				break
			}
		}
	}

	fmt.Fprintln(fout, result)
}
