package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) insert(val int) {
	if !(list.exists(val)) {
		node := &Node{data: val, next: list.last}
		list.last = node
	}
}

func (list *LinkedList) delete(val int) {
	node := list.last
	for node != nil {
		if node.next != nil && node.next.data == val {
			if node.next.next != nil {
				node.next = node.next.next
			} else {
				node.next = nil
			}
			return

		} else if node.next != nil && node.next.data != val {
			node = node.next

		} else {
			if node.data == val {
				list.last = nil
			}
			return
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
	return int(math.Abs(float64(n % 100)))
}

func main() {
	fin, _ := os.Open("set.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	mod := 10000
	table := make([]LinkedList, mod)
	var results []bool

	for scanner.Scan() {
		txt := scanner.Text()

		switch {
		case strings.HasPrefix(txt, "insert"):
			var n int
			fmt.Sscanf(txt, "insert %d", &n)
			table[hash(n, mod)].insert(n)

		case strings.HasPrefix(txt, "delete"):
			var n int
			fmt.Sscanf(txt, "delete %d", &n)
			table[hash(n, mod)].delete(n)

		case strings.HasPrefix(txt, "exists"):
			var n int
			fmt.Sscanf(txt, "exists %d", &n)
			results = append(results, table[hash(n, mod)].exists(n))
		}
	}
	strResults := make([]string, len(results))
	for index, elem := range results {
		strResults[index] = strconv.FormatBool(elem)
	}

	fout, _ := os.Create("set.out")
	fout.WriteString(strings.Join(strResults, "\n"))
	fout.Close()
}
