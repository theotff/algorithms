package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var a, b int
	numStr, _ := ioutil.ReadFile("aplusb.in")
	fmt.Sscanf(string(numStr), "%d %d", &a, &b)
	fout, _ := os.Create("aplusb.out")
	fout.WriteString(fmt.Sprint(a + b))
	fout.Close()
}
