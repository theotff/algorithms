package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func push(heap [][]int, arr []int) [][]int {
	if len(heap) == 0 {
		heap = append(heap, arr)
	} else if len(heap) == 1 {
		heap = append(heap, arr)
		if heap[1][0] < heap[0][0] {
			heap[1], heap[0] = heap[0], heap[1]
		}
	} else {
		heap = append(heap, arr)
		index := len(heap) - 1
		parent := (index - 1) / 2
		for heap[index][0] < heap[parent][0] {
			heap[index], heap[parent] = heap[parent], heap[index]
			index = parent
			parent = (index - 1) / 2
		}
	}
	return heap
}

func replace(heap [][]int, n int, m int) [][]int {
	ln := len(heap)
	index := 0
	for i := 0; i < ln; i++ {
		if heap[i][1] == n-1 {
			heap[i][0] = m
			index = i
			break
		}
	}
	if len(heap) > 1 {
		parent := (index - 1) / 2
		for heap[index][0] < heap[parent][0] {
			heap[index], heap[parent] = heap[parent], heap[index]
			index = parent
			parent = (index - 1) / 2
		}
	}

	return heap
}

func removeMin(heap [][]int) ([][]int, string) {
	var result string
	ln := len(heap)
	if ln < 1 {
		result = "*"
		return heap, result
	} else {
		result = fmt.Sprint(heap[0][0])
		heap[0] = heap[ln-1]
		heap = heap[:ln-1]
		ln = len(heap)
		if ln > 0 {
			index := 0
			count := ln - 1
			for {
				child1 := 2*index + 1
				child2 := 2*index + 2

				if child1 > count {
					child1 = index
				}
				if child2 > count {
					child2 = index
				}

				if heap[index][0] <= heap[child1][0] && heap[index][0] <= heap[child2][0] {
					break
				}

				var swapChild int
				if heap[child1][0] < heap[child2][0] {
					swapChild = child1
				} else {
					swapChild = child2
				}
				heap[index], heap[swapChild] = heap[swapChild], heap[index]
				index = swapChild
			}
		}
		return heap, result
	}
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
			heap = push(heap, arr)

		case strings.HasPrefix(txt, "decrease-key"):
			var n, m int
			fmt.Sscanf(txt, "decrease-key %d %d", &n, &m)
			heap = replace(heap, n, m)

		case strings.HasPrefix(txt, "extract-min"):
			var result string
			heap, result = removeMin(heap)
			results = append(results, result)
		}
		index += 1
	}
	fout, _ := os.Create("priorityqueue.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
