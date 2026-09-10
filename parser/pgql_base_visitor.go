// Code generated from pgql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // pgql

import "github.com/antlr4-go/antlr/v4"

type BasepgqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasepgqlVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitShowClause(ctx *ShowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitShowBody(ctx *ShowBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitShowFragment(ctx *ShowFragmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitForClause(ctx *ForClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitByClause(ctx *ByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitSortByClause(ctx *SortByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitSortBody(ctx *SortBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitDimension(ctx *DimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitMeasure(ctx *MeasureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitAggregateFunction(ctx *AggregateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitSequentialFunction(ctx *SequentialFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitScope(ctx *ScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepgqlVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}
