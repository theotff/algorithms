package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

func (s *Stack) add() {
	val1 := s.pop()
	val2 := s.pop()
	s.push(val2 + val1)
}

func (s *Stack) subtract() {
	val1 := s.pop()
	val2 := s.pop()
	s.push(val2 - val1)
}

func (s *Stack) multiply() {
	val1 := s.pop()
	val2 := s.pop()
	s.push(val2 * val1)
}

func main() {
	dataRaw, _ := ioutil.ReadFile("postfix.in")
	data := strings.Fields(string(dataRaw))

	stack := &Stack{}

	for _, elem := range data {
		if num, err := strconv.Atoi(elem); err == nil {
			stack.push(num)
		} else {
			switch elem {
			case "-":
				stack.subtract()
			case "+":
				stack.add()
			case "*":
				stack.multiply()
			}
		}
	}
	fout, _ := os.Create("postfix.out")
	fmt.Fprintln(fout, stack.last.data)
	fout.Close()
}
