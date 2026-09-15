package analyzer

import (
	"pubql/compiler/parser"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func buildStream(input string) *antlr.CommonTokenStream {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewpgqlLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	return stream
}

func createStatementContext(input string) *parser.StatementContext {
	stream := buildStream(input)
	p := parser.NewpgqlParser(stream)
	ctx := p.Statement()
	stmtCtx, _ := ctx.(*parser.StatementContext)
	return stmtCtx
}

func createForContext(input string) *parser.ForClauseContext {
	stream := buildStream(input)
	p := parser.NewpgqlParser(stream)
	ctx := p.ForClause()
	forCtx, _ := ctx.(*parser.ForClauseContext)
	return forCtx
}

func createByContext(input string) *parser.ByClauseContext {
	stream := buildStream(input)
	p := parser.NewpgqlParser(stream)
	ctx := p.ByClause()
	byCtx, _ := ctx.(*parser.ByClauseContext)
	return byCtx
}

func createShowContext(input string) *parser.ShowClauseContext {
	stream := buildStream(input)
	p := parser.NewpgqlParser(stream)
	ctx := p.ShowClause()
	showCtx, _ := ctx.(*parser.ShowClauseContext)
	return showCtx
}

func createShowFragmentContext(input string) *parser.ShowFragmentContext {
	stream := buildStream(input)
	p := parser.NewpgqlParser(stream)
	ctx := p.ShowFragment()
	fragCtx, _ := ctx.(*parser.ShowFragmentContext)
	return fragCtx
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

func createExprContext(expr string) *parser.ExprContext {
	stream := buildStream(expr)
	p := parser.NewpgqlParser(stream)
	ctx := p.Expr()
	exprCtx, _ := ctx.(*parser.ExprContext)
	return exprCtx
}

func createScopeContext(scope string) *parser.ScopeContext {
	stream := buildStream(scope)
	p := parser.NewpgqlParser(stream)
	ctx := p.Scope()
	scopeCtx, _ := ctx.(*parser.ScopeContext)
	return scopeCtx
}

func createSortByContext(input string) *parser.SortByClauseContext {
	stream := buildStream(input)
	p := parser.NewpgqlParser(stream)
	ctx := p.SortByClause()
	sortCtx, _ := ctx.(*parser.SortByClauseContext)
	return sortCtx
}

func assertEquality(t *testing.T, a any, b any, opts ...cmp.Option) {
	t.Helper()
	if diff := cmp.Diff(a, b, opts...); diff != "" {
		t.Errorf("Mismatch (-expected +got):\n%s", diff)
	}
}

func assertResult(t *testing.T, res any, expectedMsg string) {
	t.Helper()
	r, ok := res.(result)
	if !ok {
		t.Fatalf("Expected result type, got %T", res)
	}
	if expectedMsg == "" {
		if r.err != nil {
			t.Fatalf("Expected no error, got %v", r.err)
		}
		return
	}
	if r.err == nil {
		t.Fatalf("Expected error containing %q, got nil", expectedMsg)
	}
	if !strings.Contains(r.err.Error(), expectedMsg) {
		t.Fatalf("Expected error containing %q, got %q", expectedMsg, r.err.Error())
	}
}

type forTestCase struct {
	name            string
	input           *parser.ForClauseContext
	expectedGrain   Grain
	expectedMessage string
}

func TestForClause(t *testing.T) {
	testCases := []forTestCase{
		{
			name:            "valid team",
			input:           createForContext("for team ic"),
			expectedGrain:   Grain{BaseDimension: TeamDimension, BaseDimensionValue: "ic"},
			expectedMessage: "",
		},
		{
			name:            "valid player",
			input:           createForContext("for player ben"),
			expectedGrain:   Grain{BaseDimension: PlayerDimension, BaseDimensionValue: "ben"},
			expectedMessage: "",
		},
		{
			name:            "invalid dimension",
			input:           createForContext("for win ic"),
			expectedMessage: "Invalid dimension used in the FOR clause!",
		},
		{
			name:            "invalid team",
			input:           createForContext("for team i"),
			expectedMessage: "Unrecognized team i. Expected one of",
		},
		{
			name:            "invalid player",
			input:           createForContext("for player p"),
			expectedMessage: "Unrecognized player p. Expected one of",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := Analyzer{}
			res := a.VisitForClause(tc.input)
			assertResult(t, res, tc.expectedMessage)
			if tc.expectedMessage == "" {
				assertEquality(t, tc.expectedGrain, a.Grain)
			}
		})
	}
}

type byTestCase struct {
	name            string
	input           *parser.ByClauseContext
	a               Analyzer
	expectedGrain   Grain
	expectedScalar  bool
	expectedMessage string
}

func TestByClause(t *testing.T) {
	testCases := []byTestCase{
		{
			name:  "valid single dimension",
			input: createByContext("by player"),
			a: Analyzer{
				Grain: Grain{BaseDimension: TeamDimension, BaseDimensionValue: "ic"},
			},
			expectedGrain: Grain{
				BaseDimension:      TeamDimension,
				BaseDimensionValue: "ic",
				Refinements:        []Dimension{PlayerDimension},
			},
			expectedScalar:  false,
			expectedMessage: "",
		},
		{
			name:  "valid two dimensions",
			input: createByContext("by player, game"),
			a: Analyzer{
				Grain: Grain{BaseDimension: TeamDimension, BaseDimensionValue: "ic"},
			},
			expectedGrain: Grain{
				BaseDimension:      TeamDimension,
				BaseDimensionValue: "ic",
				Refinements:        []Dimension{GameDimension, PlayerDimension},
			},
			expectedScalar:  true,
			expectedMessage: "",
		},
		{
			name:  "valid player with game refinement",
			input: createByContext("by game"),
			a: Analyzer{
				Grain: Grain{BaseDimension: PlayerDimension, BaseDimensionValue: "ben"},
			},
			expectedGrain: Grain{
				BaseDimension:      PlayerDimension,
				BaseDimensionValue: "ben",
				Refinements:        []Dimension{GameDimension},
			},
			expectedScalar:  true,
			expectedMessage: "",
		},
		{
			name:  "invalid dimension",
			input: createByContext("by date"),
			a: Analyzer{
				Grain: Grain{BaseDimension: TeamDimension, BaseDimensionValue: "ic"},
			},
			expectedMessage: "Invalid dimension used in the BY clause! Got date. Expected one of 'team' or 'player'",
		},
		{
			name:  "team already in for",
			input: createByContext("by team"),
			a: Analyzer{
				Grain: Grain{BaseDimension: TeamDimension, BaseDimensionValue: "ic"},
			},
			expectedMessage: "'team' dimension already used in FOR clause!",
		},
		{
			name:  "player already in for",
			input: createByContext("by player"),
			a: Analyzer{
				Grain: Grain{BaseDimension: PlayerDimension, BaseDimensionValue: "ben"},
			},
			expectedMessage: "'player' dimension already used in FOR clause!",
		},
		{
			name:  "game used twice",
			input: createByContext("by game, game"),
			a: Analyzer{
				Grain: Grain{BaseDimension: TeamDimension, BaseDimensionValue: "ic"},
			},
			expectedMessage: "Used the same dimension 'game' more than once!",
		},
		{
			name:  "team used twice",
			input: createByContext("by team, team"),
			a: Analyzer{
				Grain: Grain{BaseDimension: PlayerDimension, BaseDimensionValue: "ben"},
			},
			expectedMessage: "Used the same dimension 'team' more than once!",
		},
		{
			name:  "player used twice",
			input: createByContext("by player, player"),
			a: Analyzer{
				Grain: Grain{BaseDimension: TeamDimension, BaseDimensionValue: "ic"},
			},
			expectedMessage: "Used the same dimension 'player' more than once!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := tc.a
			res := a.VisitByClause(tc.input)
			assertResult(t, res, tc.expectedMessage)
			if tc.expectedMessage == "" {
				assertEquality(t, tc.expectedGrain, a.Grain)
				assertEquality(t, tc.expectedScalar, a.IsScalar)
			}
		})
	}
}

