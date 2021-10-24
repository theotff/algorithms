package main

import (
	"bufio"
	"os"
	"strings"
)

type Node struct {
	key   string
	value string
	next  *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) insert(key string, value string) {
	if list.get(key) == "none" {
		node := &Node{key: key, value: value, next: list.last}
		list.last = node
	}
}

func (list *LinkedList) get(key string) string {
	node := list.last
	for node != nil {
		if node.key == key {
			return node.value
		} else {
			node = node.next
		}
	}
	return "none"
}

func (list *LinkedList) delete(key string) {
	node := list.last
	for node != nil {
		if node.next != nil {
			if node.next.key == key {
				if node.next.next != nil {
					node.next = node.next.next
				} else {
					node.next = nil
				}
				return
			} else {
				node = node.next
			}
		} else {
			if node.key == key {
				list.last = nil
			}
			return
		}
	}
}

func hash(key string, mod int) int {
	sum := 0
	for _, elem := range strings.Split(key, "") {
		sum += int([]rune(elem)[0])
	}
	return sum % mod
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
		switch {
		case strings.HasPrefix(txt, "put"):
			fields := strings.Fields(txt)
			key, value := fields[1], fields[2]
			table[hash(key, mod)].insert(key, value)

		case strings.HasPrefix(txt, "get"):
			key := strings.Fields(txt)[1]
			results = append(results, table[hash(key, mod)].get(key))

		case strings.HasPrefix(txt, "delete"):
			key := strings.Fields(txt)[1]
			table[hash(key, mod)].delete(key)
		}
	}
	fout, _ := os.Create("map.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
