package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) insert(val int) {
	if !(list.exists(val)) {
		node := &Node{
			data: val,
			next: list.last,
			prev: nil}

		if list.last != nil {
			list.last.prev = node
		}
		list.last = node
	}
}

func (list *LinkedList) delete(val int) {
	node := list.last
	for node != nil {
		if node.data == val {
			if node.next != nil {
				node.next.prev = node.prev
			}
			if node.prev != nil {
				node.prev.next = node.next
			} else {
				list.last = node.next
			}
			return
		} else {
			node = node.next
		}
	}
}

func (list *LinkedList) exists(val int) bool {
	node := list.last

	for node != nil {
		if node.data == val {
			return true
		} else {
			node = node.next
		}
	}
	return false
}

func hash(n int, mod int) int {
	return int(math.Abs(float64(n % mod)))
}

func main() {
	fin, _ := os.Open("set.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	mod := 10000
	table := make([]LinkedList, mod)
	var results []string

	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Fields(txt)
		n, _ := strconv.Atoi(fields[1])
		hashSum := hash(n, mod)

		switch fields[0] {
		case "insert":
			table[hashSum].insert(n)

		case "delete":
			table[hashSum].delete(n)

		case "exists":
			results = append(results, strconv.FormatBool(table[hashSum].exists(n)))
		}
	}

	fout, _ := os.Create("set.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