type showTestCase struct {
	name                string
	input               *parser.ShowClauseContext
	a                   Analyzer
	expectedSequential  bool
	expectedProjections []Projection
	checkIdentifiers    func(t *testing.T, a *Analyzer)
	expectedMessage     string
}

func TestShowClause(t *testing.T) {
	testCases := []showTestCase{
		{
			name:  "valid show simple measure",
			input: createShowContext("show kills"),
			a: Analyzer{
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedSequential: false,
			expectedProjections: []Projection{
				{Measure: "kills"},
			},
			expectedMessage: "",
		},
		{
			name:  "sequential function",
			input: createShowContext("show running kills"),
			a: Analyzer{
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedSequential: true,
			expectedProjections: []Projection{
				{Measure: "kills"},
			},
			expectedMessage: "",
		},
		{
			name:  "aggregate in non-scalar",
			input: createShowContext("show max kills"),
			a: Analyzer{
				IsScalar:    false,
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedProjections: []Projection{
				{Aggregate: "max", Measure: "kills"},
			},
			expectedMessage: "",
		},
		{
			name:  "aggregate in scalar",
			input: createShowContext("show average kills"),
			a: Analyzer{
				IsScalar:    true,
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedMessage: "Cannot use aggregate functions in a scalar statement!",
		},
		{
			name:  "identifiers defined in show",
			input: createShowContext("show (5 + 5) yep"),
			a: Analyzer{
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedProjections: []Projection{
				{Identifier: "yep"},
			},
			checkIdentifiers: func(t *testing.T, a *Analyzer) {
				val, exists := a.Identifiers["yep"]
				if !exists {
					t.Fatalf("Expected identifier 'yep' in Identifiers map")
				}
				if val.Expr == nil {
					t.Fatalf("Expected identifier 'yep' to have non-nil Expr")
				}
			},
			expectedMessage: "",
		},
		{
			name:  "duplicate identifier",
			input: createShowContext("show (5 + 5) yep, (2 - 3) yep"),
			a: Analyzer{
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedMessage: "Identifier \"yep\" already used!",
		},
		{
			name:  "predicate defined in show",
			input: createShowContext("show (5 > 5) yep"),
			a: Analyzer{
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedProjections: []Projection{
				{Identifier: "yep"},
			},
			checkIdentifiers: func(t *testing.T, a *Analyzer) {
				val, exists := a.Identifiers["yep"]
				if !exists {
					t.Fatalf("Expected identifier 'yep' in Identifiers map")
				}
				if val.Pred == nil {
					t.Fatalf("Expected identifier 'yep' to have non-nil Pred")
				}
			},
			expectedMessage: "",
		},
		{
			name:  "scope in show fragment",
			input: createShowContext("show ben:kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
				},
				Identifiers: make(map[string]IdentifierValue),
			},
			checkIdentifiers: func(t *testing.T, a *Analyzer) {
				if len(a.Projections) != 1 {
					t.Fatalf("Expected 1 projection, got %d", len(a.Projections))
				}
				if a.Projections[0].Scope == nil {
					t.Fatalf("Expected non-nil Scope on projection")
				}
			},
			expectedMessage: "",
		},
		{
			name:  "reference existing identifier in show",
			input: createShowContext("show yep"),
			a: Analyzer{
				Identifiers: map[string]IdentifierValue{
					"yep": {Expr: createExprContext("5 + 5")},
				},
			},
			checkIdentifiers: func(t *testing.T, a *Analyzer) {
				if len(a.Projections) != 1 {
					t.Fatalf("Expected 1 projection, got %d", len(a.Projections))
				}
				if a.Projections[0].Expression == nil {
					t.Fatalf("Expected non-nil Expression on resolved projection")
				}
			},
			expectedMessage: "",
		},
		{
			name:  "reference undeclared identifier in show",
			input: createShowContext("show nope"),
			a: Analyzer{
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedMessage: "Undeclared identifer: \"nope\"",
		},
		{
			name:  "invalid expression with undeclared identifier",
			input: createShowContext("show (nope > 5) yep"),
			a: Analyzer{
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedMessage: "Undeclared identifer: \"nope\"",
		},
		{
			name:  "invalid expression with aggregate in scalar",
			input: createShowContext("show (average kills / 5) avg"),
			a: Analyzer{
				IsScalar:    true,
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedMessage: "Cannot use aggregate functions in a scalar statement!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := tc.a
			res := a.VisitShowClause(tc.input)
			assertResult(t, res, tc.expectedMessage)
			if tc.expectedMessage == "" {
				assertEquality(t, tc.expectedSequential, a.IsSequential)
				if len(tc.expectedProjections) > 0 {
					ignoreContexts := cmpopts.IgnoreFields(Projection{}, "Expression", "Predicate", "Scope")
					assertEquality(t, tc.expectedProjections, a.Projections, ignoreContexts)
				}
				if tc.checkIdentifiers != nil {
					tc.checkIdentifiers(t, &a)
				}
			}
		})
	}
}

type exprTestCase struct {
	name            string
	input           *parser.ExprContext
	a               Analyzer
	expectedMessage string
}

func TestExpr(t *testing.T) {
	testCases := []exprTestCase{
		{
			name:            "valid binary arithmetic",
			input:           createExprContext("kills / damage"),
			a:               Analyzer{},
			expectedMessage: "",
		},
		{
			name:  "declared identifier in expr",
			input: createExprContext("yep + 10"),
			a: Analyzer{
				Identifiers: map[string]IdentifierValue{
					"yep": {Expr: createExprContext("5 + 5")},
				},
			},
			expectedMessage: "",
		},
		{
			name:  "undeclared identifier in expr",
			input: createExprContext("nope + 10"),
			a: Analyzer{
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedMessage: "Undeclared identifer: \"nope\"",
		},
		{
			name:  "aggregate in scalar mode",
			input: createExprContext("average damage"),
			a: Analyzer{
				IsScalar: true,
			},
			expectedMessage: "Cannot use aggregate functions in a scalar statement!",
		},
		{
			name:  "aggregate in non-scalar mode",
			input: createExprContext("average damage"),
			a: Analyzer{
				IsScalar: false,
			},
			expectedMessage: "",
		},
		{
			name:  "valid scope in expr",
			input: createExprContext("ben:kills + 1"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
				},
			},
			expectedMessage: "",
		},
		{
			name:  "invalid scope in expr",
			input: createExprContext("ben:kills + 1"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ic",
				},
			},
			expectedMessage: "Player \"ben\" is not on Team \"ic\"!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := tc.a
			res := a.VisitExpr(tc.input)
			assertResult(t, res, tc.expectedMessage)
		})
	}
}

type scopeTestCase struct {
	name            string
	input           *parser.ScopeContext
	a               Analyzer
	expectedScalar  bool
	expectedMessage string
}

func TestScope(t *testing.T) {
	testCases := []scopeTestCase{
		{
			name:  "valid team scope player on team",
			input: createScopeContext("ben:kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
				},
			},
			expectedMessage: "",
		},
		{
			name:  "valid team scope with g identifier",
			input: createScopeContext("g:damage"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
				},
			},
			expectedMessage: "",
		},
		{
			name:  "team and team invalid player identifier",
			input: createScopeContext("ib:kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ic",
				},
			},
			expectedMessage: "Identifier in scope cannot be mapped to a player! Got: ib. Expected one of",
		},
		{
			name:  "invalid player identifier",
			input: createScopeContext("nope:kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ic",
				},
			},
			expectedMessage: "Identifier in scope cannot be mapped to a player! Got: nope. Expected one of",
		},
		{
			name:  "team does not have player",
			input: createScopeContext("ben:kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ic",
				},
			},
			expectedMessage: "Player \"ben\" is not on Team \"ic\"!",
		},
		{
			name:  "player statement with scope in show clause",
			input: createScopeContext("g:kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      PlayerDimension,
					BaseDimensionValue: "isaac",
				},
				currentClause: ShowClause,
			},
			expectedMessage: "Scope cannot be used in the SHOW clause of a statement with \"FOR player\"",
		},
		{
			name:  "player statement with invalid scope type",
			input: createScopeContext("ic:kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      PlayerDimension,
					BaseDimensionValue: "ben",
				},
				currentClause: WhereClause,
			},
			expectedMessage: "Invalid scope type used in a statement with \"FOR player\". Got ic.",
		},
		{
			name:  "player statement with valid g scope in where clause",
			input: createScopeContext("g:kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      PlayerDimension,
					BaseDimensionValue: "ben",
				},
				currentClause: WhereClause,
			},
			expectedMessage: "",
		},
		{
			name:  "player statement with aggregate in scope",
			input: createScopeContext("g:average kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      PlayerDimension,
					BaseDimensionValue: "ben",
				},
				currentClause: WhereClause,
			},
			expectedMessage: "Cannot use aggregates in a scope in a statement with \"FOR player\"!",
		},
		{
			name:  "aggregate in scalar outside g scope",
			input: createScopeContext("ben:average kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
				},
				IsScalar: true,
			},
			expectedMessage: "Cannot use an aggregate function within a scalar statement outside of the \"g:\" scope!",
		},
		{
			name:  "scope in show with game dimension sets scalar",
			input: createScopeContext("ben:kills"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
					Refinements:        []Dimension{GameDimension},
				},
				currentClause: ShowClause,
				IsScalar:      false,
			},
			expectedScalar:  true,
			expectedMessage: "",
		},
		{
			name:            "unknown base dimension fatal error",
			input:           createScopeContext("g:kills"),
			a:               Analyzer{},
			expectedMessage: "Fatal error! Somehow there is no dimension in the FOR!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := tc.a
			res := a.VisitScope(tc.input)
			assertResult(t, res, tc.expectedMessage)
			if tc.expectedMessage == "" && tc.expectedScalar {
				assertEquality(t, true, a.IsScalar)
			}
		})
	}
}

