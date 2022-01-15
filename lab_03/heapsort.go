package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func makeHeap(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		index := i
		for index != 0 {
			parent := (index - 1) / 2
			if array[index] <= array[parent] {
				break
			}
			array[index], array[parent] = array[parent], array[index]
			index = parent
		}
	}
	return array
}

func heapSort(array []int) []int {
	for count := len(array) - 1; count > 0; count-- {
		array[0], array[count] = array[count], array[0]

		index := 0
		for {
			child1 := 2*index + 1
			child2 := 2*index + 2

			if child1 >= count {
				child1 = index
			}
			if child2 >= count {
				child2 = index
			}

			if array[index] >= array[child1] && array[index] >= array[child2] {
				break
			}

			var swapChild int
			if array[child1] > array[child2] {
				swapChild = child1
			} else {
				swapChild = child2
			}
			array[index], array[swapChild] = array[swapChild], array[index]
			index = swapChild
		}
	}
	return array
}

func main() {
	var n int
	dataRaw, _ := ioutil.ReadFile("sort.in")
	data := strings.Split(string(dataRaw), "\n")
	fmt.Sscanf(data[0], "%d", &n)
	array := make([]int, n)
	numbersRaw := strings.Split(data[1], " ")
	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(numbersRaw[i])
	}

	heap := makeHeap(array)
	sorted := heapSort(heap)
	resultRaw := make([]string, n)
	for i := 0; i < n; i++ {
		resultRaw[i] = fmt.Sprint(sorted[i])
	}
	result := strings.Join(resultRaw, " ")
	fout, _ := os.Create("sort.out")
	fout.WriteString(result)
	fout.Close()
}
