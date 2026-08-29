package main

import (
	"fmt"
	"pubql/analyzer"
	"pubql/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	inputString := `show average kills
for team nope 
where kills + 2 > damage + 3`
	input := antlr.NewInputStream(inputString)

	lexer := parser.NewpgqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewpgqlParser(stream)
	tree := parser.Statement()

	err := analyzer.Analyze(tree)
	if err != nil {
		fmt.Println(err)
		return
	}
}
