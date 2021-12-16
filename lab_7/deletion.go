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

func (tree *AVL) balance(root *Node) *Node {
	if root != nil {
		var balance int
		balance, root.height = tree.getBalance(root)

		if balance > 1 {
			rBalance, _ := tree.getBalance(root.right)
			if rBalance != -1 {
				root = tree.leftRotation(root)
			} else {
				root = tree.rightLeftRotation(root)
			}

		} else if balance < -1 {
			lBalance, _ := tree.getBalance(root.left)
			if lBalance != 1 {
				root = tree.rightRotation(root)
			} else {
				root = tree.leftRightRotation(root)
			}
		}
		return root
	} else {
		return nil
	}

}

func (tree *AVL) getBalance(node *Node) (int, int) {
	rHeight := 0
	lHeight := 0

	if node != nil {
		if node.right != nil {
			rHeight = node.right.height
		}
		if node.left != nil {
			lHeight = node.left.height
		}
		return rHeight - lHeight, intMax(rHeight, lHeight) + 1
	} else {
		return 0, 0
	}
}

func (tree *AVL) delete(root *Node, value int) (*Node, bool) {
	if root != nil {
		var state bool
		if root.value > value {
			root.left, state = tree.delete(root.left, value)
			root = tree.balance(root)
			return root, state

		} else if root.value < value {
			root.right, state = tree.delete(root.right, value)
			root = tree.balance(root)
			return root, state

		} else {
			if root.left == nil && root.right == nil {
				return nil, true
			} else if root.left == nil && root.right != nil {
				root = root.right
				return root, true
			} else {
				rightmost := tree.findRight(root.left)
				root.value = rightmost.value
				root.left, _ = tree.delete(root.left, rightmost.value)
				root = tree.balance(root)
				return root, true
			}
		}
	} else {
		return nil, false
	}
}

func (tree *AVL) findRight(root *Node) *Node {
	result := root
	for root != nil {
		result = root
		root = root.right
	}
	return result
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
	fin, _ := os.Open("deletion.in")
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
	tree.root, state = tree.delete(tree.root, newVal)

	results := make([][]string, n)
	for i := range results {
		results[i] = make([]string, 3)
	}

	tree.serialize(results, tree.root, 0)

	fout, _ := os.Create("deletion.out")
	var foutSize int
	if state {
		foutSize = n - 1
	} else {
		foutSize = n
	}

	fmt.Fprintln(fout, foutSize)
	for i := 0; i < foutSize; i++ {
		fmt.Fprintln(fout, strings.Join(results[i], " "))
	}
	fout.Close()
}
