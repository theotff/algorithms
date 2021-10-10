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
	data_raw, _ := ioutil.ReadFile("binsearch.in")
	data := strings.Split(string(data_raw), "\n")
	fmt.Sscanf(data[0], "%d", &n)
	fmt.Sscanf(data[2], "%d", &m)
	array := make([]int, n)
	requests := make([]int, m)

	array_index := 0
	requests_index := 0
	for _, elem := range strings.Fields(data[1]) {
		if num, err := strconv.Atoi(elem); err == nil {
			array[array_index] = num
			array_index += 1
		}
	}
	for _, elem := range strings.Fields(data[3]) {
		if num, err := strconv.Atoi(elem); err == nil {
			requests[requests_index] = num
			requests_index += 1
		}
	}

	var results []string

	for _, elem := range requests {
		f_index := first(array, elem)
		l_index := last(array, elem)

		if f_index != -1 && l_index != -1 {
			var str_arr = []string{fmt.Sprint(f_index + 1), fmt.Sprint(l_index + 1)}
			results = append(results, strings.Join(str_arr, " "))
		} else {
			var str_arr = []string{fmt.Sprint(f_index), fmt.Sprint(l_index)}
			results = append(results, strings.Join(str_arr, " "))
		}
	}
	fout, _ := os.Create("binsearch.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
