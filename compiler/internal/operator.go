package internal

import "pubql/compiler/parser"

type Operator int

const (
	UnknownOperator Operator = iota
	Add
	Subtract
	Divide
	Multiply
)

func CheckAndReturnOperator(ctx *parser.ExprContext) Operator {
	if ctx.SUM() != nil {
		return Add
	} else if ctx.DIFFERENCE() != nil {
		return Subtract
	} else if ctx.MULTIPLY() != nil {
		return Multiply
	} else if ctx.DIVIDE() != nil {
		return Divide
	}

	return UnknownOperator
}