type predicateTestCase struct {
	name            string
	input           *parser.PredicateContext
	a               Analyzer
	expectedMessage string
}

func TestPredicate(t *testing.T) {
	testCases := []predicateTestCase{
		{
			name:            "comparison predicate",
			input:           createPredicateContext("kills > 2"),
			a:               Analyzer{},
			expectedMessage: "",
		},
		{
			name:  "comparison predicate with undeclared identifier",
			input: createPredicateContext("nope > 2"),
			a: Analyzer{
				Identifiers: make(map[string]IdentifierValue),
			},
			expectedMessage: "Undeclared identifer: \"nope\"",
		},
		{
			name:            "logical not predicate",
			input:           createPredicateContext("not assists = 6"),
			a:               Analyzer{},
			expectedMessage: "",
		},
		{
			name:            "parenthesized logical predicate",
			input:           createPredicateContext("(rescues != 2 or damage >= 900) and kills <= 3"),
			a:               Analyzer{},
			expectedMessage: "",
		},
		{
			name:            "logical and predicate",
			input:           createPredicateContext("kills > 2 and damage < 100"),
			a:               Analyzer{},
			expectedMessage: "",
		},
		{
			name:  "logical or same grain aggregate level",
			input: createPredicateContext("kills > 2 or damage < 100"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
					Refinements:        []Dimension{},
				},
			},
			expectedMessage: "",
		},
		{
			name:  "logical or mismatched grain level without game dimension",
			input: createPredicateContext("ben:kills > 2 or damage < 100"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
					Refinements:        []Dimension{},
				},
			},
			expectedMessage: "Both sides of an OR must share the same grain level!",
		},
		{
			name:  "logical or same grain scalar level without game dimension",
			input: createPredicateContext("ben:kills > 2 or isaac:damage < 100"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
					Refinements:        []Dimension{},
				},
			},
			expectedMessage: "",
		},
		{
			name:  "logical or same grain scalar level with g scope without game dimension",
			input: createPredicateContext("ben:kills > 2 or g:damage < 100"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
					Refinements:        []Dimension{},
				},
			},
			expectedMessage: "",
		},
		{
			name:  "logical or mismatched grain level with game dimension permitted",
			input: createPredicateContext("ben:kills > 2 or damage < 100"),
			a: Analyzer{
				Grain: Grain{
					BaseDimension:      TeamDimension,
					BaseDimensionValue: "ib",
					Refinements:        []Dimension{GameDimension},
				},
			},
			expectedMessage: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := tc.a
			res := a.VisitPredicate(tc.input)
			assertResult(t, res, tc.expectedMessage)
		})
	}
}

