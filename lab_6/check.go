package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	fin, _ := os.Open("check.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
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

	result := "YES"
	for i := 0; i < n; i++ {
		elem := array[i]
		if elem[1] != 0 {
			if array[elem[1]-1][0] > elem[0] {
				result = "NO"
				break
			}
		}

		if elem[2] != 0 {
			if array[elem[2]-1][0] < elem[0] {
				result = "NO"
				break
			}
		}
	}

	fout, _ := os.Create("check.out")
	fout.WriteString(result)
	fout.Close()
}
