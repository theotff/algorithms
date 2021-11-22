package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const mod = 1000

type Node struct {
	key   string
	value string
	table []LinkedList
	prev  *Node
	next  *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) put(key string, value string) {
	result := list.get(key)
	if result == nil {
		node := &Node{
			key:  key,
			next: list.last}

		if list.last != nil {
			list.last.prev = node
		}
		list.last = node
		node.table = make([]LinkedList, mod)

		result = node
	}
	hashSum := hash(value, mod)
	result.table[hashSum].listPut(value)
}

func (list *LinkedList) listPut(value string) {
	result := list.listGet(value)
	if result == nil {
		node := &Node{
			value: value,
			next:  list.last}

		if list.last != nil {
			list.last.prev = node
		}
		list.last = node
	}
}

func (list *LinkedList) listGet(value string) *Node {
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

func (list *LinkedList) delete(key string, value string) {
	result := list.get(key)
	if result != nil {
		hashSum := hash(value, mod)
		node := result.table[hashSum].listGet(value)
		if node != nil {
			if node.next != nil {
				node.next.prev = node.prev
			}
			if node.prev != nil {
				node.prev.next = node.next
			} else {
				result.table[hashSum].last = node.next
			}
		}
	}
}

func (list *LinkedList) deleteAll(key string) {
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

func hash(key string, mod int) int {
	hashSum := 5381
	for _, elem := range key {
		hashSum = ((hashSum << 5) + hashSum) + int(elem)
	}
	return int(math.Abs(float64(hashSum % mod)))
}

func main() {
	fin, _ := os.Open("multimap.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

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
			table[hashSum].put(key, value)

		case "get":
			node := table[hashSum].get(key)
			if node != nil {
				count := 0
				var res []string
				for i := 0; i < mod; i++ {
					elem := node.table[i].last
					for elem != nil {
						count += 1
						res = append(res, elem.value)
						elem = elem.next
					}
				}
				results = append(results, fmt.Sprint(count)+" "+strings.Join(res, " "))
			} else {
				results = append(results, "0")
			}

		case "delete":
			value := fields[2]
			table[hashSum].delete(key, value)

		case "deleteall":
			table[hashSum].deleteAll(key)
		}
	}
	fout, _ := os.Create("multimap.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
