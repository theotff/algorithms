package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	last *Node
}

func (s *Stack) push(data int) {
	node := &Node{data: data, next: s.last}
	s.last = node
}

func (s *Stack) pop() int {
	value := s.last.data
	s.last = s.last.next
	return value
}

func main() {
	fin, _ := os.Open("stack.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	var results []int
	s := &Stack{}

	for scanner.Scan() {
		operation := scanner.Text()

		if len(operation) == 1 {
			results = append(results, s.pop())

		} else {
			var num int
			fmt.Sscanf(operation, "+ %d", &num)
			s.push(num)
		}
	}

	fout, _ := os.Create("stack.out")
	for i := 0; i < len(results); i++ {
		fmt.Fprintln(fout, results[i])
	}
	fout.Close()
}
