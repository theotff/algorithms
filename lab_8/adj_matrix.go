package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fin, err := os.Open("input.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fin, _ = os.Open("adj_matrix.in")
		}
	}
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	info := strings.Fields(scanner.Text())
	vertices, _ := strconv.Atoi(info[0])
	edges, _ := strconv.Atoi(info[1])

	adj_matrix := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		adj_matrix[i] = make([]int, vertices)
	}

	for i := 0; i < edges; i++ {
		scanner.Scan()
		points := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(points[0])
		b, _ := strconv.Atoi(points[1])

		adj_matrix[a-1][b-1] = 1
	}

	var fout *os.File
	if err == nil {
		fout, _ = os.Create("output.txt")
	} else {
		if errors.Is(err, os.ErrNotExist) {
			fout, _ = os.Create("adj_matrix.out")
		}
	}

	for _, line := range adj_matrix {
		arr := make([]string, vertices)
		for index, elem := range line {
			arr[index] = fmt.Sprint(elem)
		}
		fmt.Fprintln(fout, strings.Join(arr, " "))
	}
}
