package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func index(array []float64, n float64) int {
	elem_index := 0
	for index, elem := range array {
		if elem == n {
			elem_index = index
			break
		}
	}
	return elem_index
}

func main() {
	data_raw, _ := ioutil.ReadFile("sortland.in")
	data := strings.Split(string(data_raw), "\n")
	people, _ := strconv.Atoi(data[0])
	numbers_raw := strings.Split(data[1], " ")
	numbers := make([]float64, people)

	for index, elem := range numbers_raw {
		numbers[index], _ = strconv.ParseFloat(elem, 64)
	}

	sortlist := make([]float64, people)
	copy(sortlist, numbers)

	for i := 1; i < people; i++ {
		for j := i; j > 0; j-- {
			if sortlist[j] <= sortlist[j-1] {
				sortlist[j], sortlist[j-1] = sortlist[j-1], sortlist[j]
			}
		}
	}

	results := []string{
		fmt.Sprint(index(numbers, sortlist[0]) + 1),
		fmt.Sprint(index(numbers, sortlist[people/2]) + 1),
		fmt.Sprint(index(numbers, sortlist[people-1]) + 1)}

	fout, _ := os.Create("sortland.out")
	fout.WriteString(strings.Join(results, " "))
	fout.Close()
}
