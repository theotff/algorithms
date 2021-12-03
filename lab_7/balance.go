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

func (tree *BST) createNode(array [][]int, index int, results []string) (*Node, int) {
	if index != -1 {
		var lHeight, rHeight int

		node := &Node{value: array[index][0]}
		node.left, lHeight = tree.createNode(array, array[index][1]-1, results)
		node.right, rHeight = tree.createNode(array, array[index][2]-1, results)

		if lHeight > rHeight {
			node.height = lHeight + 1
			results[index] = fmt.Sprint(tree.getBalance(node))
			return node, lHeight + 1
		} else {
			node.height = rHeight + 1
			results[index] = fmt.Sprint(tree.getBalance(node))
			return node, rHeight + 1
		}

	} else {
		return nil, 0
	}
}

func (tree *BST) getBalance(node *Node) int {
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
