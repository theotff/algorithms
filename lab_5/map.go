package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

type Node struct {
	key   string
	value string
	next  *Node
	prev  *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) insert(key string, value string) {
	result := list.get(key)
	if result == nil {
		node := &Node{
			key:   key,
			value: value,
			next:  list.last,
			prev:  nil}

		if list.last != nil {
			list.last.prev = node
		}
		list.last = node
	} else {
		result.value = value
	}
}

func (list *LinkedList) delete(key string) {
	node := list.last
	for node != nil {
		if node.key == key {
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

func (list *LinkedList) get(key string) *Node {
	node := list.last
	for node != nil {
		if node.key == key {
			return node
		} else {
			node = node.next
		}
	}
	return nil
}

func hash(key string, mod int) int {
	hashSum := 5381
	for _, elem := range key {
		hashSum = ((hashSum << 5) + hashSum) + int(elem)
	}
	return int(math.Abs(float64(hashSum % mod)))
}

func main() {
	fin, _ := os.Open("map.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	mod := 10000
	table := make([]LinkedList, mod)
	var results []string

	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Fields(txt)
		key := fields[1]
		hashSum := hash(key, mod)

		switch fields[0] {
		case "put":
			value := fields[2]
			table[hashSum].insert(key, value)

		case "get":
			node := table[hashSum].get(key)
			if node != nil {
				results = append(results, node.value)
			} else {
				results = append(results, "none")
			}

		case "delete":
			table[hashSum].delete(key)
		}
	}

	fout, _ := os.Create("map.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
