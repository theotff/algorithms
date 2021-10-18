package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var a, b int
	num_str, _ := ioutil.ReadFile("aplusbb.in")
	fmt.Sscanf(string(num_str), "%d %d", &a, &b)
	fout, _ := os.Create("aplusbb.out")
	fout.WriteString(fmt.Sprint(a + b*b))
	fout.Close()
}
