package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//func countingsort(array []int) {
//	counts := make()
//}

func radixsort(array []int, n int, m int) []int {
	for i := 1; i < m; i++ {

	}
	return array
}

func main() {
	var n, m, k int
	var array []string
	data_raw, _ := ioutil.ReadFile("radixsort.in")
	data := strings.Split(string(data_raw), "\n")
	fmt.Sscanf(data[0], "%d %d %d", &n, &m, &k)
	for i := 1; i <= n; i++ {
		array = append(array, data[i])
	}
	fmt.Println(array)
}
