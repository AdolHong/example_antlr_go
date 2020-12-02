// example1.go
package main

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/c-bata/go-prompt"

	"./parser"
)

type arrayinitListener struct {
	*parser.BaseArrayInitListener
}

func (p *arrayinitListener) EnterInit(c *parser.InitContext) {
	fmt.Print("")
}

func (p *arrayinitListener) exitInit(c *parser.InitContext) {
	fmt.Print("")
}

func (p *arrayinitListener) EnterValue(c *parser.ValueContext) {
	d, _ := strconv.Atoi(c.INT().GetText())
	fmt.Printf("\\u%04x",d)
}


func doSomething(data string) {
	// Setup the input
	is := antlr.NewInputStream(data)

	// Create the Lexer
	lexer := parser.NewArrayInitLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewArrayInitParser(stream)

	var listener arrayinitListener

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Init())

}


func executor(in string) {
	fmt.Print("Answer:")
	doSomething(in)
	fmt.Println("")
}

func completer(in prompt.Document) []prompt.Suggest {
	var ret []prompt.Suggest
	return ret
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("calc"),
	)
	p.Run()
}