package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quickSort(array []string, k int) []string {
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
		result := append(quickSort(less, k), pivot)
		result = append(result, quickSort(greater, k)...)
		return result
	}
}

func radixsort(array []string, m int, k int) []string {
	var result = array
	for i := 0; i < k; i++ {
		result = quickSort(result, m-i-1)
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

	result := radixsort(array, m, k)
	fout, _ := os.Create("radixsort.out")
	fout.WriteString(strings.Join(result, "\n"))
	fout.Close()
}
