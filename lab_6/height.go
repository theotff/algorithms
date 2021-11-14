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
	depth int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (tree *BST) createNode(array [][]int, index int) *Node {
	if index != -1 {
		node := &Node{value: array[index][0], depth: 1}
		node.left = tree.createNode(array, array[index][1]-1)
		node.right = tree.createNode(array, array[index][2]-1)
		return node

	} else {
		return nil
	}
}

func (tree *BST) height() int {
	var stack [200000]*Node
	index := -1
	root := tree.root
	height := 0
	curDepth := 1

	for {
		if root == nil {
			if index == -1 {
				break
			} else {
				root = stack[index]
				index -= 1
				root.left = nil
				curDepth = root.depth
			}

		} else {

			if curDepth > height {
				height = curDepth
			}

			if root.left != nil && root.right != nil {
				index += 1
				stack[index] = root
				root = root.left
				root.depth += curDepth
				curDepth += 1

			} else {
				if root.left != nil {
					root = root.left
					root.depth += curDepth
					curDepth += 1
				} else if root.right != nil {
					root = root.right
					root.depth += curDepth
					curDepth += 1
				} else {
					root = nil
				}
			}
		}
	}

	return height
}

func main() {
	fin, _ := os.Open("height.in")
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
		tree.root = &Node{value: array[0][0], depth: 1}
		tree.root.left = tree.createNode(array, array[0][1]-1)
		tree.root.right = tree.createNode(array, array[0][2]-1)
	}

	fout, _ := os.Create("height.out")
	fout.WriteString(fmt.Sprint(tree.height()))
	fout.Close()
}
