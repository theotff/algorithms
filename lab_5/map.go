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
	prev  *Node
}

type Dll struct {
	last *Node
}

func (list *Dll) put(key string, value string) {
	if list.get(key) == "none" {
		node := &Node{key: key, value: value, next: list.last, prev: nil}
		if list.last != nil {
			list.last.prev = node
		}
		list.last = node
	}
}

func (list *Dll) get(key string) string {
	if list != nil {
		node := list.last
		for node != nil {
			if node.key == key {
				return node.value
			} else {
				node = node.next
			}
		}
	}
	return "none"
}

func (list *Dll) delete(key string) {
	node := list.last
	for node != nil {
		if node.key == key {
			if node.prev != nil {
				if node.next != nil {
					node.prev.next, node.next.prev = node.next.prev, node.prev.next
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
	mod := 100
	table := make([]Dll, mod)
	var results []string

	for scanner.Scan() {
		txt := scanner.Text()
		switch {
		case strings.HasPrefix(txt, "put"):
			split := strings.Split(txt[4:], " ")
			key, value := split[0], split[1]
			table[hash(key, mod)].put(key, value)
		case strings.HasPrefix(txt, "get"):
			key := txt[4:]
			results = append(results, table[hash(key, mod)].get(key))
		case strings.HasPrefix(txt, "delete"):
			key := txt[7:]
			table[hash(key, mod)].delete(key)
		}
	}
	fout, _ := os.Create("map.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
