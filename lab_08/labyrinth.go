package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	route string
	state bool
	end   bool
	x     int
	y     int
}

type Graph struct {
	matrix [][]*Cell
	width  int
	height int
	start  []int
}

func (g *Graph) construct(w, h int) {
	g.width = w
	g.height = h
	g.matrix = make([][]*Cell, h)
}

type Node struct {
	cell *Cell
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) push(cell *Cell) {
	node := &Node{cell: cell}
	if q.head == nil {
		q.tail = node
		q.head = q.tail
	} else {
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue) pop() *Cell {
	cell := q.head.cell
	q.head = q.head.next
	return cell
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}

func main() {
	var fin, fout *os.File
	fin, err := os.Open("input.txt")

	if errors.Is(err, os.ErrNotExist) {
		fin, _ = os.Open("labyrinth.in")
		fout, _ = os.Create("labyrinth.out")
	} else {
		fout, _ = os.Create("output.txt")
	}

	graph := &Graph{}
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	info := strings.Fields(scanner.Text())
	h, _ := strconv.Atoi(info[0])
	w, _ := strconv.Atoi(info[1])
	graph.construct(w, h)

	for i := 0; i < graph.height; i++ {
		graph.matrix[i] = make([]*Cell, graph.width)
		scanner.Scan()
		for j := 0; j < graph.width; j++ {
			switch scanner.Text()[j] {
			case 'S':
				graph.matrix[i][j] = &Cell{x: i, y: j}
				graph.start = []int{i, j}
			case 'T':
				graph.matrix[i][j] = &Cell{state: true, end: true, x: i, y: j}
			case '.':
				graph.matrix[i][j] = &Cell{state: true, x: i, y: j}
			default:
				graph.matrix[i][j] = &Cell{x: i, y: j}
			}
		}
	}

	q := &Queue{}
	var end *Cell
	q.push(graph.matrix[graph.start[0]][graph.start[1]])
	for !q.isEmpty() {
		cell := q.pop()
		if cell.end {
			end = cell
			break
		}
		if cell.x > 0 {
			oneUp := graph.matrix[cell.x-1][cell.y]
			if oneUp.state && oneUp.route == "" {
				oneUp.route = cell.route + "U"
				q.push(oneUp)
			}
		}
		if cell.x < graph.height-1 {
			oneDown := graph.matrix[cell.x+1][cell.y]
			if oneDown.state && oneDown.route == "" {
				oneDown.route = cell.route + "D"
				q.push(oneDown)
			}
		}
		if cell.y > 0 {
			oneLeft := graph.matrix[cell.x][cell.y-1]
			if oneLeft.state && oneLeft.route == "" {
				oneLeft.route = cell.route + "L"
				q.push(oneLeft)
			}
		}
		if cell.y < graph.width-1 {
			oneRight := graph.matrix[cell.x][cell.y+1]
			if oneRight.state && oneRight.route == "" {
				oneRight.route = cell.route + "R"
				q.push(oneRight)
			}
		}
	}

	if end != nil {
		fmt.Fprintln(fout, len(end.route))
		fmt.Fprintln(fout, end.route)
	} else {
		fmt.Fprintln(fout, -1)
	}
}
