package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) insert(value int) {
	if !(list.exists(value)) {
		node := &Node{
			value: value,
			next:  list.last,
			prev:  nil}

		if list.last != nil {
			list.last.prev = node
		}
		list.last = node
	}
}

func (list *LinkedList) delete(value int) {
	node := list.get(value)
	if node != nil {
		if node.next != nil {
			node.next.prev = node.prev
		}
		if node.prev != nil {
			node.prev.next = node.next
		} else {
			list.last = node.next
		}
	}
}

func (list *LinkedList) get(value int) *Node {
	node := list.last
	for node != nil {
		if node.value == value {
			return node
		} else {
			node = node.next
		}
	}
	return nil
}

func (list *LinkedList) exists(value int) bool {
	return list.get(value) != nil
}

func hash(num int, mod int) int {
	return int(math.Abs(float64(num % mod)))
}

func main() {
	fin, _ := os.Open("set.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	mod := 10000
	table := make([]LinkedList, mod)
	var results []string

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		num, _ := strconv.Atoi(fields[1])
		hashSum := hash(num, mod)

		switch fields[0] {
		case "insert":
			table[hashSum].insert(num)

		case "delete":
			table[hashSum].delete(num)

		case "exists":
			exists := table[hashSum].exists(num)
			results = append(results, strconv.FormatBool(exists))
		}
	}

	fout, _ := os.Create("set.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
