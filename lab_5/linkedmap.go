package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

type Node struct {
	key     string
	value   string
	next    *Node
	prev    *Node
	prevIns *Node
	nextIns *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) put(key string, value string, prevIns *Node) *Node {
	result := list.get(key)
	if result == nil {
		node := &Node{
			key:     key,
			value:   value,
			next:    list.last,
			prevIns: prevIns,
			nextIns: nil}

		if list.last != nil {
			list.last.prev = node
		}
		if prevIns != nil {
			prevIns.nextIns = node
		}

		list.last = node
		return node
	} else {
		result.value = value
		return prevIns
	}
}

func (list *LinkedList) delete(key string, prevIns *Node) *Node {
	node := list.get(key)
	if node != nil {
		if node.next != nil {
			node.next.prev = node.prev
		}

		if node.prev != nil {
			node.prev.next = node.next
		} else {
			list.last = node.next
		}

		if node.prevIns != nil {
			node.prevIns.nextIns = node.nextIns
		}
		if node.nextIns != nil {
			node.nextIns.prevIns = node.prevIns
		}

		if node == prevIns {
			return prevIns.prevIns
		}
	}

	return prevIns
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
	fin, _ := os.Open("linkedmap.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	mod := 1000
	table := make([]LinkedList, mod)
	var results []string

	var prevIns *Node = nil
	none := "none"

	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Fields(txt)
		key := fields[1]
		hashSum := hash(key, mod)

		switch fields[0] {
		case "put":
			value := fields[2]
			prevIns = table[hashSum].put(key, value, prevIns)

		case "get":
			result := table[hashSum].get(key)
			if result == nil {
				results = append(results, none)
			} else {
				results = append(results, result.value)
			}

		case "delete":
			prevIns = table[hashSum].delete(key, prevIns)

		case "prev":
			result := table[hashSum].get(key)

			if result != nil && result.prevIns != nil {
				results = append(results, result.prevIns.value)
			} else {
				results = append(results, none)
			}

		case "next":
			result := table[hashSum].get(key)

			if result != nil && result.nextIns != nil {
				results = append(results, result.nextIns.value)
			} else {
				results = append(results, none)
			}

		}
	}

	fout, _ := os.Create("linkedmap.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
