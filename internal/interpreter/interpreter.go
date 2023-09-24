package interpreter

import (
	"github.com/abolfazlalz/goeventer/internal/interpreter/grammar"
	"github.com/abolfazlalz/goeventer/internal/interpreter/visitor"
	"github.com/antlr4-go/antlr/v4"
)

func Code(code string) {
	input := antlr.NewInputStream(code)
	run(input)
}

func File(file string) {
	input, _ := antlr.NewFileStream(file)
	run(input)
}

func run(input antlr.CharStream) {
	lexer := grammar.NewGoEventerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := grammar.NewGoEventerParser(stream)
	p.BuildParseTrees = true
	tree := p.Parse()
	v := visitor.NewVisitor()
	v.Visit(tree)
}
