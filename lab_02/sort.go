package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func quickSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	} else {
		index := rand.Intn(n)
		pivot := array[index]
		var less []int
		var greater []int

		for i := 0; i < n; i++ {
			if i != index {
				if array[i] <= pivot {
					less = append(less, array[i])
				} else {
					greater = append(greater, array[i])
				}
			}
		}

		result := append(quickSort(less), pivot)
		result = append(result, quickSort(greater)...)
		return result
	}
}

func main() {
	var n int
	data_raw, _ := ioutil.ReadFile("sort.in")
	data := strings.Split(string(data_raw), "\n")
	fmt.Sscanf(data[0], "%d", &n)
	array := make([]int, n)
	numbersRaw := strings.Split(data[1], " ")

	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(numbersRaw[i])
	}

	sorted := quickSort(array)
	resultRaw := make([]string, n)

	for i := 0; i < n; i++ {
		resultRaw[i] = fmt.Sprint(sorted[i])
	}

	result := strings.Join(resultRaw, " ")
	fout, _ := os.Create("sort.out")
	fout.WriteString(result)
	fout.Close()
}
