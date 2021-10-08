package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func kth(array []int32, l int, r int, k int) []int32 {
	if l >= r {
		return array
	} else {
		pivot := array[(l+r)/2]
		l_count, r_count := l, r
		for l_count <= r_count {
			for array[l_count] < pivot {
				l_count += 1
			}
			for array[r_count] > pivot {
				r_count -= 1
			}
			if l_count <= r_count {
				array[l_count], array[r_count] = array[r_count], array[l_count]
				l_count += 1
				r_count -= 1
			}
		}
		if k <= r_count {
			return kth(array, l, r_count, k)
		} else {
			return kth(array, l_count, r, k)
		}

	}
}

func main() {
	var n, k int
	var a, b, c, arr_1, arr_2 int32

	data_raw, _ := ioutil.ReadFile("kth.in")
	data := strings.Split(string(data_raw), "\n")
	fmt.Sscanf(data[0], "%d %d", &n, &k)
	fmt.Sscanf(data[1], "%d %d %d %d %d", &a, &b, &c, &arr_1, &arr_2)
	array := make([]int32, n)
	array[0], array[1] = arr_1, arr_2

	for i := 2; i < n; i++ {
		array[i] = a*array[i-2] + b*array[i-1] + c
	}

	res := kth(array, 0, len(array)-1, k-1)
	fout, _ := os.Create("kth.out")
	fout.WriteString(fmt.Sprintf("%d", res[k-1]))
	fout.Close()
}
