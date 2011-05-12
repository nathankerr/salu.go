package main

import (
	"fmt"
	"os"
)

func main() {
	parser := NewParser(os.Stdin)
	ast := parser.Parse()
	fmt.Println(ast)
}
