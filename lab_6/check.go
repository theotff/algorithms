package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func is_bst(array [][]int, index int, min int, max int) bool {
	if index == 0 {
		return true
	}

	index -= 1

	if min <= array[index][0] && array[index][0] <= max {
		return is_bst(array, array[index][1], min, array[index][0]-1) &&
			is_bst(array, array[index][2], array[index][0]+1, max)
	} else {
		return false
	}
}

func main() {
	fin, _ := os.Open("check.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	state := true

	if n > 0 {
		array := make([][]int, n)
		for index := 0; index < n; index++ {
			scanner.Scan()
			elems := strings.Fields(scanner.Text())
			arr := make([]int, 3)
			for i := 0; i < 3; i++ {
				arr[i], _ = strconv.Atoi(elems[i])
			}
			array[index] = arr
		}

		state = is_bst(array, 1, -1000000000, 1000000000)
	}

	fout, _ := os.Create("check.out")
	if state {
		fout.WriteString("YES")
	} else {
		fout.WriteString("NO")
	}
	fout.Close()
}
