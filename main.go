package main

import (
	parsing2 "github.com/abolfazlalz/goeventer/internal/interpreter/grammar"
	"github.com/abolfazlalz/goeventer/internal/interpreter/visitor"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parsing2.NewGoEventerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing2.NewGoEventerParser(stream)
	p.BuildParseTrees = true
	tree := p.Parse()
	v := visitor.NewVisitor()
	v.Visit(tree)
}
