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

type AVL struct {
	root *Node
}

func intMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (tree *AVL) createNode(array [][]int, index int) (*Node, int) {
	if index != -1 {
		var lHeight, rHeight int

		node := &Node{value: array[index][0]}
		node.left, lHeight = tree.createNode(array, array[index][1]-1)
		node.right, rHeight = tree.createNode(array, array[index][2]-1)

		node.height = intMax(rHeight, lHeight) + 1
		return node, node.height

	} else {
		return nil, 0
	}
}

func (tree *AVL) setHeight(root *Node) int {
	if root != nil {
		height := intMax(tree.setHeight(root.right), tree.setHeight(root.left))

		root.height = height + 1
		return height + 1
	} else {
		return 0
	}
}

func (tree *AVL) rightRotation(root *Node) *Node {
	newRoot := root.left
	newRoot.right, root.left = root, newRoot.right
	tree.setHeight(newRoot)
	return newRoot
}

func (tree *AVL) leftRotation(root *Node) *Node {
	newRoot := root.right
	newRoot.left, root.right = root, newRoot.left
	tree.setHeight(newRoot)
	return newRoot
}

func (tree *AVL) rightLeftRotation(root *Node) *Node {
	root.right = tree.rightRotation(root.right)
	root = tree.leftRotation(root)
	return root
}

func (tree *AVL) leftRightRotation(root *Node) *Node {
	root.left = tree.leftRotation(root.left)
	root = tree.rightRotation(root)
	return root
}

func (tree *AVL) insert(root *Node, value int) (*Node, bool) {
	node := &Node{
		value:  value,
		height: 1,
		left:   nil,
		right:  nil}

	var state bool
	if tree.root != nil {
		if root.value > value {
			if root.left != nil {
				root.left, state = tree.insert(root.left, value)
				root.height = intMax(root.height, root.left.height+1)
				root = tree.balance(root)
				return root, state
			} else {
				root.left = node
				if root.right == nil {
					root.height += 1
				}
				return root, true
			}
		} else if root.value < value {
			if root.right != nil {
				root.right, state = tree.insert(root.right, value)
				root.height = intMax(root.height, root.right.height+1)
				root = tree.balance(root)
				return root, state
			} else {
				root.right = node
				if root.left == nil {
					root.height += 1
				}
				return root, true
			}
		} else {
			return root, false
		}
	} else {
		tree.root = node
		return node, true
	}
}

func (tree *AVL) getBalance(node *Node) int {
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

func (tree *AVL) balance(root *Node) *Node {
	balance := tree.getBalance(root)
	if balance > 1 {
		if tree.getBalance(root.right) != -1 {
			root = tree.leftRotation(root)
		} else {
			root = tree.rightLeftRotation(root)
		}
	} else if balance < -1 {
		if tree.getBalance(root.left) != 1 {
			root = tree.rightRotation(root)
		} else {
			root = tree.leftRightRotation(root)
		}
	}
	return root
}

func (tree *AVL) serialize(results [][]string, root *Node, index int) (int, int) {
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
	fin, _ := os.Open("addition.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	array := make([][]int, n)

	tree := &AVL{}

	for i := 0; i < n; i++ {
		scanner.Scan()
		elems := strings.Fields(scanner.Text())
		arr := make([]int, 3)
		for j := 0; j < 3; j++ {
			arr[j], _ = strconv.Atoi(elems[j])
		}
		array[i] = arr
	}

	scanner.Scan()
	newVal, _ := strconv.Atoi(scanner.Text())

	if len(array) != 0 {
		tree.root, _ = tree.createNode(array, 0)
	}

	var state bool
	tree.root, state = tree.insert(tree.root, newVal)

	results := make([][]string, n+1)
	for i := range results {
		results[i] = make([]string, 3)
	}

	tree.serialize(results, tree.root, 0)

	fout, _ := os.Create("addition.out")
	if state {
		fmt.Fprintln(fout, n+1)
		for i := 0; i < n+1; i++ {
			fmt.Fprintln(fout, strings.Join(results[i], " "))
		}
	} else {
		fmt.Fprintln(fout, n)
		for i := 0; i < n; i++ {
			fmt.Fprintln(fout, strings.Join(results[i], " "))
		}
	}
	fout.Close()
}
