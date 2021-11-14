package main

import (
	"bufio"
	"os"
	"strings"
)

type Node struct {
	key     string
	value   string
	next    *Node
	prevIns *Node
	nextIns *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) insert(key string, value string, prev *Node) *Node {
	result := list.get(key)

	if result == nil {
		node := &Node{
			key:     key,
			value:   value,
			next:    list.last,
			prevIns: prev}

		list.last = node
		return node

	} else {
		return result
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

func (list *LinkedList) delete(key string) {
	node := list.last

	for node != nil {
		if node.next != nil && node.next.key == key {
			if node.next.next != nil {
				swapLinks(node)
				node.next = node.next.next
			} else {
				swapLinks(node)
				node.next = nil
			}
			return

		} else if node.next != nil && node.next.key != key {
			node = node.next

		} else {
			if node.key == key {
				swapLinks(node)
				list.last = nil
			}

			return
		}
	}
}

func swapLinks(node *Node) {
	if node.prevIns != nil && node.nextIns != nil {
		node.prevIns.nextIns, node.nextIns.prevIns = node.nextIns, node.prevIns

	} else if node.prevIns != nil && node.nextIns == nil {
		node.prevIns.nextIns = nil

	} else if node.prevIns == nil && node.nextIns != nil {
		node.nextIns.prevIns = nil
	}
}

func hash(key string, mod int) int {
	sum := 0
	for _, elem := range key {
		sum += int(elem)
	}
	return sum % mod
}

func main() {
	fin, _ := os.Open("linkedmap.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	mod := 10000
	table := make([]LinkedList, mod)
	var results []string

	var insert *Node

	for scanner.Scan() {
		txt := scanner.Text()
		switch {
		case strings.HasPrefix(txt, "put"):
			fields := strings.Fields(txt)
			key, value := fields[1], fields[2]
			node := table[hash(key, mod)].insert(key, value, insert)

			if insert != nil {
				insert.nextIns = node
			}

			insert = node

		case strings.HasPrefix(txt, "get"):
			key := strings.Fields(txt)[1]
			result := table[hash(key, mod)].get(key)

			if result == nil {
				results = append(results, "none")
			} else {
				results = append(results, result.value)
			}

		case strings.HasPrefix(txt, "delete"):
			key := strings.Fields(txt)[1]
			table[hash(key, mod)].delete(key)

		case strings.HasPrefix(txt, "prev"):
			key := strings.Fields(txt)[1]
			hashVal := hash(key, mod)
			result := table[hashVal].get(key)

			if result != nil {
				if result.prevIns != nil {
					results = append(results, result.prevIns.value)
				} else {
					results = append(results, "none")
				}
			} else {
				results = append(results, "none")
			}

		case strings.HasPrefix(txt, "next"):
			key := strings.Fields(txt)[1]
			hashVal := hash(key, mod)
			result := table[hashVal].get(key)

			if result != nil {
				if result.nextIns != nil {
					results = append(results, result.nextIns.value)
				} else {
					results = append(results, "none")
				}
			} else {
				results = append(results, "none")
			}

		}
	}
	fout, _ := os.Create("linkedmap.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
