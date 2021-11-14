package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func mergeSort(array []int) ([]int, int) {
	if len(array) < 2 {
		return array, 0
	}

	half := len(array) / 2
	left, leftInversions := mergeSort(array[:half])
	right, rightInversions := mergeSort(array[half:])
	var result []int
	inversions := 0 + leftInversions + rightInversions

	leftIndex, rightIndex := 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex += 1
		} else {
			result = append(result, right[rightIndex])
			rightIndex += 1
			inversions += len(left) - leftIndex
		}
	}

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result, inversions
}

func main() {
	var n int
	dataRaw, _ := ioutil.ReadFile("inversions.in")
	data := strings.Split(string(dataRaw), "\n")
	fmt.Sscanf(data[0], "%d", &n)
	array := make([]int, n)
	numbersRaw := strings.Split(data[1], " ")

	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(numbersRaw[i])
	}

	_, result := mergeSort(array)
	fout, _ := os.Create("inversions.out")
	fout.WriteString(fmt.Sprint(result))
	fout.Close()
}
