package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func checkheap(array []int) string {
	result := "YES"
	n := len(array)
	for i := 0; i < n; i++ {
		index := i
		for index != 0 {
			parent := (index - 1) / 2
			if array[index] >= array[parent] {
				break
			}
			result = "NO"
			break
		}
	}
	return result
}

func main() {
	var n int
	data_raw, _ := ioutil.ReadFile("isheap.in")
	data := strings.Split(string(data_raw), "\n")
	fmt.Sscanf(data[0], "%d", &n)
	array := make([]int, n)
	numbers_raw := strings.Split(data[1], " ")

	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(numbers_raw[i])
	}

	fout, _ := os.Create("isheap.out")
	fout.WriteString(checkheap(array))
	fout.Close()
}
