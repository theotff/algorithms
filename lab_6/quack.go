package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	key   string
	value int
	next  *Node
}

type LinkedList struct {
	last  *Node
	first *Node
}

func (queue *LinkedList) queue_put(value int) {
	if queue.last == nil {
		node := &Node{value: value, next: nil}
		queue.first = node
		queue.last = node
	} else {
		node := &Node{value: value, next: nil}
		queue.first.next = node
		queue.first = node
	}
}

func (queue *LinkedList) queue_get() int {
	value := queue.last.value
	queue.last = queue.last.next
	return value
}

func (list *LinkedList) hash_put(key string, index int) {
	if list.hash_get(key) == nil {
		node := &Node{key: key, value: index, next: list.last}
		list.last = node
	}
}

func (list *LinkedList) hash_get(key string) *Node {
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
	sum := 0
	for _, elem := range key {
		sum += int(elem)
	}
	return sum % mod
}

func main() {
	data, _ := ioutil.ReadFile("quack.in")
	callstack := strings.Fields(string(data))

	intsize := 65536
	queue := &LinkedList{}
	mod := 1000
	table := make([]LinkedList, mod)
	registers := make([]int, 26)
	index := 0
	ln := len(callstack)
	var results []string

	find_func_index := func(label string) (label_index int) {
		for index, elem := range callstack {
			if strings.HasPrefix(elem, ":") && len(elem) > 1 {
				if elem[1:] == label {
					table[hash(label, mod)].hash_put(label, index)
					return index - 1
				}
			}
		}
		return
	}

	for index < ln {
		op := callstack[index]
		switch {
		case op == "+":
			value := queue.queue_get() + queue.queue_get()
			queue.queue_put(value & intsize)
			index += 1

		case op == "-":
			value := queue.queue_get() - queue.queue_get()
			queue.queue_put(value & intsize)
			index += 1

		case op == "*":
			value := queue.queue_get() * queue.queue_get()
			queue.queue_put(value & intsize)
			index += 1

		case op == "/":
			value := queue.queue_get() / queue.queue_get()
			queue.queue_put(value & intsize)
			index += 1

		case op == "%":
			value := queue.queue_get() % queue.queue_get()
			queue.queue_put(value & intsize)
			index += 1

		case op == "P":
			results = append(results, fmt.Sprint(queue.queue_get()))
			index += 1

		case op == "C":
			results = append(results, string(rune(queue.queue_get()%256)))
			index += 1

		case op == "Q":
			index = ln

		case strings.HasPrefix(op, "<"):
			reg_index := int([]rune(strings.ToLower(op[1:]))[0]) - 97
			queue.queue_put(registers[reg_index])
			index += 1

		case strings.HasPrefix(op, ">"):
			reg_index := int([]rune(strings.ToLower(op[1:]))[0]) - 97
			registers[reg_index] = queue.queue_get()
			index += 1

		case strings.HasPrefix(op, "P"):
			reg_index := int([]rune(strings.ToLower(op[1:]))[0]) - 97
			results = append(results, fmt.Sprint(registers[reg_index]))
			index += 1

		case strings.HasPrefix(op, "C"):
			reg_index := int([]rune(strings.ToLower(op[1:]))[0]) - 97
			results = append(results, string(rune(registers[reg_index])%256))
			index += 1

		case strings.HasPrefix(op, ":"):
			key := op[1:]
			table[hash(key, mod)].hash_put(key, index)
			index += 1

		case strings.HasPrefix(op, "J"):
			key := op[1:]
			label_node := table[hash(key, mod)].hash_get(key)
			var label_index int

			if label_node != nil {
				label_index = label_node.value
			} else {
				label_index = find_func_index(key)
			}
			index = label_index + 1

		case strings.HasPrefix(op, "Z"):
			reg_index := int([]rune(strings.ToLower(op[1:2]))[0]) - 97

			if registers[reg_index] == 0 {
				label := op[2:]
				label_node := table[hash(label, mod)].hash_get(label)
				var label_index int

				if label_node != nil {
					label_index = label_node.value
				} else {
					label_index = find_func_index(label)
				}
				index = label_index + 1
			}
			index += 1

		case strings.HasPrefix(op, "E"):
			reg_index_1 := int([]rune(strings.ToLower(op[1:2]))[0]) - 97
			reg_index_2 := int([]rune(strings.ToLower(op[2:3]))[0]) - 97

			if registers[reg_index_1] == registers[reg_index_2] {
				label := op[3:]
				label_node := table[hash(label, mod)].hash_get(label)
				var label_index int

				if label_node != nil {
					label_index = label_node.value
				} else {
					label_index = find_func_index(label)
				}
				index = label_index + 1
			}
			index += 1

		case strings.HasPrefix(op, "G"):
			reg_index_1 := int([]rune(strings.ToLower(op[1:2]))[0]) - 97
			reg_index_2 := int([]rune(strings.ToLower(op[2:3]))[0]) - 97

			if registers[reg_index_1] > registers[reg_index_2] {
				label := op[3:]
				label_node := table[hash(label, mod)].hash_get(label)
				var label_index int

				if label_node != nil {
					label_index = label_node.value
				} else {
					label_index = find_func_index(label)
				}
				index = label_index + 1
			}
			index += 1

		default:
			num, _ := strconv.Atoi(op)
			queue.queue_put(num)
			index += 1
		}
	}
	fout, _ := os.Create("quack.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
