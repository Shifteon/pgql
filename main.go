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
by game
where kills + 2 > damage + 3 or (damage < 2 and 2 < 3 or (5 = 5 and 2 < 3))`
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
