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

func sum(array []int) int {
	result := 0
	for _, elem := range array {
		result += elem
	}
	return result
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

	graphSum := 0

	for i := 0; i < graph.verts; i++ {
		scanner.Scan()
		data := strings.Fields(scanner.Text())
		for j := 0; j < graph.verts; j++ {
			graph.adjMatrix[i][j], _ = strconv.Atoi(data[j])
		}
		graphSum += sum(graph.adjMatrix[i])
	}

	result := "NO"
	if graphSum >= graph.verts*2 && graphSum%2 == 0 {
		result = "YES"
	}
	fmt.Fprintln(fout, result)
}