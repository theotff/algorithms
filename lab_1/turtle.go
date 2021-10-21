package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	data_raw, _ := ioutil.ReadFile("turtle.in")
	data := strings.Split(string(data_raw), "\n")
	var h, w int
	fmt.Sscanf(data[0], "%d %d", &h, &w)
	var matrix [][]int

	for i := 1; i <= h; i++ {
		arr := strings.Split(data[i], " ")
		nums := make([]int, w)
		for j := 0; j < w; j++ {
			nums[j], _ = strconv.Atoi(arr[j])
		}
		matrix = append(matrix, nums)
	}

	for i := h - 1; i > -1; i-- {
		for j := 0; j < w; j++ {
			if i == h-1 {
				if j != 0 {
					matrix[i][j] += matrix[i][j-1]
				}
			} else {
				if j == 0 {
					matrix[i][j] += matrix[i+1][j]
				} else {
					if matrix[i][j-1] > matrix[i+1][j] {
						matrix[i][j] += matrix[i][j-1]
					} else {
						matrix[i][j] += matrix[i+1][j]
					}
				}
			}
		}
	}

	fout, _ := os.Create("turtle.out")
	fout.WriteString(fmt.Sprint(matrix[0][w-1]))
	fout.Close()
}
