package main

import (
	"bytes"
	"fmt"
	"os"
	"lexer"
	"parser"
)
func main() {
	str := os.Args[1]

	expr := bytes.NewBufferString(str+"\n")
	lxr := lexer.NewFromFile(expr)
	psr := parser.New(lxr)
	xpr := psr.Parse()

	fmt.Printf("%T\n", xpr)
	xpr.Print(os.Stdout)
}

