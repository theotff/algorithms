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

type BracketStack struct {
	last *Node
}

func (bs *BracketStack) insert(data string) {
	if bs.last == nil {
		node := &Node{data: data}
		bs.last = node

	} else {
		if bs.last.data == "(" && data == ")" {
			bs.last = bs.last.next

		} else if bs.last.data == "[" && data == "]" {
			bs.last = bs.last.next

		} else {
			node := &Node{data: data, next: bs.last}
			bs.last = node
		}
	}
}

func (bs *BracketStack) checkState() bool {
	if bs.last == nil {
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
		bs := &BracketStack{}

		for _, elem := range brackets {
			bs.insert(elem)
		}

		if bs.checkState() {
			results = append(results, "YES")
		} else {
			results = append(results, "NO")
		}
	}
	fout, _ := os.Create("brackets.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
