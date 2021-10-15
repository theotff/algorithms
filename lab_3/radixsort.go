package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quicksort(array []string, k int) []string {
	n := len(array)
	if n < 2 {
		return array
	} else {
		index := len(array) / 2
		pivot := array[index]
		var less []string
		var greater []string

		for i := 0; i < n; i++ {
			if i != index {
				if array[i][k] < pivot[k] {
					less = append(less, array[i])
				} else if array[i][k] > pivot[k] {
					greater = append(greater, array[i])
				} else {
					if i < index {
						less = append(less, array[i])
					} else {
						greater = append(greater, array[i])
					}
				}
			}
		}
		result := append(quicksort(less, k), pivot)
		result = append(result, quicksort(greater, k)...)
		return result
	}
}

func radixsort(array []string, k int) []string {
	var result = array
	n := len(array)
	for i := 0; i < k; i++ {
		result = quicksort(result, n-i-1)
	}
	return result
}

func main() {
	var n, m, k int
	var array []string
	fin, _ := os.Open("radixsort.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d %d", &n, &m, &k)

	for scanner.Scan() {
		array = append(array, scanner.Text())
	}

	result := radixsort(array, k)
	fout, _ := os.Create("radixsort.out")
	fout.WriteString(strings.Join(result, "\n"))
	fout.Close()
}
