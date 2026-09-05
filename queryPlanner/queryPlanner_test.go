package queryplanner

import (
	"pubql/analyzer"
	"pubql/parser"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/google/go-cmp/cmp"
)

// TODO: Put this somewhere both tests can access it
func buildStream(input string) *antlr.CommonTokenStream {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewpgqlLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	return stream
}

type predicateTestCase struct {
	input             string
	expectedPredicate Predicate
}

func TestPredicate(t *testing.T) {
	singlePredicateStatement := `show kills
for team ic
where kills > 2`
	singlePredicateWant := Predicate{
		Type:               PredComparison,
		LeftExpr:           &Expression{Type: ExprMeasure, Measure: "kills"},
		RightExpr:          &Expression{Type: ExprLiteral, LiteralValue: "2"},
		ComparisonOperator: GreaterThan,
	}

	logicalPredicate := `show kills
for team ic
where kills > 2 or damage < 100`
	logicalPredicateWant := Predicate{
		Type: PredLogical,
		LeftPred: &Predicate{
			Type:               PredComparison,
			LeftExpr:           &Expression{Type: ExprMeasure, Measure: "kills"},
			RightExpr:          &Expression{Type: ExprLiteral, LiteralValue: "2"},
			ComparisonOperator: GreaterThan,
		},
		RightPred: &Predicate{
			Type:               PredComparison,
			LeftExpr:           &Expression{Type: ExprMeasure, Measure: "damage"},
			RightExpr:          &Expression{Type: ExprLiteral, LiteralValue: "100"},
			ComparisonOperator: LesserThan,
		},
		LogicalOperator: LogicalOr,
	}

	notPredicate := `show kills
for team ic
where not assists = 6`
	notPredicateWant := Predicate{
		Type: PredNot,
		LeftPred: &Predicate{
			Type:               PredComparison,
			LeftExpr:           &Expression{Type: ExprMeasure, Measure: "assists"},
			RightExpr:          &Expression{Type: ExprLiteral, LiteralValue: "6"},
			ComparisonOperator: Equal,
		},
	}

	parenPredicate := `show kills
for team ic
where (rescues != 2 or damage >= 900) and kills <= 3`
	parenPredicateWant := Predicate{
		Type:            PredLogical,
		LogicalOperator: LogicalAnd,
		LeftPred: &Predicate{
			Type:            PredLogical,
			LogicalOperator: LogicalOr,
			IsWithinParens:  true,
			LeftPred: &Predicate{
				Type:               PredComparison,
				LeftExpr:           &Expression{Type: ExprMeasure, Measure: "rescues"},
				RightExpr:          &Expression{Type: ExprLiteral, LiteralValue: "2"},
				ComparisonOperator: NotEqual,
			},
			RightPred: &Predicate{
				Type:               PredComparison,
				LeftExpr:           &Expression{Type: ExprMeasure, Measure: "damage"},
				RightExpr:          &Expression{Type: ExprLiteral, LiteralValue: "900"},
				ComparisonOperator: GreaterOrEqual,
			},
		},
		RightPred: &Predicate{
			Type:               PredComparison,
			LeftExpr:           &Expression{Type: ExprMeasure, Measure: "kills"},
			RightExpr:          &Expression{Type: ExprLiteral, LiteralValue: "3"},
			ComparisonOperator: LesserOrEqual,
		},
	}

	testCases := []predicateTestCase{
		{singlePredicateStatement, singlePredicateWant},
		{logicalPredicate, logicalPredicateWant},
		{notPredicate, notPredicateWant},
		{parenPredicate, parenPredicateWant},
	}

	for _, testCase := range testCases {
		stream := buildStream(testCase.input)
		p := parser.NewpgqlParser(stream)
		tree := p.Statement()
		// TODO: Instead of calling this should probably just construct an Analyzer{}
		analyzer, _ := analyzer.Analyze(tree)
		got := PlanQuery(analyzer, tree)

		if diff := cmp.Diff(testCase.expectedPredicate, got.Predicate); diff != "" {
			t.Errorf("Predicate mismatch (-expected +got):\n%s", diff)
		}
	}
}

func createExprContext(expr string) *parser.ExprContext {
	stream := buildStream(expr)
	p := parser.NewpgqlParser(stream)
	ctx := p.Expr()
	exprCtx, _ := ctx.(*parser.ExprContext)
	return exprCtx
}

type exprTestCase struct {
	input    parser.ExprContext
	expected Expression
	a        analyzer.Analyzer
}

func TestExpression(t *testing.T) {
	binaryOp := createExprContext("kills + 1")
	binaryOpWant := Expression{
		Type:     ExprBinaryOp,
		Operator: Add,
		Left:     &Expression{Type: ExprMeasure, Measure: "kills"},
		Right:    &Expression{Type: ExprLiteral, LiteralValue: "1"},
	}

	aggregate := createExprContext("average kills")
	aggregateWant := Expression{Type: ExprAggregate, Aggregate: Average, Measure: "kills"}

	parens := createExprContext("(rescues - 2) * 3")
	parensWant := Expression{
		Type:     ExprBinaryOp,
		Operator: Multiply,
		Left: &Expression{
			Type:           ExprBinaryOp,
			IsWithinParens: true,
			Operator:       Subtract,
			Left:           &Expression{Type: ExprMeasure, Measure: "rescues"},
			Right:          &Expression{Type: ExprLiteral, LiteralValue: "2"},
		},
		Right: &Expression{Type: ExprLiteral, LiteralValue: "3"},
	}

	identifiers := createExprContext("ident / kills")
	identifiersWant := Expression{
		Type:     ExprBinaryOp,
		Operator: Divide,
		Left: &Expression{
			Type:       ExprBinaryOp,
			Operator:   Add,
			Identifier: "ident",
			Left:       &Expression{Type: ExprLiteral, LiteralValue: "3"},
			Right:      &Expression{Type: ExprMeasure, Measure: "recalls"},
		},
		Right: &Expression{Type: ExprMeasure, Measure: "kills"},
	}
	identifiersAnalyzer := analyzer.Analyzer{
		Identifiers: map[string]analyzer.IdentifierValue{
			"ident": {Expr: createExprContext("3 + recalls")},
		},
	}

	testCases := []exprTestCase{
		{*binaryOp, binaryOpWant, analyzer.Analyzer{}},
		{*aggregate, aggregateWant, analyzer.Analyzer{}},
		{*parens, parensWant, analyzer.Analyzer{}},
		{*identifiers, identifiersWant, identifiersAnalyzer},
	}

	for _, testCase := range testCases {
		q := QueryPlanner{
			Analyzer: testCase.a,
		}
		expr := q.VisitExpr(&testCase.input).(Expression)
		if diff := cmp.Diff(testCase.expected, expr); diff != "" {
			t.Errorf("Expression mismatch (-expected +got):\n%s", diff)
		}
	}
}
