package main

import (
	"fmt"
	"pubql/analyzer"
	"pubql/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	inputString := `show kills, ic:damage
for player ben
by game
where kills + 2 > damage + 3`
	input := antlr.NewInputStream(inputString)

	lexer := parser.NewpgqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewpgqlParser(stream)
	tree := parser.Statement()

	_, err := analyzer.Analyze(tree)
	if err != nil {
		fmt.Println(err)
		return
	}
}
