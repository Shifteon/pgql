package main

import (
	"fmt"
	"pubql/parser"
	"pubql/visitor"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	inputString := "show average kills for team \"test\" where kills + 2 > damage + 3;"
	input := antlr.NewInputStream(inputString)

	lexer := parser.NewpgqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewpgqlParser(stream)
	tree := parser.Statement()

	printVisitor := &visitor.ToStringVisitor{}
	result := printVisitor.Visit(tree)

	fmt.Println(result)
}
