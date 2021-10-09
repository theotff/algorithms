package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func makeheap(array []int) []int {
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

func heapsort(array []int) []int {
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

			var swap_child int
			if array[child1] > array[child2] {
				swap_child = child1
			} else {
				swap_child = child2
			}
			array[index], array[swap_child] = array[swap_child], array[index]
			index = swap_child
		}
	}
	return array
}

func main() {
	var n int
	data_raw, _ := ioutil.ReadFile("sort.in")
	data := strings.Split(string(data_raw), "\n")
	fmt.Sscanf(data[0], "%d", &n)
	array := make([]int, n)
	numbers_raw := strings.Split(data[1], " ")
	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(numbers_raw[i])
	}

	heap := makeheap(array)
	sorted := heapsort(heap)
	result_raw := make([]string, n)
	for i := 0; i < n; i++ {
		result_raw[i] = fmt.Sprint(sorted[i])
	}
	result := strings.Join(result_raw, " ")
	fout, _ := os.Create("sort.out")
	fout.WriteString(result)
	fout.Close()
}
