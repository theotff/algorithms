package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type Queue struct {
	first *Node
	last  *Node
}

func (q *Queue) push(data int) {
	if q.last == nil {
		node := &Node{data: data}
		q.first = node
		q.last = node

	} else {
		node := &Node{data: data}
		q.first.next = node
		q.first = node
	}
}

func (q *Queue) pop() int {
	value := q.last.data
	q.last = q.last.next
	return value
}

func main() {
	var n int
	fin, _ := os.Open("queue.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)

	var results []int
	q := &Queue{}

	for scanner.Scan() {
		operation := scanner.Text()

		if len(operation) == 1 {
			results = append(results, q.pop())

		} else {
			var num int
			fmt.Sscanf(operation, "+ %d", &num)
			q.push(num)
		}
	}

	resultArr := make([]string, len(results))

	for i := 0; i < len(results); i++ {
		resultArr[i] = fmt.Sprint(results[i])
	}

	fout, _ := os.Create("queue.out")
	fout.WriteString(strings.Join(resultArr, "\n"))
	fout.Close()
}
