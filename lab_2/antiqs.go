package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func antiqs(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		array[i], array[i/2] = array[i/2], array[i]
	}
	return array
}

func main() {
	var n int
	data, _ := ioutil.ReadFile("antiqs.in")
	fmt.Sscanf(string(data), "%d", &n)
	array := make([]int, n)

	for i := 0; i < n; i++ {
		array[i] = i + 1
	}

	antisorted := antiqs(array)
	result_raw := make([]string, n)

	for i := 0; i < n; i++ {
		result_raw[i] = fmt.Sprint(antisorted[i])
	}

	result := strings.Join(result_raw, " ")
	fout, _ := os.Create("antiqs.out")
	fout.WriteString(result)
	fout.Close()
}
