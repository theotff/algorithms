package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	last *Node
}

func (list *LinkedList) insert(data int) {
	node := &Node{data: data, next: list.last}
	list.last = node
}

func (list *LinkedList) remove() int {
	value := list.last.data
	list.last = list.last.next
	return value
}

func (list *LinkedList) add() {
	val1 := list.remove()
	val2 := list.remove()
	list.insert(val2 + val1)
}

func (list *LinkedList) subtract() {
	val1 := list.remove()
	val2 := list.remove()
	list.insert(val2 - val1)
}

func (list *LinkedList) multiply() {
	val1 := list.remove()
	val2 := list.remove()
	list.insert(val2 * val1)
}

func main() {
	data_raw, _ := ioutil.ReadFile("postfix.in")
	data := strings.Split(string(data_raw), "")

	stack := &LinkedList{}

	for _, elem := range data {
		if num, err := strconv.Atoi(elem); err == nil {
			stack.insert(num)
		} else {
			switch elem {
			case "-":
				stack.subtract()
			case "+":
				stack.add()
			case "*":
				stack.multiply()
			}
		}
	}
	fout, _ := os.Create("postfix.out")
	fout.WriteString(fmt.Sprint(stack.last.data))
	fout.Close()
}
