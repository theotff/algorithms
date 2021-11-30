package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value  int
	height int
	left   *Node
	right  *Node
}

type BST struct {
	root *Node
}

func (tree *BST) createNode(array [][]int, index int) (*Node, int) {
	if index != -1 {
		var lHeight, rHeight int

		node := &Node{value: array[index][0]}
		node.left, lHeight = tree.createNode(array, array[index][1]-1)
		node.right, rHeight = tree.createNode(array, array[index][2]-1)

		if lHeight > rHeight {
			node.height = lHeight + 1
			return node, lHeight + 1
		} else {
			node.height = rHeight + 1
			return node, rHeight + 1
		}

	} else {
		return nil, 0
	}
}

func (tree *BST) getBalance(value int) int {
	node := tree.get(value)
	var rHeight int
	var lHeight int

	if node.right != nil {
		rHeight = node.right.height
	} else {
		rHeight = 0
	}

	if node.left != nil {
		lHeight = node.left.height
	} else {
		lHeight = 0
	}
	return rHeight - lHeight
}

func (tree *BST) get(value int) *Node {
	node := tree.root
	for node != nil {
		if node.value == value {
			return node
		} else if node.value > value {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

func main() {
	fin, _ := os.Open("balance.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	array := make([][]int, n)

	tree := &BST{}

	for i := 0; i < n; i++ {
		scanner.Scan()
		elems := strings.Fields(scanner.Text())
		arr := make([]int, 3)
		for j := 0; j < 3; j++ {
			arr[j], _ = strconv.Atoi(elems[j])
		}
		array[i] = arr
	}

	if len(array) != 0 {
		tree.root, _ = tree.createNode(array, 0)
	}

	results := make([]string, n)

	for index, elem := range array {
		results[index] = fmt.Sprint(tree.getBalance(elem[0]))
	}

	fout, _ := os.Create("balance.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
