package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var stack = make([]int, 150)
var index = 0

func add() {
	stack[index-2] = stack[index-2] + stack[index-1]
	stack[index-1] = 0
	index -= 1
}

func subtract() {
	stack[index-2] = stack[index-2] - stack[index-1]
	stack[index-1] = 0
	index -= 1
}

func multiply() {
	stack[index-2] = stack[index-2] * stack[index-1]
	stack[index-1] = 0
	index -= 1
}

func main() {
	data_raw, _ := ioutil.ReadFile("postfix.in")
	data := strings.Split(string(data_raw), " ")

	for _, elem := range data {
		if num, err := strconv.Atoi(elem); err == nil {
			stack[index] = num
			index += 1
		} else {
			if elem == "-" {
				subtract()
			} else if elem == "+" {
				add()
			} else if elem == "*" {
				multiply()
			}
		}
	}
	fmt.Println(stack)
	fout, _ := os.Create("postfix.out")
	fout.WriteString(fmt.Sprint(stack[0]))
	fout.Close()
}
