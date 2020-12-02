// example1.go
package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"./parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("a := 1;")

	// Create the Lexer
	lexer := parser.NewAssignLexer(is)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewAssignParser(tokens)

	tree := parser.Assign()

	fmt.Println(tree.ToStringTree([]string{}, parser))
}
