package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ListNode struct {
	key  string
	last *Node
	next *ListNode
}

type Node struct {
	key  string
	next *Node
}

type LinkedList struct {
	last *ListNode
}

func (list *LinkedList) insert(key string, value string) {
	qnode := list.last

	for qnode != nil {
		if qnode.key == key {

			if qnode.last != nil {
				elem := &Node{key: value, next: qnode.last}
				qnode.last = elem
			} else {
				elem := &Node{key: value, next: nil}
				qnode.last = elem
			}
			return

		} else {
			qnode = qnode.next
		}
	}

	elem := &Node{key: value, next: nil}
	node := &ListNode{key: key, last: elem, next: list.last}
	list.last = node
}

func (list *LinkedList) get(key string) *ListNode {
	listNode := list.last

	for listNode != nil {
		if listNode.key == key {
			return listNode
		} else {
			listNode = listNode.next
		}
	}

	return nil
}

func (list *LinkedList) delete(key string, value string) {
	listNode := list.get(key)

	if listNode != nil {
		node := listNode.last
		for node != nil {
			if node.next != nil && node.next.key == value {
				if node.next.next != nil {
					node.next = node.next.next
				} else {
					node.next = nil
				}
				return

			} else if node.next != nil && node.next.key != value {
				node = node.next

			} else {
				if node.key == value {
					list.last = nil
				}
				return
			}
		}
	}
}

func (list *LinkedList) deleteAll(key string) {
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
	for _, elem := range key {
		sum += int(elem)
	}
	return sum % mod
}

func main() {
	fin, _ := os.Open("multimap.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	mod := 100
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
			node := table[hash(key, mod)].get(key)
			if node != nil {
				count := 0
				var res []string
				elem := node.last
				for elem != nil {
					count += 1
					res = append(res, elem.key)
					elem = elem.next
				}
				results = append(results, fmt.Sprint(count)+" "+strings.Join(res, " "))
			} else {
				results = append(results, "0")
			}

		case strings.HasPrefix(txt, "delete "):
			fields := strings.Fields(txt)
			key, value := fields[1], fields[2]
			table[hash(key, mod)].delete(key, value)

		case strings.HasPrefix(txt, "deleteall"):
			key := strings.Fields(txt)[1]
			table[hash(key, mod)].deleteAll(key)
		}
	}
	fout, _ := os.Create("multimap.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
