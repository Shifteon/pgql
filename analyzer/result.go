package analyzer

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type result struct {
	err error
}

func ok() result {
	return result{}
}

func fail(msg string, token antlr.Token) result {
	lineNumber := token.GetLine()
	charPosition := token.GetTokenSource().GetCharPositionInLine()
	return result{err: fmt.Errorf("line %v : char %v - error: %v", lineNumber, charPosition, msg)}
}

func message(token antlr.BaseToken) {
	fmt.Printf("")
}
