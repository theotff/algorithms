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
			if value < root.value {
				if root.left != nil {
					root = root.left
				} else {
					node := &Node{value: value, parent: root, left: nil, right: nil}
					root.left = node
				}
			} else if value > root.value {
				if root.right != nil {
					root = root.right
				} else {
					node := &Node{value: value, parent: root, left: nil, right: nil}
					root.right = node
				}
			} else {
				return
			}
		}
	} else {
		node := &Node{value: value, parent: nil, left: nil, right: nil}
		tree.root = node
	}
}

func (tree *BST) search(node *Node, value int) *Node {
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

func (tree *BST) exists(value int) bool {
	node := tree.search(tree.root, value)
	return node != nil
}

func (tree *BST) delete(value int) {
	node := tree.search(tree.root, value)
	if node != nil {
		if node.right == nil && node.left == nil {
			if node.parent != nil {
				if node.parent.left == node {
					node.parent.left = nil
				} else {
					node.parent.right = nil
				}
			} else {
				tree.root = nil
			}
		} else if node.right != nil && node.left != nil {
			successor := tree.next(node.value)

			if successor.left != nil {
				fmt.Println(tree.root.parent.value)
			}

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
			if node.right != nil {
				if node.parent == nil {
					tree.root = node.right
					node.right.parent = nil
				} else if node.parent.left == node {
					node.parent.left = node.right
					node.parent.left.parent = node.parent
				} else {
					node.parent.right = node.right
					node.parent.right.parent = node.parent
				}
			} else {
				if node.parent == nil {
					tree.root = node.left
					node.left.parent = nil
				} else if node.parent.left == node {
					node.parent.left = node.left
					node.parent.left.parent = node.parent
				} else {
					node.parent.right = node.left
					node.parent.right.parent = node.parent
				}
			}
		}
	}
}

func (tree *BST) prev(value int) *Node {
	var result *Node = nil
	root := tree.root
	for root != nil {
		if root.value < value {
			result = root
			root = root.right
		} else {
			root = root.left
		}
	}

	return result
}

func (tree *BST) next(value int) *Node {
	var result *Node = nil
	root := tree.root
	for root != nil {
		if root.value > value {
			result = root
			root = root.left
		} else {
			root = root.right
		}
	}
	return result
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
			next := tree.next(value)
			if next != nil {
				results = append(results, fmt.Sprint(next.value))
			} else {
				results = append(results, "none")
			}

		case strings.HasPrefix(txt, "prev"):
			var value int
			fmt.Sscanf(txt, "prev %d", &value)
			prev := tree.prev(value)
			if prev != nil {
				results = append(results, fmt.Sprint(prev.value))
			} else {
				results = append(results, "none")
			}
		}
	}
	fout, _ := os.Create("bstsimple.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
