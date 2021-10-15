package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	data     int
	position int
	next     *Node
}

type LinkedList struct {
	last *Node
}

func (queue *LinkedList) push(data int, pos int) {
	node := &Node{data: data, position: pos, next: queue.last}

	qnode := queue.last
	if qnode != nil {
		if qnode.next != nil {
			for qnode != nil {
				if data < qnode.data {
					queue.last = node
					break
				} else if qnode.data <= data && qnode.next == nil {
					node.next = nil
					qnode.next = node
					break
				} else if qnode.data <= data && qnode.next.data > data {
					node.next = qnode.next
					qnode.next = node
					break
				} else {
					qnode = qnode.next
				}
			}
		} else {
			if data < qnode.data {
				queue.last = node
			} else {
				node.next = nil
				queue.last.next = node
			}
		}
	} else {
		queue.last = node
	}
}

func (queue *LinkedList) replace(data int, pos int) {
	node := queue.last
	next := node.next

	for next != nil {
		if next.position == pos {
			node.next = next.next
			break
		} else {
			node = node.next
			next = node.next
		}
	}

	queue.push(data, pos)
}

func (queue *LinkedList) remove_min() (int, bool) {
	if queue.last != nil {
		value := queue.last.data
		queue.last = queue.last.next
		return value, true
	} else {
		return 0, false
	}
}

func main() {
	fin, _ := os.Open("priorityqueue.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	queue := &LinkedList{}
	pos := 1

	var results []string

	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "push") {
			var n int
			fmt.Sscanf(txt, "push %d", &n)
			queue.push(n, pos)
			pos += 1
		} else if strings.HasPrefix(txt, "decrease-key") {
			var n, m int
			fmt.Sscanf(txt, "decrease-key %d %d", &n, &m)
			queue.replace(m, n)
			pos += 1
		} else {
			result, state := queue.remove_min()
			if state {
				results = append(results, fmt.Sprint(result))
			} else {
				results = append(results, "*")
			}
			pos += 1
		}
	}
	fout, _ := os.Create("priorityqueue.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
