package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var a, b int
	numStr, _ := ioutil.ReadFile("aplusbb.in")
	fmt.Sscanf(string(numStr), "%d %d", &a, &b)
	fout, _ := os.Create("aplusbb.out")
	fout.WriteString(fmt.Sprint(a + b*b))
	fout.Close()
}
