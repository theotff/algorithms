package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

func quicksort(array [][]string) [][]string {
	n := len(array)
	if n < 2 {
		return array
	} else {
		index := rand.Intn(n)
		pivot := array[index]
		var less [][]string
		var greater [][]string

		for i := 0; i < n; i++ {
			if i != index {
				if array[i][0] < pivot[0] {
					less = append(less, array[i])
				} else if array[i][0] > pivot[0] {
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

		result := append(quicksort(less), pivot)
		result = append(result, quicksort(greater)...)
		return result
	}
}

func main() {
	var n int
	data_raw, _ := ioutil.ReadFile("race.in")
	data := strings.Split(string(data_raw), "\n")
	fmt.Sscanf(data[0], "%d", &n)
	array := make([][]string, n)

	for i := 1; i <= n; i++ {
		array[i-1] = (strings.Split(data[i], " "))
	}

	result := quicksort(array)
	var result_final []string

	country := ""
	var names []string

	for i := 0; i < n; i++ {
		if country != result[i][0] {
			if country != "" {
				result_final = append(result_final, "=== "+country+" ==="+"\n"+strings.Join(names, "\n"))
			}
			country = result[i][0]
			names = nil
		}
		names = append(names, result[i][1])
	}
	if country != "" {
		result_final = append(result_final, "=== "+country+" ==="+"\n"+strings.Join(names, "\n"))
	}

	fout, _ := os.Create("race.out")
	fout.WriteString(strings.Join(result_final, "\n"))
	fout.Close()
}
