package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(array []int) int {
	result := 0
	for _, elem := range array {
		result += elem
	}
	return result
}

func main() {
	fin, err := os.Open("input.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fin, _ = os.Open("is_oriented.in")
		}
	}
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	vertices, _ := strconv.Atoi(scanner.Text())

	adj_matrix := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		adj_matrix[i] = make([]int, vertices)
	}

	graph_sum := 0

	for i := 0; i < vertices; i++ {
		scanner.Scan()
		data := strings.Fields(scanner.Text())
		for j := 0; j < vertices; j++ {
			adj_matrix[i][j], _ = strconv.Atoi(data[j])
		}
		graph_sum += sum(adj_matrix[i])
	}

	var fout *os.File
	if err == nil {
		fout, _ = os.Create("output.txt")
	} else {
		if errors.Is(err, os.ErrNotExist) {
			fout, _ = os.Create("is_oriented.out")
		}
	}

	result := "NO"
	if graph_sum >= vertices*2 && graph_sum%2 == 0 {
		result = "YES"
	}
	fmt.Fprintf(fout, result)
}
