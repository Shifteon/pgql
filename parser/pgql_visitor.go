// Code generated from pgql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // pgql

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by pgqlParser.
type pgqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by pgqlParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by pgqlParser#showClause.
	VisitShowClause(ctx *ShowClauseContext) interface{}

	// Visit a parse tree produced by pgqlParser#showBody.
	VisitShowBody(ctx *ShowBodyContext) interface{}

	// Visit a parse tree produced by pgqlParser#showFragment.
	VisitShowFragment(ctx *ShowFragmentContext) interface{}

	// Visit a parse tree produced by pgqlParser#forClause.
	VisitForClause(ctx *ForClauseContext) interface{}

	// Visit a parse tree produced by pgqlParser#byClause.
	VisitByClause(ctx *ByClauseContext) interface{}

	// Visit a parse tree produced by pgqlParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by pgqlParser#sortByClause.
	VisitSortByClause(ctx *SortByClauseContext) interface{}

	// Visit a parse tree produced by pgqlParser#sortBody.
	VisitSortBody(ctx *SortBodyContext) interface{}

	// Visit a parse tree produced by pgqlParser#dimension.
	VisitDimension(ctx *DimensionContext) interface{}

	// Visit a parse tree produced by pgqlParser#measure.
	VisitMeasure(ctx *MeasureContext) interface{}

	// Visit a parse tree produced by pgqlParser#aggregateFunction.
	VisitAggregateFunction(ctx *AggregateFunctionContext) interface{}

	// Visit a parse tree produced by pgqlParser#sequentialFunction.
	VisitSequentialFunction(ctx *SequentialFunctionContext) interface{}

	// Visit a parse tree produced by pgqlParser#logicalOperator.
	VisitLogicalOperator(ctx *LogicalOperatorContext) interface{}

	// Visit a parse tree produced by pgqlParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by pgqlParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by pgqlParser#specificity.
	VisitSpecificity(ctx *SpecificityContext) interface{}

	// Visit a parse tree produced by pgqlParser#expr.
	VisitExpr(ctx *ExprContext) interface{}
}
