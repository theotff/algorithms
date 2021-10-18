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
	prev *Node
}

type Dll struct {
	last *Node
}

func (list *Dll) insert(val int) {
	if !(list.exists(val)) {
		node := &Node{data: val, next: list.last, prev: nil}
		if list.last != nil {
			list.last.prev = node
		}
		list.last = node
	}
}

func (list *Dll) delete(val int) {
	node := list.last
	for node != nil {
		if node.data == val {
			if node.prev != nil {
				if node.next != nil {
					node.prev.next, node.next.prev = node.next, node.prev
				} else {
					node.prev.next = nil
				}
			} else {
				if node.next != nil {
					node.next.prev = nil
					list.last = node.next
				} else {
					list.last = nil
				}
			}
			return
		} else {
			node = node.next
		}
	}
}

func (list *Dll) exists(val int) bool {
	if list != nil {
		node := list.last
		for node != nil {
			if node.data == val {
				return true
			} else {
				node = node.next
			}
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
	table := make([]Dll, mod)
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
	str_results := make([]string, len(results))
	for index, elem := range results {
		str_results[index] = strconv.FormatBool(elem)
	}

	fout, _ := os.Create("set.out")
	fout.WriteString(strings.Join(str_results, "\n"))
	fout.Close()
}
