package main

import (
	"arithmetic/lexer"
	"arithmetic/parser"
	"bytes"
	"fmt"
	"os"
)

func main() {

	graphVizOutput := false
	str := os.Args[1]

	if str == "-g" {
		graphVizOutput = true
		str = os.Args[2]
	}

	expr := bytes.NewBufferString(str + "\n")
	lxr := lexer.NewFromFile(expr)
	psr := parser.New(lxr)
	xpr := psr.Parse()

	if graphVizOutput {
		xpr.GraphNode(os.Stdout)
	} else {
		fmt.Printf("%q\n", xpr)
	}

	fmt.Printf("/* %v */\n", xpr.Eval().Const)
}
