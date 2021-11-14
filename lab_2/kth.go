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
		leftCount, rightCount := l, r

		for leftCount <= rightCount {
			for array[leftCount] < pivot {
				leftCount += 1
			}
			for array[rightCount] > pivot {
				rightCount -= 1
			}

			if leftCount <= rightCount {
				array[leftCount], array[rightCount] = array[rightCount], array[leftCount]
				leftCount += 1
				rightCount -= 1
			}

		}

		if k <= rightCount {
			return kth(array, l, rightCount, k)
		} else {
			return kth(array, leftCount, r, k)
		}
	}
}

func main() {
	var n, k int
	var a, b, c, arr1, arr2 int32

	dataRaw, _ := ioutil.ReadFile("kth.in")
	data := strings.Split(string(dataRaw), "\n")
	fmt.Sscanf(data[0], "%d %d", &n, &k)
	fmt.Sscanf(data[1], "%d %d %d %d %d", &a, &b, &c, &arr1, &arr2)
	array := make([]int32, n)
	array[0], array[1] = arr1, arr2

	for i := 2; i < n; i++ {
		array[i] = a*array[i-2] + b*array[i-1] + c
	}

	res := kth(array, 0, len(array)-1, k-1)
	fout, _ := os.Create("kth.out")
	fout.WriteString(fmt.Sprintf("%d", res[k-1]))
	fout.Close()
}
