package main

import (
	"fmt"
	"pubql/compiler/analyzer"
	"pubql/compiler/parser"
	queryplanner "pubql/compiler/queryPlanner"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	inputString := `show kills
for player ben
where kills > 500 or g:damage = 10`
	input := antlr.NewInputStream(inputString)

	lexer := parser.NewpgqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewpgqlParser(stream)
	tree := parser.Statement()

	analyzer, err := analyzer.Analyze(tree)
	if err != nil {
		fmt.Println(err)
		return
	}
	queryplanner.PlanQuery(analyzer, tree)
}
