package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) insert(data int) {
	if list.last == nil {
		node := &Node{data: data}
		list.last = node
	} else {
		node := &Node{data: data, next: list.last}
		list.last = node
	}
}

func (list *LinkedList) remove() int {
	value := list.last.data
	list.last = list.last.next
	return value
}

func main() {
	var n int
	fin, _ := os.Open("stack.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)

	var results []int
	list := &LinkedList{}

	for scanner.Scan() {
		operation := scanner.Text()
		if len(operation) == 1 {
			results = append(results, list.remove())
		} else {
			var num int
			fmt.Sscanf(operation, "+ %d", &num)
			list.insert(num)
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
