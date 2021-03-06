package node

// Parse tree - a binary tree of objects of type Node,
// and associated utility functions and methods.

import (
	"arithmetic/lexer"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

// Node has all elements exported, everything reaches inside instances
// of Node to find things out, or to change Left and Right. Private
// elements would cost me gross ol' getter and setter boilerplate.
type Node struct {
	Op    lexer.TokenType
	Const int
	Left  *Node
	Right *Node
}

// NewOpNode creates interior nodes of a parse tree, which will
// all have a +, -, *, / operator associated
func NewOpNode(op lexer.TokenType) *Node {
	return &Node{Op: op}
}

// NewIdentNode creates leaf nodes of a parse tree, which should all be
// lexer.CONSTANT identifier nodes.
func NewConstantNode(stringrepresentation string) *Node {
	var n Node
	n.Op = lexer.CONSTANT
	n.Const, _ = strconv.Atoi(stringrepresentation)
	return &n
}

func (p *Node) Eval() *Node {
	switch p.Op {
	case lexer.CONSTANT:
		return p
	case lexer.PLUS:
		return &Node{Const: p.Left.Eval().Const + p.Right.Eval().Const}
	case lexer.MINUS:
		return &Node{Const: p.Left.Eval().Const - p.Right.Eval().Const}
	case lexer.DIVIDE:
		return &Node{Const: p.Left.Eval().Const / p.Right.Eval().Const}
	case lexer.MULT:
		return &Node{Const: p.Left.Eval().Const * p.Right.Eval().Const}
	}
	return nil
}

// Print puts a human-readable, nicely formatted string representation
// of a parse tree onto the io.Writer, w.  Essentially just an in-order
// traversal of a binary tree, with accommodating a few oddities, like
// parenthesization, and the "~" (not) operator being a prefix.
func (p *Node) Print(w io.Writer) {

	if p.Left != nil {
		printParen := false
		if p.Left.Op != lexer.CONSTANT {
			fmt.Fprintf(w, "(")
			printParen = true
		}
		p.Left.Print(w)
		if printParen {
			fmt.Fprintf(w, ")")
		}
	}

	var oper rune
	switch p.Op {
	case lexer.MULT:
		oper = '*'
	case lexer.DIVIDE:
		oper = '/'
	case lexer.PLUS:
		oper = '+'
	case lexer.MINUS:
		oper = '-'
	case lexer.CONSTANT:
		oper = 0
	}
	if oper != 0 {
		fmt.Fprintf(w, " %c ", oper)
	}

	if p.Op == lexer.CONSTANT {
		fmt.Fprintf(w, "%d", p.Const)
	}

	if p.Right != nil {
		printParen := false
		if p.Right.Op != lexer.CONSTANT {
			fmt.Fprintf(w, "(")
			printParen = true
		}
		p.Right.Print(w)
		if printParen {
			fmt.Fprintf(w, ")")
		}
	}
}

// ExpressionToString creates a Golang string with a human readable
// representation of a parse tree in it.
func ExpressionToString(root *Node) string {
	var sb bytes.Buffer
	root.Print(&sb)
	return sb.String()
}

func (p *Node) String() string {
	return ExpressionToString(p)
}

func (p *Node) graphNode(w io.Writer) {

	var label string

	switch p.Op {
	case lexer.CONSTANT:
		label = fmt.Sprintf("%d", p.Const)
	case lexer.MINUS:
		label = "-"
	case lexer.PLUS:
		label = "+"
	case lexer.DIVIDE:
		label = "/"
	case lexer.MULT:
		label = "*"
	}

	fmt.Fprintf(w, "n%p [label=\"%s\"];\n", p, label)

	if p.Left != nil {
		p.Left.graphNode(w)
		fmt.Fprintf(w, "n%p -> n%p;\n", p, p.Left)
	}
	if p.Right != nil {
		p.Right.graphNode(w)
		fmt.Fprintf(w, "n%p -> n%p;\n", p, p.Right)
	}
}

// GraphNode puts a dot-format text representation of
// a parse tree on w io.Writer.
func (p *Node) GraphNode(w io.Writer) {
	fmt.Fprintf(w, "digraph g {\n")
	p.graphNode(w)
	fmt.Fprintf(w, "}\n")
}
