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
	left   *Node
	right  *Node
	parent *Node
}

type BST struct {
	root *Node
}

func (tree *BST) insert(value int) {
	if tree.root != nil {
		root := tree.root
		for root != nil {
			node := &Node{value: value, parent: root}
			if value < root.value {
				if root.left != nil {
					root = root.left
				} else {
					root.left = node
				}
			} else if value > root.value {
				if root.right != nil {
					root = root.right
				} else {
					root.right = node
				}
			} else {
				return
			}
		}
	} else {
		node := &Node{value: value, parent: nil}
		tree.root = node
	}
}

func (tree *BST) search(node *Node, value int) *Node {
	if node == nil || node.value == value {
		return node
	} else if node.value > value {
		return tree.search(node.left, value)
	} else {
		return tree.search(node.right, value)
	}
}

func (tree *BST) exists(value int) bool {
	node := tree.search(tree.root, value)
	return node != nil
}

func (tree *BST) min(root *Node) *Node {
	if root.left == nil {
		return root
	} else {
		return tree.min(root.left)
	}
}

func (tree *BST) next_node(node *Node) *Node {
	if node.right != nil {
		return tree.min(node.left)
	}
	parent := node.parent
	for parent != nil && node == parent.right {
		node = parent
		parent = parent.parent
	}
	return parent
}

func (tree *BST) delete(value int) {
	node := tree.search(tree.root, value)
	if node != nil {
		if node.right == nil && node.left == nil {
			if node.parent.left == node {
				node.parent.left = nil
			} else {
				node.parent.right = node
			}
		} else if node.right != nil && node.left != nil {
			successor := tree.next_node(node)
			node.value = successor.value
			if successor.parent.left == successor {
				successor.parent.left = successor.right
				if successor.right != nil {
					successor.right.parent = successor.parent
				}
			} else {
				successor.parent.right = successor.right
				if successor.right != nil {
					successor.right.parent = successor.parent
				}
			}

		} else {
			if node.left == nil {
				if node.parent.left == node {
					node.parent.left = node.right
				} else {
					node.parent.right = node.right
				}
			} else {
				if node.parent.left == node {
					node.parent.left = node.left
				} else {
					node.parent.right = node.left
				}
			}
		}
	}
}

func (tree *BST) prev(value int) string {
	result := "none"
	root := tree.root
	for root != nil {
		if root.value < value {
			if root.right != nil {
				root = root.right
			} else {
				result = fmt.Sprint(root.value)
				break
			}
		} else if root.value > value {
			root = root.left
		}
	}
	return result
}

func (tree *BST) next(value int) string {
	result := 0
	status := false
	root := tree.root
	for root != nil {
		if root.value > value {
			result, status = root.value, true
			root = root.left
		} else if root.value < value {
			root = root.right
		}
	}
	if status {
		return fmt.Sprint(result)
	} else {
		return "none"
	}
}

func main() {
	fin, _ := os.Open("bstsimple.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	tree := &BST{}
	var results []string

	for scanner.Scan() {
		txt := scanner.Text()

		switch {
		case strings.HasPrefix(txt, "insert"):
			var value int
			fmt.Sscanf(txt, "insert %d", &value)
			tree.insert(value)

		case strings.HasPrefix(txt, "exists"):
			var value int
			fmt.Sscanf(txt, "exists %d", &value)
			results = append(results, strconv.FormatBool(tree.exists(value)))

		case strings.HasPrefix(txt, "delete"):
			var value int
			fmt.Sscanf(txt, "delete %d", &value)
			tree.delete(value)

		case strings.HasPrefix(txt, "next"):
			var value int
			fmt.Sscanf(txt, "next %d", &value)
			results = append(results, tree.next(value))

		case strings.HasPrefix(txt, "prev"):
			var value int
			fmt.Sscanf(txt, "prev %d", &value)
			results = append(results, tree.prev(value))
		}
	}
	fout, _ := os.Create("bstsimple.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
