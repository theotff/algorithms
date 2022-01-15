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

func (tree *BST) height(root *Node) int {
	if root != nil {
		height := intMax(tree.height(root.right), tree.height(root.left))

		root.height = height + 1
		return height + 1
	} else {
		return 0
	}
}

func (tree *BST) rightRotation(root *Node) *Node {
	newRoot := root.left
	newRoot.right, root.left = root, newRoot.right
	newRoot.height = tree.height(newRoot)
	newRoot.right.height = tree.height(newRoot.right)
	return newRoot
}

func (tree *BST) leftRotation(root *Node) *Node {
	newRoot := root.right
	newRoot.left, root.right = root, newRoot.left
	newRoot.height = tree.height(newRoot)
	newRoot.left.height = tree.height(newRoot.left)
	return newRoot
}

func (tree *BST) rightLeftRotation(root *Node) *Node {
	root.right = tree.rightRotation(root.right)
	root = tree.leftRotation(root)
	return root
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

func (tree *BST) balance() {
	if tree.getBalance(tree.root.right) != -1 {
		tree.root = tree.leftRotation(tree.root)
	} else {
		tree.root = tree.rightLeftRotation(tree.root)
	}
}

func (tree *BST) serialize(results [][]string, root *Node, index int) (int, int) {
	if root != nil {
		var position int
		curIndex := index
		results[curIndex][0] = fmt.Sprint(root.value)

		for i := 1; i < 3; i++ {
			node := root.right
			if i == 1 {
				node = root.left
			}
			index, position = tree.serialize(results, node, index+1)
			if position != 0 {
				position += 1
			}
			results[curIndex][i] = fmt.Sprint(position)
		}
		return index, curIndex
	} else {
		return index - 1, 0
	}
}

func main() {
	fin, _ := os.Open("rotation.in")
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

	tree.balance()

	results := make([][]string, n)
	for i := range results {
		results[i] = make([]string, 3)
	}

	tree.serialize(results, tree.root, 0)

	output := make([]string, n)
	for i := 0; i < n; i++ {
		output[i] = strings.Join(results[i], " ")
	}

	fout, _ := os.Create("rotation.out")
	fout.WriteString(fmt.Sprint(n) + "\n" + strings.Join(output, "\n"))
	fout.Close()
}