func TestWhereClause(t *testing.T) {
	t.Run("valid where clause sets current clause", func(t *testing.T) {
		a := Analyzer{}
		res := a.VisitWhereClause(createWhereContext("where kills > 2"))
		assertResult(t, res, "")
	})

	t.Run("where clause propagates predicate error", func(t *testing.T) {
		a := Analyzer{
			Identifiers: make(map[string]IdentifierValue),
		}
		res := a.VisitWhereClause(createWhereContext("where nope > 2"))
		assertResult(t, res, "Undeclared identifer: \"nope\"")
	})
}

func TestSortByClause(t *testing.T) {
	t.Run("valid sort by", func(t *testing.T) {
		a := Analyzer{}
		res := a.VisitSortByClause(createSortByContext("sort by kills desc"))
		assertResult(t, res, "")
	})

	t.Run("valid multiple sort by", func(t *testing.T) {
		a := Analyzer{}
		res := a.VisitSortByClause(createSortByContext("sort by kills desc, damage asc"))
		assertResult(t, res, "")
	})

	t.Run("undeclared identifier in sort by", func(t *testing.T) {
		a := Analyzer{
			Identifiers: make(map[string]IdentifierValue),
		}
		res := a.VisitSortByClause(createSortByContext("sort by nope desc"))
		assertResult(t, res, "Undeclared identifer: \"nope\"")
	})
}

