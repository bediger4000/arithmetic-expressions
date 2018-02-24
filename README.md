# Arithmetic Expressions

Golang version of a simple algebraic-order-of-operations
arithmetic expression evaluator.

## Build

	$ GOPATH=$PWD go build expr
	$ GOPATH=$PWD go build eval

`expr` parses and prints a single arithmetic expression,
passed to `expr` as a command line arguments:

    $ ./expr '1 + 3*4'
    "1 + (3 * 4)"

`eval` does much the same as `expr`, but it also evaluates
the successfully parsed command line string:

    $ ./eval '1 + 3*4'
    "1 + (3 * 4)"
    13

## Design

Recursive descent parser, implemented by `lexer.go` and `parser.go`.
I started with lexer and parser code from
[Semantic Tableaux in Go](https://github.com/bediger4000/tableaux-in-go)
which comprehends propositional logic expressions. Luckily, propositional
logic has a lot in common with simple arithmethic expressions
on the syntactic and grammar levels. In fact, I had originally
cribbed the propositional logic grammar from an article on
evaluating arithmetic expressions.
Grammatical precedence of the oparations implements
algebraic order of operations

This code could be a transliteration of C code to do the same
thing, barring some Go idioms. Nothing in particular uses Golang
features (interfaces, goroutines) to do work.

## Why

I need a simple, working, debugged lexer and parser I can start
with to do things
like make an object oriented arithmetic expression evaluator,
or to implement other forms of calculation.
