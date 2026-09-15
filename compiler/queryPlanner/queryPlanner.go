package queryplanner

import (
	"pubql/compiler/analyzer"
	"pubql/compiler/internal"
	"pubql/compiler/parser"
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

type Scope struct {
	EntityType  EntityType
	EntityValue string
	Measure     string
	Aggregate
}

type ExpressionType int

const (
	ExprLiteral ExpressionType = iota
	ExprMeasure
	ExprScope
	ExprAggregate
	ExprBinaryOp // a + b, a / b
)

type Expression struct {
	Type ExpressionType

	// For binary arithmetic
	Left     *Expression
	Right    *Expression
	Operator internal.Operator

	// Leaf data
	LiteralValue string
	Measure      string
	Aggregate    Aggregate
	Scope        Scope
	Identifier   string

	IsWithinParens bool
	GrainLevel     analyzer.GrainLevel
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
	GrainLevel     analyzer.GrainLevel
}

type ProjectionType int

const (
	ProjMeasure ProjectionType = iota
	ProjAggregate
	ProjScope
	ProjExpr
	ProjPred
)

type Projection struct {
	// Metadata
	Type ProjectionType
	Name string

	Predicate  Predicate
	Expression Expression
	Scope      Scope
	Measure    string
	Aggregate  Aggregate
}

type QueryPlan struct {
	Grain        analyzer.Grain
	Predicate    Predicate
	IsSequential bool
	Projections  []Projection
}

type QueryPlanner struct {
	parser.BasepgqlVisitor
	QueryPlan QueryPlan
	// TODO: Should this be private?
	Analyzer analyzer.Analyzer
}

func PlanQuery(a analyzer.Analyzer, tree antlr.ParseTree) QueryPlan {
	queryPlanner := QueryPlanner{
		QueryPlan: QueryPlan{
			Grain:        a.Grain,
			IsSequential: a.IsSequential,
			Projections:  make([]Projection, 0),
		},
		Analyzer: a,
	}

	queryPlanner.Visit(tree)
	queryPlanner.createProjections()

	return queryPlanner.QueryPlan
}

func (q *QueryPlanner) createProjections() {
	for _, projection := range q.Analyzer.Projections {
		queryProjection := Projection{}
		if projection.Aggregate != "" {
			queryProjection.Type = ProjAggregate
			// TODO: If I move the Aggregate type to a different file I can do this in the analyzer instead
			switch projection.Aggregate {
			case "average":
				queryProjection.Aggregate = Average
			case "total":
				queryProjection.Aggregate = Sum
			case "min":
				queryProjection.Aggregate = Min
			case "max":
				queryProjection.Aggregate = Max
			default:
				return
			}
			queryProjection.Measure = projection.Measure
		} else if projection.Measure != "" {
			queryProjection.Type = ProjMeasure
			queryProjection.Measure = projection.Measure
		} else if projection.Expression.ExprCtx != nil {
			queryProjection.Type = ProjExpr
			// TODO: Is there a better place/way to do this type assertion?
			queryProjection.Expression = q.VisitExpr(projection.Expression.ExprCtx).(Expression)
			queryProjection.Name = projection.Identifier
		} else if projection.Predicate.PredCtx != nil {
			queryProjection.Type = ProjPred
			// TODO: Is there a better place/way to do this type assertion?
			queryProjection.Predicate = q.VisitPredicate(projection.Predicate.PredCtx).(Predicate)
			queryProjection.Name = projection.Identifier
		} else if projection.Scope != nil {
			queryProjection.Type = ProjScope
			// TODO: Is there a better place/way to do this type assertion?
			queryProjection.Scope = q.VisitScope(projection.Scope.(*parser.ScopeContext)).(Scope)
		}

		q.QueryPlan.Projections = append(q.QueryPlan.Projections, queryProjection)
	}
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
	if ctx.Predicate() != nil {
		q.QueryPlan.Predicate = q.Visit(ctx.Predicate()).(Predicate)
		// makes it easier to test
		return q.QueryPlan.Predicate
	}

	return true
}

func (q *QueryPlanner) VisitPredicate(ctx *parser.PredicateContext) any {
	predicate := Predicate{}
	predicate.GrainLevel = q.Analyzer.GetPredicateGrainLevel(ctx)

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

func (q *QueryPlanner) VisitExpr(ctx *parser.ExprContext) any {
	expression := Expression{}
	// This is not efficient since we walk the tree twice
	// but it is better for readability etc and I don't think it will be an issue
	expression.GrainLevel = q.Analyzer.GetExpressionGrainLevel(ctx)

	if ctx.LPAREN() != nil {
		expression = q.Visit(ctx.Expr(0)).(Expression)
		expression.IsWithinParens = true
	} else if ctx.IDENTIFIER() != nil {
		expr := q.Analyzer.Identifiers[ctx.IDENTIFIER().GetText()].Expr
		if expr.ExprCtx != nil {
			expression = q.Visit(expr.ExprCtx).(Expression)
			expression.Identifier = ctx.IDENTIFIER().GetText()
		}
	} else if operator := internal.CheckAndReturnOperator(ctx); operator != internal.UnknownOperator {
		expression.Operator = operator
		lhs := q.Visit(ctx.Expr(0)).(Expression)
		rhs := q.Visit(ctx.Expr(1)).(Expression)
		expression.Left = &lhs
		expression.Right = &rhs
		expression.Type = ExprBinaryOp
	} else { //leaf
		if ctx.Scope() != nil {
			expression.Type = ExprScope
			expression.Scope = q.Visit(ctx.Scope()).(Scope)
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

func (q *QueryPlanner) VisitScope(ctx *parser.ScopeContext) any {
	scope := Scope{}

	identifier := strings.ToLower(ctx.IDENTIFIER().GetText())
	scope.EntityValue = identifier
	if identifier == "g" {
		if q.Analyzer.Grain.BaseDimension == analyzer.TeamDimension {
			scope.EntityType = GameAggregateEntity
		} else {
			scope.EntityType = GameEntity
		}
	} else {
		scope.EntityType = PlayerEntity
	}
	if ctx.Measure() != nil {
		scope.Measure = strings.ToLower(ctx.Measure().GetText())
	}
	if ctx.AggregateFunction() != nil {
		aggr := ctx.AggregateFunction()
		if aggr.Measure() != nil {
			scope.Measure = strings.ToLower(aggr.Measure().GetText())
		}
		if aggr.AVERAGE() != nil {
			scope.Aggregate = Average
		} else if aggr.TOTAL() != nil {
			scope.Aggregate = Sum
		} else if aggr.MAX() != nil {
			scope.Aggregate = Max
		} else if aggr.MIN() != nil {
			scope.Aggregate = Min
		}
	}

	return scope
}
