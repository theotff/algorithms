package main

import (
	"bufio"
	"os"
	"strings"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) insert(data string) {
	if list.last == nil {
		node := &Node{data: data}
		list.last = node
	} else {
		if list.last.data == "(" && data == ")" {
			list.last = list.last.next

		} else if list.last.data == "[" && data == "]" {
			list.last = list.last.next

		} else {
			node := &Node{data: data, next: list.last}
			list.last = node
		}
	}
}

func (list *LinkedList) check_state() bool {
	if list.last == nil {
		return true
	} else {
		return false
	}
}

func main() {
	fin, _ := os.Open("brackets.in")
	scanner := bufio.NewScanner(fin)
	var results []string

	for scanner.Scan() {
		brackets := strings.Split(scanner.Text(), "")
		list := &LinkedList{}

		for _, elem := range brackets {
			list.insert(elem)
		}

		if list.check_state() {
			results = append(results, "YES")
		} else {
			results = append(results, "NO")
		}
	}
	fout, _ := os.Create("brackets.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
