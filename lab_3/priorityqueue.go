package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func makeheap(array [][]int) [][]int {
	n := len(array)
	for i := 0; i < n; i++ {
		index := i
		for index != 0 {
			parent := (index - 1) / 2
			if array[index][0] >= array[parent][0] {
				break
			}
			array[index], array[parent] = array[parent], array[index]
			index = parent
		}
	}
	return array
}

func main() {
	fin, _ := os.Open("priorityqueue.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	var results []string
	var heap [][]int
	index := 0

	for scanner.Scan() {
		txt := scanner.Text()
		switch {
		case strings.HasPrefix(txt, "push"):
			var n int
			fmt.Sscanf(txt, "push %d", &n)
			arr := []int{n, index}
			heap = makeheap(append(heap, arr))
			index += 1

		case strings.HasPrefix(txt, "decrease-key"):
			var n, m int
			fmt.Sscanf(txt, "decrease-key %d %d", &n, &m)
			len := len(heap)

			for i := 0; i < len; i++ {
				if heap[i][1] == n-1 {
					heap[i][0] = m
					break
				}
			}
			heap = makeheap(heap)

		default:
			var result string
			if len(heap) >= 1 {
				result = fmt.Sprint(heap[0][0])
				heap = heap[1:]
			} else {
				result = "*"
			}
			results = append(results, result)
			heap = makeheap(heap)
		}
	}

	fout, _ := os.Create("priorityqueue.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
