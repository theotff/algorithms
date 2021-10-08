package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fin, _ := os.Open("queue.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)

	queue := make([]int, n)
	l_index := 0
	r_index := 0
	var results []int

	for scanner.Scan() {
		operation := scanner.Text()
		if len(operation) == 1 {
			results = append(results, queue[l_index])
			queue[l_index] = 0
			l_index += 1
		} else {
			var num int
			fmt.Sscanf(operation, "+ %d", &num)
			queue[r_index] = num
			r_index += 1
		}
	}
	result_arr := make([]string, len(results))

	for i := 0; i < len(results); i++ {
		result_arr[i] = fmt.Sprint(results[i])
	}

	fout, _ := os.Create("queue.out")
	fout.WriteString(strings.Join(result_arr, "\n"))
	fout.Close()
}
