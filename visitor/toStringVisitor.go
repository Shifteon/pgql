package visitor

import (
	"pubql/parser"
	"strings"

	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type ToStringVisitor struct {
	parser.BasepgqlVisitor
}

func (v *ToStringVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ToStringVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	showClause := v.Visit(ctx.ShowClause())
	forClause := v.Visit(ctx.ForClause())
	// byClause := v.Visit(ctx.ByClause())
	whereClause := v.Visit(ctx.WhereClause())
	// sortByClause := v.Visit(ctx.SortByClause())

	return fmt.Sprintf("%v %v %v", showClause, forClause, whereClause)
}

func (v *ToStringVisitor) VisitShowClause(ctx *parser.ShowClauseContext) interface{} {
	showClause := make([]string, 1)
	for _, child := range ctx.GetChildren() {
		res := v.Visit(child.(antlr.ParseTree))
		showClause = append(showClause, res.(string))
	}

	return strings.Join(showClause, " ")
}

func (v *ToStringVisitor) VisitShowBody(ctx *parser.ShowBodyContext) interface{} {
	showBody := make([]string, 1)
	for _, child := range ctx.GetChildren() {
		res := v.Visit(child.(antlr.ParseTree))
		showBody = append(showBody, res.(string))
	}

	return strings.Join(showBody, " ")
}

func (v *ToStringVisitor) VisitShowFragment(ctx *parser.ShowFragmentContext) interface{} {
	showFragment := make([]string, 1)
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)

		if terminal, ok := child.(antlr.TerminalNode); ok {
			showFragment = append(showFragment, terminal.GetText())
		} else if ruleCtx, ok := child.(antlr.ParserRuleContext); ok {
			showFragment = append(showFragment, v.Visit(ruleCtx).(string))
		}
	}

	return strings.Join(showFragment, "")
}

func (v *ToStringVisitor) VisitShowFunction(ctx *parser.ShowFunctionContext) interface{} {
	return ctx.GetText()
}

func (v *ToStringVisitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	expr := make([]string, 1)
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)

		if terminal, ok := child.(antlr.TerminalNode); ok {
			expr = append(expr, terminal.GetText())
		} else if ruleCtx, ok := child.(antlr.ParserRuleContext); ok {
			expr = append(expr, v.Visit(ruleCtx).(string))
		}
	}

	return strings.Join(expr, "")
}

func (v *ToStringVisitor) VisitMeasure(ctx *parser.MeasureContext) interface{} {
	return ctx.GetText()
}