type expressionGrainTestCase struct {
	expected GrainLevel
	input    string
	a        Analyzer
	name     string
}

func TestGetExpressionGrainLevel(t *testing.T) {

	testCases := []expressionGrainTestCase{
		{
			name:     "Simple aggregate",
			input:    "kills + 1",
			expected: GrainAggregate,
		},
		{
			name:     "Simple scope",
			input:    "ben:kills",
			expected: GrainScalar,
		},
		{
			name:     "Scope with literal value on right",
			input:    "ben:kills + 2",
			expected: GrainScalar,
		},
		{
			name:     "Scope with literal value on left",
			input:    "2 + ben:kills",
			expected: GrainScalar,
		},
		{
			name:     "Nested scope",
			input:    "(ben: kills + 1) * 2",
			expected: GrainScalar,
		},
		{
			name:     "Deeply parentheszied scope",
			input:    "((ben:kills))",
			expected: GrainScalar,
		},
		{
			name:     "Aggregate function",
			input:    "average kills",
			expected: GrainAggregate,
		},
		{
			name:  "Coarse grain identifier",
			input: "coarse",
			a: Analyzer{
				Grain:       Grain{BaseDimension: TeamDimension},
				Identifiers: map[string]IdentifierValue{"coarse": {Expr: createExprContext("ben:kills + 30")}},
			},
			expected: GrainAggregate,
		},
		{
			name:  "Fine grain identifier",
			input: "fine",
			a: Analyzer{
				Grain:       Grain{BaseDimension: TeamDimension, Refinements: []Dimension{GameDimension}},
				Identifiers: map[string]IdentifierValue{"fine": {Expr: createExprContext("ben:kills / 2")}},
			},
			expected: GrainScalar,
		},
		{
			name:     "Scope with aggregate function",
			input:    "ben:average damage",
			expected: GrainAggregate,
		},
		{
			name:     "Scope with unscoped measure",
			input:    "ben:damage / damage",
			expected: GrainAggregate,
		},
		{
			name:     "Aggregate scope with  unscoped measure",
			input:    "cody:max damage / damage",
			expected: GrainAggregate,
		},
		{
			name:  "Scope with unscoped measure in game level statement",
			input: "cody:assists * damage",
			a: Analyzer{
				Grain: Grain{BaseDimension: PlayerDimension, Refinements: []Dimension{GameDimension}},
			},
			expected: GrainScalar,
		},
	}

	for _, testCase := range testCases {
		got := testCase.a.getExpressionGrainLevel(createExprContext(testCase.input))
		if got != testCase.expected {
			t.Errorf("Test Failed: %v. Expected %v, got %v", testCase.name, testCase.expected, got)
		}
	}
}

