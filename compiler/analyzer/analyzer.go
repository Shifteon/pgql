package analyzer

import (
	"fmt"
	"pubql/compiler/internal"
	"pubql/compiler/parser"
	"slices"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Dimension int

const (
	UnkownDimension Dimension = iota
	TeamDimension
	PlayerDimension
	GameDimension
)

type Clause int

const (
	UnkownClause Clause = iota
	ShowClause
	ForClause
	ByClause
	WhereClause
	SortByClause
)

type GrainLevel int

const (
	UnkownGrain GrainLevel = iota
	GrainAggregate
	GrainScalar
)

type IdenitifierType int

const (
	UnkownIdentifier IdenitifierType = iota
	IdenitifierExpr
	IdenitifierPred
)

// all of the valid team options
var teams = []string{"ictb", "itb", "icb", "ctb", "ict", "ib", "ic", "it"}

// all of the valid players
var players = []string{"isaac", "cody", "trenton", "ben"}

var teamPlayers = map[string][]string{
	"ictb": {"isaac", "cody", "trenton", "ben"},
	"itb":  {"isaac", "trenton", "ben"},
	"icb":  {"isaac", "cody", "ben"},
	"ctb":  {"cody", "trenton", "ben"},
	"ict":  {"isaac", "cody", "trenton"},
	"ib":   {"isaac", "ben"},
	"ic":   {"isaac", "cody"},
	"it":   {"isaac", "trenton"},
}

type Expression struct {
	ExprCtx    *parser.ExprContext
	GrainLevel GrainLevel
}

type Predicate struct {
	PredCtx    *parser.PredicateContext
	GrainLevel GrainLevel
}

type IdentifierValue struct {
	Type IdenitifierType
	Expr Expression
	Pred Predicate
}

type Grain struct {
	BaseDimension      Dimension
	BaseDimensionValue string
	Refinements        []Dimension
}

type Projection struct {
	Measure    string
	Aggregate  string
	Identifier string
	Expression Expression
	Predicate  Predicate
	Scope      parser.IScopeContext
}

type Analyzer struct {
	parser.BasepgqlVisitor
	Grain
	Projections   []Projection
	IsScalar      bool
	IsSequential  bool
	Identifiers   map[string]IdentifierValue
	currentClause Clause
}

func Analyze(tree antlr.ParseTree) (Analyzer, error) {
	analyzer := Analyzer{
		Identifiers: make(map[string]IdentifierValue),
		Grain:       Grain{Refinements: make([]Dimension, 0)},
		Projections: make([]Projection, 0),
	}
	result := analyzer.Visit(tree).(result)
	return analyzer, result.err
}

func (a *Analyzer) Visit(tree antlr.ParseTree) any {
	return tree.Accept(a)
}

func (a *Analyzer) VisitStatement(ctx *parser.StatementContext) any {
	a.currentClause = ForClause
	forResult := a.Visit(ctx.ForClause()).(result)
	if forResult.err != nil {
		return forResult
	}

	if ctx.ByClause() != nil {
		a.currentClause = ByClause
		byResult := a.Visit(ctx.ByClause()).(result)
		if byResult.err != nil {
			return byResult
		}
	}

	a.currentClause = ShowClause
	showResult := a.Visit(ctx.ShowClause()).(result)
	if showResult.err != nil {
		return showResult
	}

	if ctx.WhereClause() != nil {
		a.currentClause = WhereClause
		whereResult := a.Visit(ctx.WhereClause()).(result)
		if whereResult.err != nil {
			return whereResult
		}
	}

	if ctx.SortByClause() != nil {
		a.currentClause = SortByClause
		sortByResult := a.Visit(ctx.SortByClause()).(result)
		if sortByResult.err != nil {
			return sortByResult
		}
	}

	return ok()
}

func (a *Analyzer) VisitForClause(ctx *parser.ForClauseContext) any {
	dimension := ctx.Dimension()
	identifier := strings.ToLower(ctx.IDENTIFIER().GetText())
	if dimension.PLAYER() != nil {
		a.Grain.BaseDimension = PlayerDimension
		player := identifier
		if slices.Contains(players, player) {
			a.Grain.BaseDimensionValue = player
			return ok()
		}
		return fail(fmt.Sprintf("Unrecognized player %v. Expected one of %v", player, players), ctx.IDENTIFIER().GetSymbol())
	}
	if dimension.TEAM() != nil {
		a.Grain.BaseDimension = TeamDimension
		team := identifier
		if slices.Contains(teams, team) {
			a.Grain.BaseDimensionValue = team
			return ok()
		}
		return fail(fmt.Sprintf("Unrecognized team %v. Expected one of %v", team, teams), ctx.IDENTIFIER().GetSymbol())
	}

	return fail(fmt.Sprintf("Invalid dimension used in the FOR clause! Got %v. Expected one of team or player", dimension.GetText()), dimension.GetStart())
}

func (a *Analyzer) VisitByClause(ctx *parser.ByClauseContext) any {
	if len(ctx.AllGAME()) > 1 {
		game := ctx.AllGAME()[1]
		return fail("Used the same dimension 'game' more than once!", game.GetSymbol())
	}

	isPlayerUsed := false
	isGameUsed := len(ctx.AllGAME()) == 1

	if isGameUsed {
		a.Grain.Refinements = append(a.Grain.Refinements, GameDimension)
	}

	var prevDimension Dimension
	for i, dimension := range ctx.AllDimension() {
		if dimension.PLAYER() != nil {
			player := dimension.PLAYER()
			if prevDimension == PlayerDimension {
				return fail("Used the same dimension 'player' more than once!", player.GetSymbol())
			}
			if a.Grain.BaseDimension == PlayerDimension {
				return fail("'player' dimension already used in FOR clause!", player.GetSymbol())
			}
			isPlayerUsed = true
			prevDimension = PlayerDimension
			a.Grain.Refinements = append(a.Grain.Refinements, PlayerDimension)
			continue
		}
		if dimension.TEAM() != nil {
			team := dimension.TEAM()
			if prevDimension == TeamDimension {
				return fail("Used the same dimension 'team' more than once!", team.GetSymbol())
			}
			if a.Grain.BaseDimension == TeamDimension {
				return fail("'team' dimension already used in FOR clause!", team.GetSymbol())
			}
			prevDimension = TeamDimension
			a.Grain.Refinements = append(a.Grain.Refinements, TeamDimension)
			continue
		}

		invalidDimension := ctx.Dimension(i).GetStart()
		return fail(fmt.Sprintf("Invalid dimension used in the BY clause! Got %v. Expected one of 'team' or 'player'", invalidDimension.GetText()), invalidDimension)
	}

	a.IsScalar = a.Grain.BaseDimension == PlayerDimension && isGameUsed || isGameUsed && isPlayerUsed
	return ok()
}

func (a *Analyzer) VisitShowClause(ctx *parser.ShowClauseContext) any {
	if ctx.SequentialFunction() != nil {
		a.IsSequential = true
	}
	result := a.Visit(ctx.ShowBody()).(result)
	if result.err != nil {
		return result
	}

	return ok()
}

func (a *Analyzer) VisitShowBody(ctx *parser.ShowBodyContext) any {
	for _, fragment := range ctx.AllShowFragment() {
		result := a.Visit(fragment).(result)
		if result.err != nil {
			return result
		}
	}

	return ok()
}

func (a *Analyzer) VisitShowFragment(ctx *parser.ShowFragmentContext) any {
	projection := Projection{}

	// TODO: Should I consolidate this logic?
	var grainLevel GrainLevel
	if slices.Contains(a.Grain.Refinements, GameDimension) {
		grainLevel = GrainScalar
	} else {
		grainLevel = GrainAggregate
	}

	if ctx.Scope() != nil {
		result := a.Visit(ctx.Scope()).(result)
		if result.err != nil {
			return result
		}
		projection.Scope = ctx.Scope()
	}
	if ctx.Measure() != nil {
		projection.Measure = ctx.Measure().GetText()
	}
	if ctx.AggregateFunction() != nil {
		if a.IsScalar {
			return fail("Cannot use aggregate functions in a scalar statement!", ctx.AggregateFunction().GetStart())
		}
		aggr := ctx.AggregateFunction()
		if aggr.AVERAGE() != nil {
			projection.Aggregate = aggr.AVERAGE().GetText()
		} else if aggr.TOTAL() != nil {
			projection.Aggregate = aggr.TOTAL().GetText()
		} else if aggr.MAX() != nil {
			projection.Aggregate = aggr.MAX().GetText()
		} else if aggr.MIN() != nil {
			projection.Aggregate = aggr.MIN().GetText()
		}
		projection.Measure = aggr.Measure().GetText()
	}

	if ctx.Predicate() != nil || ctx.Expr() != nil {
		identifier := ctx.IDENTIFIER().GetText()
		_, alreadyUsed := a.Identifiers[identifier]
		if alreadyUsed {
			return fail(fmt.Sprintf("Identifier \"%v\" already used!", identifier), ctx.IDENTIFIER().GetSymbol())
		}
		if ctx.Expr() != nil {
			result := a.Visit(ctx.Expr()).(result)
			if result.err != nil {
				return result
			}
			expr := Expression{ExprCtx: ctx.Expr().(*parser.ExprContext), GrainLevel: grainLevel}
			a.Identifiers[identifier] = IdentifierValue{Expr: expr, Type: IdenitifierExpr}
			projection.Expression = expr
		} else if ctx.Predicate() != nil {
			result := a.Visit(ctx.Predicate()).(result)
			if result.err != nil {
				return result
			}
			pred := Predicate{PredCtx: ctx.Predicate().(*parser.PredicateContext), GrainLevel: grainLevel}
			a.Identifiers[identifier] = IdentifierValue{Pred: pred, Type: IdenitifierPred}
			projection.Predicate = pred
		} else {
			// This should never happen since the parser should catch this. Just being safe
			return fail(fmt.Sprintf("Could not find expression or predicate preceding identifier \"%q\"", identifier), ctx.IDENTIFIER().GetSymbol())
		}
		projection.Identifier = identifier
	} else if ctx.IDENTIFIER() != nil {
		identifier := ctx.IDENTIFIER().GetText()
		value, exists := a.Identifiers[identifier]

		if !exists {
			// TODO: Probably consolidate error messages
			return fail(fmt.Sprintf("Undeclared identifer: \"%v\"", identifier), ctx.IDENTIFIER().GetSymbol())
		}

		switch value.Type {
		case IdenitifierExpr:
			projection.Expression = value.Expr
		case IdenitifierPred:
			projection.Predicate = value.Pred
		default:
			// This should not happen. Major bug if so
			return fail(fmt.Sprintf("Fatal! Somehow identifier exists but does not map to an expression or a predicate! \"%v\"", identifier), ctx.IDENTIFIER().GetSymbol())
		}
	}

	a.Projections = append(a.Projections, projection)
	return ok()
}

func (a *Analyzer) VisitWhereClause(ctx *parser.WhereClauseContext) any {
	if ctx.Predicate() != nil {
		result := a.Visit(ctx.Predicate()).(result)
		if result.err != nil {
			return result
		}
	}

	return ok()
}

func (a *Analyzer) VisitSortByClause(ctx *parser.SortByClauseContext) any {
	result := a.Visit(ctx.SortBody()).(result)
	if result.err != nil {
		return result
	}
	return ok()
}

func (a *Analyzer) VisitSortBody(ctx *parser.SortBodyContext) any {
	exprResult := a.Visit(ctx.Expr()).(result)
	if exprResult.err != nil {
		return exprResult
	}

	for _, sortBody := range ctx.AllSortBody() {
		result := a.Visit(sortBody).(result)
		if result.err != nil {
			return result
		}
	}

	return ok()
}

func (a *Analyzer) GetExpressionGrainLevel(expr *parser.ExprContext) GrainLevel {
	// We consider it scalar if we are at the game level
	if slices.Contains(a.Grain.Refinements, GameDimension) {
		return GrainScalar
	}

	grainLevel := GrainAggregate

	if expr.LPAREN() != nil {
		grainLevel = a.GetExpressionGrainLevel(expr.Expr(0).(*parser.ExprContext))
	} else if expr.IDENTIFIER() != nil {
		// Identifiers inherit the statement grain level no matter what since they are defined in the SHOW
		e := a.Identifiers[expr.IDENTIFIER().GetText()]
		if e.Type != UnkownIdentifier {
			// we can only get to this point if the statement is not at the game level so it must be aggregate
			grainLevel = GrainAggregate
		}
	} else if operator := internal.CheckAndReturnOperator(expr); operator != internal.UnknownOperator {
		var leftGrainLevel, rightGrainLevel GrainLevel
		leftExpr := expr.Expr(0)
		rightExpr := expr.Expr(1)

		leftGrainLevel = a.GetExpressionGrainLevel(leftExpr.(*parser.ExprContext))
		rightGrainLevel = a.GetExpressionGrainLevel(rightExpr.(*parser.ExprContext))

		// The rule is that if any side of an expression contains an unscoped measure
		// then the whole expression is evaluated at the statement grain
		if leftGrainLevel == GrainAggregate || rightGrainLevel == GrainAggregate {
			grainLevel = GrainAggregate
		} else {
			grainLevel = GrainScalar
		}
	} else { // leafs
		if expr.Scope() != nil {
			grainLevel = GrainScalar
			if expr.Scope().AggregateFunction() != nil {
				grainLevel = GrainAggregate
			}
		} else if expr.Measure() != nil || expr.AggregateFunction() != nil {
			grainLevel = GrainAggregate
		} else if expr.NUMBER() != nil {
			// we set this to grain scalar so the other side can override it if it is aggregate
			grainLevel = GrainScalar
		}
	}

	return grainLevel
}

func (a *Analyzer) GetPredicateGrainLevel(ctx *parser.PredicateContext) GrainLevel {
	grainLevel := GrainAggregate

	if slices.Contains(a.Grain.Refinements, GameDimension) {
		return GrainScalar
	}

	if ctx.ComparisonOperator() != nil {
		leftGrainLevel := a.GetExpressionGrainLevel(ctx.Expr(0).(*parser.ExprContext))
		rightGrainLevel := a.GetExpressionGrainLevel(ctx.Expr(1).(*parser.ExprContext))

		// If either side is aggregate the whole thing is aggregate
		if leftGrainLevel == GrainAggregate || rightGrainLevel == GrainAggregate {
			grainLevel = GrainAggregate
		} else {
			grainLevel = GrainScalar
		}
	} else if ctx.LOGICALNOT() != nil || ctx.LPAREN() != nil {
		grainLevel = a.GetPredicateGrainLevel(ctx.Predicate(0).(*parser.PredicateContext))
	} else if ctx.LOGICALOR() != nil || ctx.LOGICALAND() != nil {
		leftGrainLevel := a.GetPredicateGrainLevel(ctx.Predicate(0).(*parser.PredicateContext))
		rightGrainLevel := a.GetPredicateGrainLevel(ctx.Predicate(1).(*parser.PredicateContext))
		if leftGrainLevel == GrainAggregate || rightGrainLevel == GrainAggregate {
			grainLevel = GrainAggregate
		}
	}

	return grainLevel
}

func (a *Analyzer) VisitPredicate(ctx *parser.PredicateContext) any {
	if ctx.LPAREN() != nil || ctx.LOGICALNOT() != nil {
		result := a.Visit(ctx.Predicate(0)).(result)

		if result.err != nil {
			return result
		}
	} else if ctx.LOGICALAND() != nil || ctx.LOGICALOR() != nil {
		leftResult := a.Visit(ctx.Predicate(0)).(result)
		rightResult := a.Visit(ctx.Predicate(1)).(result)

		if leftResult.err != nil {
			return leftResult
		}
		if rightResult.err != nil {
			return rightResult
		}

		// both sides of an OR must have the same grain level
		if ctx.LOGICALOR() != nil && !slices.Contains(a.Grain.Refinements, GameDimension) {
			leftGrainLevel := a.GetPredicateGrainLevel(ctx.Predicate(0).(*parser.PredicateContext))
			rightGrainLevel := a.GetPredicateGrainLevel(ctx.Predicate(1).(*parser.PredicateContext))
			if leftGrainLevel != rightGrainLevel {
				// TODO: Better message
				return fail("Both sides of an OR must share the same grain level!", ctx.GetStart())
			}
		}
	} else if ctx.ComparisonOperator() != nil {
		leftResult := a.Visit(ctx.Expr(0)).(result)
		rightResult := a.Visit(ctx.Expr(1)).(result)

		if leftResult.err != nil {
			return leftResult
		}
		if rightResult.err != nil {
			return rightResult
		}
	}

	return ok()
}

func (a *Analyzer) VisitExpr(ctx *parser.ExprContext) any {
	if ctx.IDENTIFIER() != nil {
		_, exists := a.Identifiers[ctx.IDENTIFIER().GetText()]
		if !exists {
			return fail(fmt.Sprintf("Undeclared identifer: \"%v\"", ctx.IDENTIFIER().GetText()), ctx.IDENTIFIER().GetSymbol())
		}
	}
	if a.IsScalar && ctx.AggregateFunction() != nil {
		return fail("Cannot use aggregate functions in a scalar statement!", ctx.AggregateFunction().GetStart())
	}
	if ctx.Scope() != nil {
		result := a.Visit(ctx.Scope()).(result)
		if result.err != nil {
			return result
		}
	}

	if len(ctx.AllExpr()) > 0 {
		for _, expr := range ctx.AllExpr() {
			result := a.Visit(expr).(result)
			if result.err != nil {
				return result
			}
		}
	}

	return ok()
}

func (a *Analyzer) VisitScope(ctx *parser.ScopeContext) any {
	identifier := strings.ToLower(ctx.IDENTIFIER().GetText())
	identifierToken := ctx.IDENTIFIER().GetSymbol()

	switch a.Grain.BaseDimension {
	case PlayerDimension:
		if identifier != "g" {
			return fail(fmt.Sprintf("Invalid scope type used in a statement with \"FOR player\". Got %v. Expected \"g\"", identifier), identifierToken)
		}
		if a.currentClause == ShowClause {
			return fail("Scope cannot be used in the SHOW clause of a statement with \"FOR player\"", ctx.GetStart())
		}
		if ctx.AggregateFunction() != nil {
			return fail("Cannot use aggregates in a scope in a statement with \"FOR player\"!", ctx.AggregateFunction().GetStart())
		}
	case TeamDimension:
		if identifier != "g" && !slices.Contains(players, identifier) {
			message := fmt.Sprintf("Identifier in scope cannot be mapped to a player! Got: %v. Expected one of %v", identifier, players)
			return fail(message, identifierToken)
		}
		if ctx.AggregateFunction() != nil && identifier != "g" && a.IsScalar {
			return fail("Cannot use an aggregate function within a scalar statement outside of the \"g:\" scope!", ctx.AggregateFunction().GetStart())
		}
		team := teamPlayers[a.Grain.BaseDimensionValue]
		if identifier != "g" && !slices.Contains(team, identifier) {
			return fail(fmt.Sprintf("Player \"%v\" is not on Team \"%v\"!", identifier, a.Grain.BaseDimensionValue), identifierToken)
		}

		if a.currentClause == ShowClause && slices.Contains(a.Grain.Refinements, GameDimension) {
			a.IsScalar = true
		}
	default:
		return fail("Fatal error! Somehow there is no dimension in the FOR!", ctx.GetStart())
	}

	return ok()
}
