package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func search(left float64, right float64, a float64, n int) float64 {
	elem := (left + right) / 2
	array := make([]float64, n)
	array[0], array[1] = a, elem
	min := a

	for i := 2; i < n; i++ {
		newElem := 2*array[i-1] + 2 - array[i-2]
		array[i] = newElem

		if newElem < min {
			min = newElem
		}
	}

	if math.Abs(min) < 0.00001 {
		return array[n-1]
	} else if min < 0 {
		return search(elem, right, a, n)
	} else {
		return search(left, elem, a, n)
	}
}

func main() {
	var n int
	data, _ := ioutil.ReadFile("garland.in")
	numberStrings := strings.Fields(string(data))
	fmt.Sscanf(numberStrings[0], "%d", &n)
	a, _ := strconv.ParseFloat(numberStrings[1], 64)
	result := search(0, a, a, n)

	fout, _ := os.Create("garland.out")
	resultString := fmt.Sprintf("%.2f", result)
	fout.WriteString(resultString)
	fout.Close()
}
