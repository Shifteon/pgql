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

func assertEquality(t *testing.T, a interface{}, b interface{}) {
	if diff := cmp.Diff(a, b); diff != "" {
		t.Errorf("Expression mismatch (-expected +got):\n%s", diff)
	}
}

// TODO: Add tests for PlanQuery

type predicateTestCase struct {
	input             string
	expectedPredicate Predicate
}

// TODO: Make these tests more like the expression ones so I just call visit predicate
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

		assertEquality(t, testCase.expectedPredicate, got.Predicate)
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

	aggregateSum := createExprContext("total kills")
	aggregateSumWant := Expression{Type: ExprAggregate, Aggregate: Sum, Measure: "kills"}

	aggregateMax := createExprContext("max kills")
	aggregateMaxWant := Expression{Type: ExprAggregate, Aggregate: Max, Measure: "kills"}

	aggregateMin := createExprContext("min kills")
	aggregateMinWant := Expression{Type: ExprAggregate, Aggregate: Min, Measure: "kills"}

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

	specificity := createExprContext("ben:kills / .5")
	specificityWant := Expression{
		Type:     ExprBinaryOp,
		Operator: Divide,
		Left: &Expression{
			Type: ExprSpecificity,
			Specificity: Specificity{
				EntityType:  PlayerEntity,
				EntityValue: "ben",
				Measure:     "kills",
			},
		},
		Right: &Expression{Type: ExprLiteral, LiteralValue: ".5"},
	}

	testCases := []exprTestCase{
		{*binaryOp, binaryOpWant, analyzer.Analyzer{}},
		{*aggregate, aggregateWant, analyzer.Analyzer{}},
		{*aggregateSum, aggregateSumWant, analyzer.Analyzer{}},
		{*aggregateMax, aggregateMaxWant, analyzer.Analyzer{}},
		{*aggregateMin, aggregateMinWant, analyzer.Analyzer{}},
		{*parens, parensWant, analyzer.Analyzer{}},
		{*identifiers, identifiersWant, identifiersAnalyzer},
		{*specificity, specificityWant, analyzer.Analyzer{}},
	}

	for _, testCase := range testCases {
		q := QueryPlanner{
			Analyzer: testCase.a,
		}
		expr := q.VisitExpr(&testCase.input).(Expression)
		assertEquality(t, testCase.expected, expr)
	}
}

func createSpecificityContext(specificity string) *parser.SpecificityContext {
	stream := buildStream(specificity)
	p := parser.NewpgqlParser(stream)
	ctx := p.Specificity()
	specificityCtx, _ := ctx.(*parser.SpecificityContext)
	return specificityCtx
}

type specificityTestCase struct {
	name     string
	input    parser.SpecificityContext
	expected Specificity
	a        analyzer.Analyzer
}

func TestSpecificity(t *testing.T) {
	game := createSpecificityContext("g:damage")
	gameWant := Specificity{EntityType: GameEntity, EntityValue: "g", Measure: "damage"}

	aggregateGameWant := Specificity{EntityType: GameAggregateEntity, EntityValue: "g", Measure: "damage"}
	aggregateGameAnalyzer := analyzer.Analyzer{
		Grain: analyzer.Grain{BaseDimension: analyzer.TeamDimension},
	}

	player := createSpecificityContext("cody:assists")
	playerWant := Specificity{EntityType: PlayerEntity, EntityValue: "cody", Measure: "assists"}

	average := createSpecificityContext("ben:average kills")
	averageWant := Specificity{EntityType: PlayerEntity, EntityValue: "ben", Aggregate: Average, Measure: "kills"}

	// TODO: Should we enforce this only being in for team during query planning?
	sum := createSpecificityContext("g: total damage")
	sumWant := Specificity{EntityType: GameEntity, EntityValue: "g", Aggregate: Sum, Measure: "damage"}

	max := createSpecificityContext("trenton: max damage")
	maxWant := Specificity{EntityType: PlayerEntity, EntityValue: "trenton", Aggregate: Max, Measure: "damage"}

	min := createSpecificityContext("isaac: min kills")
	minWant := Specificity{EntityType: PlayerEntity, EntityValue: "isaac", Aggregate: Min, Measure: "kills"}

	testCases := []specificityTestCase{
		{"Game", *game, gameWant, analyzer.Analyzer{}},
		{"Aggregate game", *game, aggregateGameWant, aggregateGameAnalyzer},
		{"Player", *player, playerWant, analyzer.Analyzer{}},
		{"Average aggregate", *average, averageWant, analyzer.Analyzer{}},
		{"Sum aggregate", *sum, sumWant, analyzer.Analyzer{}},
		{"Max aggregate", *max, maxWant, analyzer.Analyzer{}},
		{"Min aggregate", *min, minWant, analyzer.Analyzer{}},
	}

	for _, testCase := range testCases {
		q := QueryPlanner{
			Analyzer: testCase.a,
		}
		specificity := q.VisitSpecificity(&testCase.input).(Specificity)
		assertEquality(t, testCase.expected, specificity)
	}
}
