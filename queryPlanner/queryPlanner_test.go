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

func assertEquality(t *testing.T, a any, b any) {
	if diff := cmp.Diff(a, b); diff != "" {
		t.Errorf("Expression mismatch (-expected +got):\n%s", diff)
	}
}

// TODO: Add tests for PlanQuery

type predicateTestCase struct {
	input             *parser.WhereClauseContext
	expectedPredicate Predicate
}

func createWhereContext(input string) *parser.WhereClauseContext {
	stream := buildStream(input)
	p := parser.NewpgqlParser(stream)
	ctx := p.WhereClause()
	whereCtx, _ := ctx.(*parser.WhereClauseContext)
	return whereCtx
}

func createPredicateContext(input string) *parser.PredicateContext {
	stream := buildStream(input)
	p := parser.NewpgqlParser(stream)
	ctx := p.Predicate()
	predCtx, _ := ctx.(*parser.PredicateContext)
	return predCtx
}

func TestPredicate(t *testing.T) {
	singlePredicateStatement := createWhereContext("where kills > 2")
	singlePredicateWant := Predicate{
		Type:               PredComparison,
		LeftExpr:           &Expression{Type: ExprMeasure, Measure: "kills"},
		RightExpr:          &Expression{Type: ExprLiteral, LiteralValue: "2"},
		ComparisonOperator: GreaterThan,
	}

	logicalPredicate := createWhereContext("where kills > 2 or damage < 100")
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

	notPredicate := createWhereContext("where not assists = 6")
	notPredicateWant := Predicate{
		Type: PredNot,
		LeftPred: &Predicate{
			Type:               PredComparison,
			LeftExpr:           &Expression{Type: ExprMeasure, Measure: "assists"},
			RightExpr:          &Expression{Type: ExprLiteral, LiteralValue: "6"},
			ComparisonOperator: Equal,
		},
	}

	parenPredicate := createWhereContext("where (rescues != 2 or damage >= 900) and kills <= 3")
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
		q := QueryPlanner{}
		got := q.VisitWhereClause(testCase.input)

		assertEquality(t, testCase.expectedPredicate, got)
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

	scope := createExprContext("ben:kills / .5")
	scopeWant := Expression{
		Type:     ExprBinaryOp,
		Operator: Divide,
		Left: &Expression{
			Type: ExprScope,
			Scope: Scope{
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
		{*scope, scopeWant, analyzer.Analyzer{}},
	}

	for _, testCase := range testCases {
		q := QueryPlanner{
			Analyzer: testCase.a,
		}
		expr := q.VisitExpr(&testCase.input).(Expression)
		assertEquality(t, testCase.expected, expr)
	}
}

func createScopeContext(scope string) *parser.ScopeContext {
	stream := buildStream(scope)
	p := parser.NewpgqlParser(stream)
	ctx := p.Scope()
	scopeCtx, _ := ctx.(*parser.ScopeContext)
	return scopeCtx
}

type scopeTestCase struct {
	name     string
	input    parser.ScopeContext
	expected Scope
	a        analyzer.Analyzer
}

func TestScope(t *testing.T) {
	game := createScopeContext("g:damage")
	gameWant := Scope{EntityType: GameEntity, EntityValue: "g", Measure: "damage"}

	aggregateGameWant := Scope{EntityType: GameAggregateEntity, EntityValue: "g", Measure: "damage"}
	aggregateGameAnalyzer := analyzer.Analyzer{
		Grain: analyzer.Grain{BaseDimension: analyzer.TeamDimension},
	}

	player := createScopeContext("cody:assists")
	playerWant := Scope{EntityType: PlayerEntity, EntityValue: "cody", Measure: "assists"}

	average := createScopeContext("ben:average kills")
	averageWant := Scope{EntityType: PlayerEntity, EntityValue: "ben", Aggregate: Average, Measure: "kills"}

	// TODO: Should we enforce this only being in for team during query planning?
	sum := createScopeContext("g: total damage")
	sumWant := Scope{EntityType: GameEntity, EntityValue: "g", Aggregate: Sum, Measure: "damage"}

	max := createScopeContext("trenton: max damage")
	maxWant := Scope{EntityType: PlayerEntity, EntityValue: "trenton", Aggregate: Max, Measure: "damage"}

	min := createScopeContext("isaac: min kills")
	minWant := Scope{EntityType: PlayerEntity, EntityValue: "isaac", Aggregate: Min, Measure: "kills"}

	testCases := []scopeTestCase{
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
		scope := q.VisitScope(&testCase.input).(Scope)
		assertEquality(t, testCase.expected, scope)
	}
}

type createProjectionTestCase struct {
	name                string
	a                   analyzer.Analyzer
	expectedProjections []Projection
}

func TestCreateProjections(t *testing.T) {
	measureProj := analyzer.Analyzer{
		Projections: []analyzer.Projection{
			{Measure: "kills"},
		},
	}
	measureProjWant := []Projection{{Type: ProjMeasure, Measure: "kills"}}

	scopeProj := analyzer.Analyzer{
		Projections: []analyzer.Projection{
			{Scope: createScopeContext("ben:damage")},
		},
	}
	scopeWant := []Projection{
		{
			Type:  ProjScope,
			Scope: Scope{EntityType: PlayerEntity, EntityValue: "ben", Measure: "damage"},
		},
	}

	aggregateProj := analyzer.Analyzer{
		Projections: []analyzer.Projection{
			{Aggregate: "max", Measure: "assists"},
		},
	}
	aggregateWant := []Projection{{Type: ProjAggregate, Aggregate: Max, Measure: "assists"}}

	expressionProj := analyzer.Analyzer{
		Projections: []analyzer.Projection{
			{Expression: createExprContext("kills + 2"), Identifier: "test"},
		},
	}
	expressionProjWant := []Projection{
		{
			Type: ProjExpr,
			Name: "test",
			Expression: Expression{
				Type:     ExprBinaryOp,
				Operator: Add,
				Left:     &Expression{Type: ExprMeasure, Measure: "kills"},
				Right:    &Expression{Type: ExprLiteral, LiteralValue: "2"},
			},
		},
	}

	predicateProj := analyzer.Analyzer{
		Projections: []analyzer.Projection{
			{Predicate: createPredicateContext("damage > 300"), Identifier: "test"},
		},
	}
	predicateProjWant := []Projection{
		{
			Type: ProjPred,
			Name: "test",
			Predicate: Predicate{
				Type:               PredComparison,
				ComparisonOperator: GreaterThan,
				LeftExpr:           &Expression{Type: ExprMeasure, Measure: "damage"},
				RightExpr:          &Expression{Type: ExprLiteral, LiteralValue: "300"},
			},
		},
	}

	multipleProjections := analyzer.Analyzer{
		Projections: []analyzer.Projection{
			{Measure: "kills"},
			{Aggregate: "min", Measure: "damage"},
		},
	}
	multipleProjectionsWant := []Projection{
		{Type: ProjMeasure, Measure: "kills"},
		{Type: ProjAggregate, Aggregate: Min, Measure: "damage"},
	}

	testCases := []createProjectionTestCase{
		{"Measure projection", measureProj, measureProjWant},
		{"Scope projection", scopeProj, scopeWant},
		{"Aggregate projection", aggregateProj, aggregateWant},
		{"Expression projection", expressionProj, expressionProjWant},
		{"Predicate projection", predicateProj, predicateProjWant},
		{"Multiple projections", multipleProjections, multipleProjectionsWant},
	}

	for _, testCase := range testCases {
		q := QueryPlanner{
			Analyzer: testCase.a,
			QueryPlan: QueryPlan{
				Projections: make([]Projection, 0),
			},
		}
		q.createProjections()
		assertEquality(t, testCase.expectedProjections, q.QueryPlan.Projections)
	}
}
