package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	fin, _ := os.Open("brackets.in")
	scanner := bufio.NewScanner(fin)
	var results []string

	for scanner.Scan() {
		brackets := strings.Split(scanner.Text(), "")
		stack := make([]string, len(brackets))
		stack_index := 0
		for _, elem := range brackets {
			if stack_index != 0 {
				if stack[stack_index-1] == "(" && elem == ")" {
					stack_index -= 1
					stack[stack_index] = ""
				} else if stack[stack_index-1] == "[" && elem == "]" {
					stack_index -= 1
					stack[stack_index] = ""
				} else {
					stack[stack_index] = elem
					stack_index += 1
				}
			} else {
				stack[stack_index] = elem
				stack_index += 1
			}
		}
		state := true
		for _, elem := range stack {
			if elem != "" {
				state = false
			}
		}
		if state {
			results = append(results, "YES")
		} else {
			results = append(results, "NO")
		}
	}
	fout, _ := os.Create("brackets.out")
	fout.WriteString(strings.Join(results, "\n"))
	fout.Close()
}