type predicateGrainTestCase struct {
	expected GrainLevel
	input    string
	a        Analyzer
	name     string
}

func TestGetPredicateGrainLevel(t *testing.T) {
	testCases := []predicateGrainTestCase{
		{
			name:     "Aggregate",
			input:    "kills > 2",
			expected: GrainAggregate,
		},
		{
			name:     "Scalar for left-hand side scope comparison",
			input:    "ben:kills > 2",
			expected: GrainScalar,
		},
		{
			name:     "Unscoped and scoped measure",
			input:    "kills > ben:kills",
			expected: GrainAggregate,
		},
		{
			name:     "Aggregate AND",
			input:    "kills > 2 and damage < 100",
			expected: GrainAggregate,
		},
		{
			name:     "Scalar for left-hand side scalar AND",
			input:    "ben:kills > 2 and damage < 100",
			expected: GrainScalar,
		},
		{
			name:     "Scalar for right-hand side scalar AND",
			input:    "kills > 2 and ben:damage < 100",
			expected: GrainScalar,
		},
		{
			name:     "Scalar for left-hand side scalar OR",
			input:    "ben:kills > 2 or damage < 100",
			expected: GrainScalar,
		},
		{
			name:     "Scalar for right-hand side scalar OR",
			input:    "damage < 100 or ben:kills > 2",
			expected: GrainScalar,
		},
		{
			name:     "Scalar for NOT scoped comparison",
			input:    "not (ben:kills > 2)",
			expected: GrainScalar,
		},
		{
			name:  "Idenitifier in game level statement",
			input: "scalarstat > 2",
			a: Analyzer{
				Identifiers: map[string]IdentifierValue{
					"aggstat": {Expr: createExprContext("damage")},
				},
				Grain: Grain{BaseDimension: TeamDimension, Refinements: []Dimension{GameDimension}},
			},
			expected: GrainScalar,
		},
		{
			name:  "Identifier in aggregate statement",
			input: "aggstat > 2",
			a: Analyzer{
				Identifiers: map[string]IdentifierValue{
					"aggstat": {Expr: createExprContext("damage")},
				},
			},
			expected: GrainAggregate,
		},
	}

	for _, testCase := range testCases {
		got := testCase.a.getPredicateGrainLevel(createPredicateContext(testCase.input))
		if got != testCase.expected {
			t.Errorf("Test Failed: %v. Expected %v, got %v", testCase.name, testCase.expected, got)
		}
	}
}

