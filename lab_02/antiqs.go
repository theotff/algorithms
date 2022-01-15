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

	antiSorted := antiqs(array)
	resultRaw := make([]string, n)

	for i := 0; i < n; i++ {
		resultRaw[i] = fmt.Sprint(antiSorted[i])
	}

	result := strings.Join(resultRaw, " ")
	fout, _ := os.Create("antiqs.out")
	fout.WriteString(result)
	fout.Close()
}
