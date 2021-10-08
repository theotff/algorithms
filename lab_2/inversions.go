package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func mergesort(array []int) ([]int, int) {
	if len(array) < 2 {
		return array, 0
	}

	half := len(array) / 2
	left, l_inversions := mergesort(array[:half])
	right, r_inversions := mergesort(array[half:])
	var result []int
	inversions := 0 + l_inversions + r_inversions

	l_index, r_index := 0, 0

	for l_index < len(left) && r_index < len(right) {
		if left[l_index] <= right[r_index] {
			result = append(result, left[l_index])
			l_index += 1
		} else {
			result = append(result, right[r_index])
			r_index += 1
			inversions += len(left) - l_index
		}
	}

	result = append(result, left[l_index:]...)
	result = append(result, right[r_index:]...)

	return result, inversions
}

func main() {
	var n int
	data_raw, _ := ioutil.ReadFile("inversions.in")
	data := strings.Split(string(data_raw), "\n")
	fmt.Sscanf(data[0], "%d", &n)
	array := make([]int, n)
	numbers_raw := strings.Split(data[1], " ")

	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(numbers_raw[i])
	}

	_, result := mergesort(array)
	fout, _ := os.Create("inversions.out")
	fout.WriteString(fmt.Sprint(result))
	fout.Close()
}
