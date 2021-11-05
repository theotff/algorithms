package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fin, _ := os.Open("height.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	array := make([][]int, n)

	for index := 0; index < n; index++ {
		scanner.Scan()
		elems := strings.Fields(scanner.Text())
		arr := make([]int, 4)
		for i := 0; i < 4; i++ {
			if i != 3 {
				arr[i], _ = strconv.Atoi(elems[i])
			} else {
				arr[i] = 1
			}
		}
		array[index] = arr
	}

	len := 0

	for i := 0; i < n; i++ {
		elem := array[i]
		if elem[1] != 0 {
			array[elem[1]-1][3] += elem[3]
		}

		if elem[2] != 0 {
			array[elem[2]-1][3] += elem[3]
		}

		if elem[3] > len {
			len = elem[3]
		}
	}

	fout, _ := os.Create("height.out")
	fout.WriteString(fmt.Sprint(len))
	fout.Close()
}
