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
I 
