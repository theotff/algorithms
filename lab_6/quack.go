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

func (queue *LinkedList) queuePut(value int) {
	node := &Node{value: value, next: nil}
	if queue.last == nil {
		queue.last = node
	} else {
		queue.first.next = node
	}
	queue.first = node
}

func (queue *LinkedList) queueGet() int {
	value := queue.last.value
	queue.last = queue.last.next
	return value
}

func (list *LinkedList) hashPut(key string, index int) {
	if list.hashGet(key) == nil {
		node := &Node{key: key, value: index, next: list.last}
		list.last = node
	}
}

func (list *LinkedList) hashGet(key string) *Node {
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

func asciiIndex(char string) int {
	return int([]rune(strings.ToLower(char))[0]) - 97
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

	findFuncIndex := func(label string) (labelIndex int) {
		for index, elem := range callstack {
			if strings.HasPrefix(elem, ":") && len(elem) > 1 {
				if elem[1:] == label {
					table[hash(label, mod)].hashPut(label, index)
					return index
				}
			}
		}
		return
	}

	for index < ln {
		op := callstack[index]

		switch string(op[0]) {
		case "+":
			value := queue.queueGet() + queue.queueGet()
			queue.queuePut(value % intsize)
			index += 1

		case "-":
			value := queue.queueGet() - queue.queueGet()

			if value < 0 {
				value = intsize + value
			}

			queue.queuePut(value % intsize)
			index += 1

		case "*":
			value := queue.queueGet() * queue.queueGet()
			queue.queuePut(value % intsize)
			index += 1

		case "/":
			x := queue.queueGet()
			y := queue.queueGet()
			var value int

			if y != 0 {
				value = x / y
			} else {
				value = 0
			}

			queue.queuePut(value % intsize)
			index += 1

		case "%":
			x := queue.queueGet()
			y := queue.queueGet()
			var value int

			if y != 0 {
				value = x % y
			} else {
				value = 0
			}

			queue.queuePut(value % intsize)
			index += 1

		case "P":
			if len(op) == 1 {
				results = append(results, fmt.Sprint(queue.queueGet()), "\n")
				index += 1
			} else {
				regIndex := int([]rune(strings.ToLower(op[1:]))[0]) - 97
				results = append(results, fmt.Sprint(registers[regIndex]), "\n")
				index += 1
			}

		case "C":
			if len(op) == 1 {
				results = append(results, string(rune(queue.queueGet()%256)))
				index += 1
			} else {
				regIndex := asciiIndex(op[1:])
				results = append(results, string(rune(registers[regIndex])%256))
				index += 1
			}

		case "Q":
			index = ln

		case "<":
			regIndex := asciiIndex(op[1:])
			queue.queuePut(registers[regIndex])
			index += 1

		case ">":
			regIndex := asciiIndex(op[1:])
			registers[regIndex] = queue.queueGet()
			index += 1

		case ":":
			key := op[1:]
			table[hash(key, mod)].hashPut(key, index)
			index += 1

		case "J":
			key := op[1:]
			labelNode := table[hash(key, mod)].hashGet(key)
			var labelIndex int

			if labelNode != nil {
				labelIndex = labelNode.value
			} else {
				labelIndex = findFuncIndex(key)
			}
			index = labelIndex + 1

		case "Z":
			regIndex := asciiIndex(op[1:2])

			if registers[regIndex] == 0 {
				label := op[2:]
				labelNode := table[hash(label, mod)].hashGet(label)
				var labelIndex int

				if labelNode != nil {
					labelIndex = labelNode.value
				} else {
					labelIndex = findFuncIndex(label)
				}
				index = labelIndex
			}
			index += 1

		case "E":
			regIndex1 := asciiIndex(op[1:2])
			regIndex2 := asciiIndex(op[2:3])

			if registers[regIndex1] == registers[regIndex2] {
				label := op[3:]
				labelNode := table[hash(label, mod)].hashGet(label)
				var labelIndex int

				if labelNode != nil {
					labelIndex = labelNode.value
				} else {
					labelIndex = findFuncIndex(label)
				}
				index = labelIndex
			}
			index += 1

		case "G":
			regIndex1 := asciiIndex(op[1:2])
			regIndex2 := asciiIndex(op[2:3])

			if registers[regIndex1] > registers[regIndex2] {
				label := op[3:]
				labelNode := table[hash(label, mod)].hashGet(label)
				var labelIndex int

				if labelNode != nil {
					labelIndex = labelNode.value
				} else {
					labelIndex = findFuncIndex(label)
				}
				index = labelIndex
			}
			index += 1

		default:
			num, err := strconv.Atoi(op)
			if err == nil {
				queue.queuePut(num % intsize)
			}
			index += 1
		}
	}

	fout, _ := os.Create("quack.out")
	fout.WriteString(strings.Join(results, ""))
	fout.Close()
}
