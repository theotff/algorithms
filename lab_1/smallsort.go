package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	data_raw, _ := ioutil.ReadFile("smallsort.in")
	data := strings.Split(string(data_raw), "\n")
	n, _ := strconv.Atoi(data[0])
	numbers_raw := strings.Split(data[1], " ")
	numbers := make([]int, n)

	for index, elem := range numbers_raw {
		numbers[index], _ = strconv.Atoi(elem)
	}

	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if numbers[j] < numbers[j-1] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			} else {
				break
			}
		}
	}

	fout, _ := os.Create("smallsort.out")
	fout.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), " "), "[]"))
	fout.Close()
}
