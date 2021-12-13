package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	fin, _ := os.Open("avlset.in")
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
