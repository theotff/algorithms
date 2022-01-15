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
	fin, err := os.Open("input.txt")
	if errors.Is(err, os.ErrNotExist) {
		fin, _ = os.Open("adj_matrix.in")
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

	for i := 0; i < graph.edges; i++ {
		scanner.Scan()
		vertices := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(vertices[0])
		b, _ := strconv.Atoi(vertices[1])

		graph.adjMatrix[a-1][b-1] = 1
	}

	var fout *os.File
	if err == nil {
		fout, _ = os.Create("output.txt")
	} else if errors.Is(err, os.ErrNotExist) {
		fout, _ = os.Create("adj_matrix.out")
	}

	for _, line := range graph.adjMatrix {
		arr := make([]string, graph.verts)
		for index, elem := range line {
			arr[index] = fmt.Sprint(elem)
		}
		fmt.Fprintln(fout, strings.Join(arr, " "))
	}
}
