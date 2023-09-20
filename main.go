package main

import (
	"github.com/abolfazlalz/goeventer/internal/interpreter"
	"github.com/abolfazlalz/goeventer/parsing"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parsing.NewGoEventerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewGoEventerParser(stream)
	p.BuildParseTrees = true
	tree := p.Parse()
	v := interpreter.NewVisitor()
	v.Visit(tree)
}
