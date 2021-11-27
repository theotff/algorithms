package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (tree *BST) createNode(array [][]int, index int) *Node {
	if index != -1 {
		node := &Node{value: array[index][0]}
		node.left = tree.createNode(array, array[index][1]-1)
		node.right = tree.createNode(array, array[index][2]-1)
		return node
	} else {
		return nil
	}
}

func (tree *BST) height(root *Node) int {
	if root != nil {
		var height int
		rHeight := tree.height(root.right)
		lHeight := tree.height(root.left)

		if rHeight > lHeight {
			height = rHeight
		} else {
			height = lHeight
		}

		return height + 1
	} else {
		return 0
	}
}

func (tree *BST) balance(value int) int {
	node := tree.get(value)
	balance := tree.height(node.right) - tree.height(node.left)
	return balance
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
		tree.root = &Node{value: array[0][0]}
		tree.root.left = tree.createNode(array, array[0][1]-1)
		tree.root.right = tree.createNode(array, array[0][2]-1)
	}

	results := make([]string, n)

	for index, elem := range array {
		results[index] = fmt.Sprint(tree.balance(elem[0]))
	}

	fout, _ := os.Create("balance.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
