package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fin, _ := os.Open("stack.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)

	stack := make([]int, n)
	stack_index := 0
	var results []int

	for scanner.Scan() {
		operation := scanner.Text()
		if len(operation) == 1 {
			results = append(results, stack[stack_index])
			stack[stack_index] = 0
			stack_index -= 1
		} else {
			var num int
			fmt.Sscanf(operation, "+ %d", &num)
			stack_index += 1
			stack[stack_index] = num
		}
	}
	result_arr := make([]string, len(results))

	for i := 0; i < len(results); i++ {
		result_arr[i] = fmt.Sprint(results[i])
	}

	fout, _ := os.Create("stack.out")
	fout.WriteString(strings.Join(result_arr, "\n"))
	fout.Close()
}
