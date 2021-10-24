package main

import (
	"bufio"
	"os"
	"strings"
)

type Node struct {
	key      string
	value    string
	next     *Node
	prev_ins *Node
	next_ins *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) insert(key string, value string, prev *Node) *Node {
	result := list.get(key)
	if result == nil {
		node := &Node{
			key:      key,
			value:    value,
			next:     list.last,
			prev_ins: prev}

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
		if node.next != nil {
			if node.next.key == key {
				if node.next.next != nil {
					swap_links(node)
					node.next = node.next.next
				} else {
					swap_links(node)
					node.next = nil
				}
				return
			} else {
				node = node.next
			}
		} else {
			if node.key == key {
				swap_links(node)
				list.last = nil
			}
			return
		}
	}
}

func swap_links(node *Node) {
	if node.prev_ins != nil && node.next_ins != nil {
		node.prev_ins.next_ins, node.next_ins.prev_ins = node.next_ins, node.prev_ins
	} else if node.prev_ins != nil && node.next_ins == nil {
		node.prev_ins.next_ins = nil
	} else if node.prev_ins == nil && node.next_ins != nil {
		node.next_ins.prev_ins = nil
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
				insert.next_ins = node
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
			hash_val := hash(key, mod)
			result := table[hash_val].get(key)

			if result != nil {
				if result.prev_ins != nil {
					results = append(results, result.prev_ins.value)
				} else {
					results = append(results, "none")
				}
			} else {
				results = append(results, "none")
			}

		case strings.HasPrefix(txt, "next"):
			key := strings.Fields(txt)[1]
			hash_val := hash(key, mod)
			result := table[hash_val].get(key)

			if result != nil {
				if result.next_ins != nil {
					results = append(results, result.next_ins.value)
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
