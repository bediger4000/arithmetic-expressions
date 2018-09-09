# Arithmetic Expressions

Golang version of a simple algebraic-order-of-operations
arithmetic expression evaluator.

## Build

    $ cd $GOPATH/src
    $ git clone https://github.com/bediger4000/arithmetic-expressions.git
    $ cd arithmetic-expressions
    $ go build arithmetic-expressions

`arithmetic-expressions` parses and prints a single arithmetic expression,
passed to `arithmetic-expressions` as a command line arguments:

    $ ./arithmetic-expressions '1 + 3*4'
    "1 + (3 * 4)"
    /* 13 */

With a `-g` command line flag,
`arithmetic-expressions` prints a [GraphViz](http://graphviz.org/) `dot` format
representation of the parse tree,
evaluates the parse tree and prints the value.
You would do something like this:

    $ ./arithmetic-expressions -g '1 + 3*4' > x.dot
    $ dot -Tpng -o x.png x.dot
    $ feh x.png

You would be rewarded with a deluxe little imag:

![parse tree](https://raw.githubusercontent.com/bediger4000/arithmetic-expressions/master/parse_tree.png)

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
