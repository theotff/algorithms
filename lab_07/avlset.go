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

func (tree *AVL) insert(root *Node, value int) *Node {
	node := &Node{
		value:  value,
		height: 1,
		left:   nil,
		right:  nil}

	if tree.root != nil {
		if root.value > value {
			if root.left != nil {
				root.left = tree.insert(root.left, value)
				root.height = intMax(root.height, root.left.height+1)
				root = tree.balance(root)
				return root
			} else {
				root.left = node
				if root.right == nil {
					root.height += 1
				}
				return root
			}
		} else if root.value < value {
			if root.right != nil {
				root.right = tree.insert(root.right, value)
				root.height = intMax(root.height, root.right.height+1)
				root = tree.balance(root)
				return root
			} else {
				root.right = node
				if root.left == nil {
					root.height += 1
				}
				return root
			}
		} else {
			return root
		}
	} else {
		tree.root = node
		return node
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

func (tree *AVL) exists(value int) bool {
	root := tree.root
	for root != nil {
		if root.value < value {
			root = root.right
		} else if root.value > value {
			root = root.left
		} else {
			return true
		}
	}
	return false
}

func main() {
	fin, _ := os.Open("avlset.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	fout, _ := os.Create("avlset.out")
	defer fout.Close()

	tree := &AVL{}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		num, _ := strconv.Atoi(fields[1])

		switch fields[0] {
		case "A":
			tree.root = tree.insert(tree.root, num)
			balance, _ := tree.getBalance(tree.root)
			fmt.Fprintln(fout, balance)
		case "C":
			exists := tree.exists(num)
			if exists {
				fmt.Fprintln(fout, "Y")
			} else {
				fmt.Fprintln(fout, "N")
			}
		case "D":
			var height int
			tree.root, _ = tree.delete(tree.root, num)
			if tree.root != nil {
				tree.root.height = height
			}
			balance, _ := tree.getBalance(tree.root)
			fmt.Fprintln(fout, balance)
		}
	}
}