type statementTestCase struct {
	name            string
	input           string
	expectedMessage string
}

func TestAnalyzeStatements(t *testing.T) {
	testCases := []statementTestCase{
		{
			name: "valid full query",
			input: `show kills
for team ic
by player, game
where kills > 2
sort by kills desc`,
			expectedMessage: "",
		},
		{
			name: "valid query with declared identifier and scope",
			input: `show (5 + 5) yep, yep, ben:kills
for team ib
where kills > 2`,
			expectedMessage: "",
		},
		{
			name: "invalid query duplicate identifier in show",
			input: `show (5 + 5) yep, (2 - 3) yep
for team it`,
			expectedMessage: "Identifier \"yep\" already used!",
		},
		{
			name: "invalid query undeclared identifier in where",
			input: `show kills
for team ic
where nope > 5`,
			expectedMessage: "Undeclared identifer: \"nope\"",
		},
		{
			name: "invalid query player dimension mismatch",
			input: `show kills
for player p`,
			expectedMessage: "Unrecognized player p. Expected one of",
		},
		{
			name: "invalid query aggregate in scalar",
			input: `show average kills
for player ben
by game`,
			expectedMessage: "Cannot use aggregate functions in a scalar statement!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stmt := createStatementContext(tc.input)
			a, err := Analyze(stmt)
			if tc.expectedMessage == "" {
				if err != nil {
					t.Fatalf("Expected no error, got %v", err)
				}
				if a.Grain.BaseDimension == UnkownDimension {
					t.Errorf("Expected valid BaseDimension, got UnkownDimension")
				}
			} else {
				if err == nil {
					t.Fatalf("Expected error containing %q, got nil", tc.expectedMessage)
				}
				if !strings.Contains(err.Error(), tc.expectedMessage) {
					t.Fatalf("Expected error containing %q, got %q", tc.expectedMessage, err.Error())
				}
			}
		})
	}
}
