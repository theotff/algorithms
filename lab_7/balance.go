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

func intMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (tree *BST) createNode(array [][]int, index int, results []string) (*Node, int) {
	if index != -1 {
		var lHeight, rHeight int

		node := &Node{value: array[index][0]}
		node.left, lHeight = tree.createNode(array, array[index][1]-1, results)
		node.right, rHeight = tree.createNode(array, array[index][2]-1, results)

		node.height = intMax(rHeight, lHeight) + 1
		results[index] = fmt.Sprint(tree.getBalance(node))
		return node, node.height

	} else {
		return nil, 0
	}
}

func (tree *BST) getBalance(node *Node) int {
	rHeight := 0
	lHeight := 0

	if node.right != nil {
		rHeight = node.right.height
	}

	if node.left != nil {
		lHeight = node.left.height
	}

	return rHeight - lHeight
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

	results := make([]string, n)

	if len(array) != 0 {
		tree.root, _ = tree.createNode(array, 0, results)
	}

	fout, _ := os.Create("balance.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
