package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func index(array []float64, n float64) int {
	elemIndex := 0
	for index, elem := range array {
		if elem == n {
			elemIndex = index
			break
		}
	}
	return elemIndex
}

func main() {
	dataRaw, _ := ioutil.ReadFile("sortland.in")
	data := strings.Split(string(dataRaw), "\n")
	people, _ := strconv.Atoi(data[0])
	numbersRaw := strings.Split(data[1], " ")
	numbers := make([]float64, people)

	for index, elem := range numbersRaw {
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
