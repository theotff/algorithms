package main

import (
	"bufio"
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

func (tree *BST) isBst(node *Node, min int, max int) bool {
	if node == nil {
		return true
	}

	if min <= node.value && node.value <= max {
		return tree.isBst(node.left, min, node.value-1) &&
			tree.isBst(node.right, node.value+1, max)
	} else {
		return false
	}
}

func main() {
	fin, _ := os.Open("check.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	array := make([][]int, n)

	state := true
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

	state = tree.isBst(tree.root, -1000000000, 1000000000)

	fout, _ := os.Create("check.out")
	if state {
		fout.WriteString("YES")
	} else {
		fout.WriteString("NO")
	}
	fout.Close()
}
