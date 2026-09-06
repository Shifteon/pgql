package queryplanner

import (
	"pubql/analyzer"
	"pubql/parser"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

/**
* For planning a query, I basically need to answer a few core questions
* 1. Whose data?
* 2. What level of detail?
* These are both answered by grain. The base grain for whose data and refinements for the level of detail
* 3. Are we looking at data by game, like a sequence? So I need game number
* 4. Filters. There are two types of filters
*    a. What games qualify? Filters that filter out individual games
*    b. What aggregates or buckets qualify? Filters that filter out aggregated data
* 5. What are we actually looking at? The projections, or rather the measures or expressions that will be returned.
*
* For game filters you are either filtering on the game total or individual stat. This matters because if you are
* going to filter on game totals, then you need to calculate the game totals in a CTE first and then you can filter.
* Aggregate filters are simpler because you basically just put them in the having
**/

type EntityType int

const (
	UnkownEntity EntityType = iota
	PlayerEntity
	GameEntity
	GameAggregateEntity
)

type Operator int

const (
	UnknownOperator Operator = iota
	Add
	Subtract
	Divide
	Multiply
)

type ComparisonOperator int

const (
	UnkownComparisonOperator ComparisonOperator = iota
	LesserThan
	GreaterThan
	LesserOrEqual
	GreaterOrEqual
	Equal
	NotEqual
)

type LogicalOperator int

const (
	UnkownLogicalOperator LogicalOperator = iota
	LogicalAnd
	LogicalOr
)

type Aggregate int

const (
	UnkownAggregate Aggregate = iota
	Average
	Sum
	Max
	Min
)

type Specificity struct {
	EntityType  EntityType
	EntityValue string
	Measure     string
	Aggregate
}

type ExpressionType int

const (
	ExprLiteral ExpressionType = iota
	ExprMeasure
	ExprSpecificity
	ExprAggregate
	ExprBinaryOp // a + b, a / b
)

type Expression struct {
	Type ExpressionType

	// For binary arithmetic
	Left     *Expression
	Right    *Expression
	Operator Operator

	// Leaf data
	LiteralValue string
	Measure      string
	Aggregate    Aggregate
	Specificity  Specificity
	Identifier   string

	IsWithinParens bool
}

type PredicateType int

const (
	PredComparison PredicateType = iota
	PredLogical
	PredNot
)

type Predicate struct {
	Type PredicateType

	// Comparison leaf
	LeftExpr           *Expression
	ComparisonOperator ComparisonOperator
	RightExpr          *Expression

	// Logical compound
	LeftPred        *Predicate
	LogicalOperator LogicalOperator
	RightPred       *Predicate

	IsWithinParens bool
}

type flatProjection struct {
	measure   string
	name      string
	aggregate Aggregate
}

type QueryPlan struct {
	Grain analyzer.Grain
	Predicate
	IsSequential bool
	Projections  []flatProjection
}

type QueryPlanner struct {
	parser.BasepgqlVisitor
	QueryPlan     QueryPlan
	Analyzer      analyzer.Analyzer
	inWhereClause bool
}

func PlanQuery(a analyzer.Analyzer, tree antlr.ParseTree) QueryPlan {
	queryPlanner := QueryPlanner{
		QueryPlan: QueryPlan{
			Grain:        a.Grain,
			IsSequential: a.IsSequential,
			Projections:  make([]flatProjection, 0),
		},
		Analyzer: a,
	}

	queryPlanner.Visit(tree)

	return queryPlanner.QueryPlan
}

func (q *QueryPlanner) Visit(tree antlr.ParseTree) any {
	return tree.Accept(q)
}

func (q *QueryPlanner) VisitStatement(ctx *parser.StatementContext) any {
	if ctx.WhereClause() != nil {
		q.Visit(ctx.WhereClause())
	}
	return true
}

func (q *QueryPlanner) VisitWhereClause(ctx *parser.WhereClauseContext) any {
	q.inWhereClause = true
	if ctx.Predicate() != nil {
		q.QueryPlan.Predicate = q.Visit(ctx.Predicate()).(Predicate)
	}

	q.inWhereClause = false
	return true
}

func (q *QueryPlanner) VisitPredicate(ctx *parser.PredicateContext) any {
	if q.inWhereClause {
		predicate := Predicate{}
		if ctx.LPAREN() != nil {
			predicate = q.Visit(ctx.Predicate(0)).(Predicate)
			predicate.IsWithinParens = true
		} else if ctx.LOGICALNOT() != nil {
			predicate.Type = PredNot
			lhs := q.Visit(ctx.Predicate(0)).(Predicate)
			predicate.LeftPred = &lhs
		} else if ctx.LOGICALAND() != nil || ctx.LOGICALOR() != nil {
			predicate.Type = PredLogical
			if ctx.LOGICALAND() != nil {
				predicate.LogicalOperator = LogicalAnd
			} else if ctx.LOGICALOR() != nil {
				predicate.LogicalOperator = LogicalOr
			}
			lhs := q.Visit(ctx.Predicate(0)).(Predicate)
			rhs := q.Visit(ctx.Predicate(1)).(Predicate)
			predicate.LeftPred = &lhs
			predicate.RightPred = &rhs
		} else if ctx.ComparisonOperator() != nil {
			predicate.Type = PredComparison
			comparisonOperator := ctx.ComparisonOperator()

			if comparisonOperator.LESSER() != nil {
				predicate.ComparisonOperator = LesserThan
			} else if comparisonOperator.GREATER() != nil {
				predicate.ComparisonOperator = GreaterThan
			} else if comparisonOperator.LESSEREQUAL() != nil {
				predicate.ComparisonOperator = LesserOrEqual
			} else if comparisonOperator.GREATEREQUAL() != nil {
				predicate.ComparisonOperator = GreaterOrEqual
			} else if comparisonOperator.EQUAL() != nil {
				predicate.ComparisonOperator = Equal
			} else if comparisonOperator != nil {
				predicate.ComparisonOperator = NotEqual
			}

			lhs := q.Visit(ctx.Expr(0)).(Expression)
			rhs := q.Visit(ctx.Expr(1)).(Expression)
			predicate.LeftExpr = &lhs
			predicate.RightExpr = &rhs
		}
		return predicate
	}

	return true
}

func checkAndReturnOperator(ctx *parser.ExprContext) Operator {
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

func (q *QueryPlanner) VisitExpr(ctx *parser.ExprContext) any {
	expression := Expression{}
	if ctx.LPAREN() != nil {
		expression = q.Visit(ctx.Expr(0)).(Expression)
		expression.IsWithinParens = true
	} else if ctx.IDENTIFIER() != nil {
		expr := q.Analyzer.Identifiers[ctx.IDENTIFIER().GetText()].Expr
		if expr != nil {
			expression = q.Visit(expr).(Expression)
			expression.Identifier = ctx.IDENTIFIER().GetText()
		}
	} else if operator := checkAndReturnOperator(ctx); operator != UnknownOperator {
		expression.Operator = operator
		lhs := q.Visit(ctx.Expr(0)).(Expression)
		rhs := q.Visit(ctx.Expr(1)).(Expression)
		expression.Left = &lhs
		expression.Right = &rhs
		expression.Type = ExprBinaryOp
	} else { //leaf
		if ctx.Specificity() != nil {
			expression.Type = ExprSpecificity
			expression.Specificity = q.Visit(ctx.Specificity()).(Specificity)
		} else if ctx.AggregateFunction() != nil {
			expression.Type = ExprAggregate
			aggr := ctx.AggregateFunction()
			if aggr.TOTAL() != nil {
				expression.Aggregate = Sum
			} else if aggr.AVERAGE() != nil {
				expression.Aggregate = Average
			} else if aggr.MAX() != nil {
				expression.Aggregate = Max
			} else if aggr.MIN() != nil {
				expression.Aggregate = Min
			}
			expression.Measure = aggr.Measure().GetText()
		} else if ctx.Measure() != nil {
			expression.Type = ExprMeasure
			expression.Measure = ctx.Measure().GetText()
		} else if ctx.NUMBER() != nil {
			expression.Type = ExprLiteral
			expression.LiteralValue = ctx.NUMBER().GetText()
		}
	}

	return expression
}

func (q *QueryPlanner) VisitSpecificity(ctx *parser.SpecificityContext) any {
	specificity := Specificity{}

	identifier := strings.ToLower(ctx.IDENTIFIER().GetText())
	specificity.EntityValue = identifier
	if identifier == "g" {
		if q.Analyzer.Grain.BaseDimension == analyzer.TeamDimension {
			specificity.EntityType = GameAggregateEntity
		} else {
			specificity.EntityType = GameEntity
		}
	} else {
		specificity.EntityType = PlayerEntity
	}
	if ctx.Measure() != nil {
		specificity.Measure = strings.ToLower(ctx.Measure().GetText())
	}
	if ctx.AggregateFunction() != nil {
		aggr := ctx.AggregateFunction()
		if aggr.Measure() != nil {
			specificity.Measure = strings.ToLower(aggr.Measure().GetText())
		}
		if aggr.AVERAGE() != nil {
			specificity.Aggregate = Average
		} else if aggr.TOTAL() != nil {
			specificity.Aggregate = Sum
		} else if aggr.MAX() != nil {
			specificity.Aggregate = Max
		} else if aggr.MIN() != nil {
			specificity.Aggregate = Min
		}
	}

	return specificity
}
