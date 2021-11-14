package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func first(array []int, target int) int {
	min := 0
	max := len(array) - 1
	result := -1

	for min <= max {
		mid := (min + max) / 2
		if target < array[mid] {
			max = mid - 1
		} else if target > array[mid] {
			min = mid + 1
		} else {
			result = mid
			max = mid - 1
		}
	}

	return result
}

func last(array []int, target int) int {
	min := 0
	max := len(array) - 1
	result := -1

	for min <= max {
		mid := (min + max) / 2
		if target < array[mid] {
			max = mid - 1
		} else if target > array[mid] {
			min = mid + 1
		} else {
			result = mid
			min = mid + 1
		}
	}

	return result
}

func main() {
	var n, m int
	dataRaw, _ := ioutil.ReadFile("binsearch.in")
	data := strings.Split(string(dataRaw), "\n")
	fmt.Sscanf(data[0], "%d", &n)
	fmt.Sscanf(data[2], "%d", &m)
	array := make([]int, n)
	requests := make([]int, m)

	arrayIndex := 0
	requestsIndex := 0

	for _, elem := range strings.Fields(data[1]) {
		if num, err := strconv.Atoi(elem); err == nil {
			array[arrayIndex] = num
			arrayIndex += 1
		}
	}

	for _, elem := range strings.Fields(data[3]) {
		if num, err := strconv.Atoi(elem); err == nil {
			requests[requestsIndex] = num
			requestsIndex += 1
		}
	}

	var results []string

	for _, elem := range requests {
		firstIndex := first(array, elem)
		lastIndex := last(array, elem)

		if firstIndex != -1 && lastIndex != -1 {
			var strArr = []string{fmt.Sprint(firstIndex + 1), fmt.Sprint(lastIndex + 1)}
			results = append(results, strings.Join(strArr, " "))
		} else {
			var strArr = []string{fmt.Sprint(firstIndex), fmt.Sprint(lastIndex)}
			results = append(results, strings.Join(strArr, " "))
		}
	}

	fout, _ := os.Create("binsearch.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
