package analyzer

import (
	"fmt"
	"pubql/parser"
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

type identifierValue struct {
	expr parser.IExprContext
	pred parser.IPredicateContext
}

type Grain struct {
	baseDimension      Dimension
	baseDimensionValue string
	refinements        []Dimension
}

type projection struct {
	measure     string
	aggregate   string
	expression  parser.IExprContext
	predicate   parser.IPredicateContext
	specificity parser.ISpecificityContext
}

type Analyzer struct {
	parser.BasepgqlVisitor
	Grain
	Projections   []projection
	IsScalar      bool
	IsSequential  bool
	Identifiers   map[string]identifierValue
	currentClause Clause
}

func Analyze(tree antlr.ParseTree) (Analyzer, error) {
	analyzer := Analyzer{
		Identifiers: make(map[string]identifierValue),
		Grain:       Grain{refinements: make([]Dimension, 0)},
		Projections: make([]projection, 0),
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
		a.Grain.baseDimension = PlayerDimension
		player := identifier
		if slices.Contains(players, player) {
			a.Grain.baseDimensionValue = player
			return ok()
		}
		return fail(fmt.Sprintf("Unrecognized player %v. Expected one of %v", player, players), ctx.IDENTIFIER().GetSymbol())
	}
	if dimension.TEAM() != nil {
		a.Grain.baseDimension = TeamDimension
		team := identifier
		if slices.Contains(teams, team) {
			a.Grain.baseDimensionValue = team
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
		a.Grain.refinements = append(a.Grain.refinements, GameDimension)
	}

	var prevDimension Dimension
	for i, dimension := range ctx.AllDimension() {
		if dimension.PLAYER() != nil {
			player := dimension.PLAYER()
			if prevDimension == PlayerDimension {
				return fail("Used the same dimension 'player' more than once!", player.GetSymbol())
			}
			if a.Grain.baseDimension == PlayerDimension {
				return fail("'player' dimension already used in FOR clause!", player.GetSymbol())
			}
			isPlayerUsed = true
			prevDimension = PlayerDimension
			a.Grain.refinements = append(a.Grain.refinements, PlayerDimension)
			continue
		}
		if dimension.TEAM() != nil {
			team := dimension.TEAM()
			if prevDimension == TeamDimension {
				return fail("Used the same dimension 'team' more than once!", team.GetSymbol())
			}
			if a.Grain.baseDimension == TeamDimension {
				return fail("'team' dimension already used in FOR clause!", team.GetSymbol())
			}
			prevDimension = TeamDimension
			a.Grain.refinements = append(a.Grain.refinements, TeamDimension)
			continue
		}

		invalidDimension := ctx.Dimension(i).GetStart()
		return fail(fmt.Sprintf("Invalid dimension used in the BY clause! Got %v. Expected one of 'team' or 'player'", invalidDimension.GetText()), invalidDimension)
	}

	a.IsScalar = a.Grain.baseDimension == PlayerDimension && isGameUsed || isGameUsed && isPlayerUsed
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
	projection := projection{}
	if ctx.Specificity() != nil {
		result := a.Visit(ctx.Specificity()).(result)
		if result.err != nil {
			return result
		}
		projection.specificity = ctx.Specificity()
	}
	if ctx.Measure() != nil {
		projection.measure = ctx.Measure().GetText()
	}
	if ctx.AggregateFunction() != nil {
		if a.IsScalar {
			return fail("Cannot use aggregate functions in a scalar statement!", ctx.AggregateFunction().GetStart())
		}
		projection.aggregate = ctx.GetText()
	}
	if ctx.IDENTIFIER() != nil {
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
			a.Identifiers[identifier] = identifierValue{expr: ctx.Expr()}
			projection.expression = ctx.Expr()
		} else if ctx.Predicate() != nil {
			result := a.Visit(ctx.Predicate()).(result)
			if result.err != nil {
				return result
			}
			a.Identifiers[identifier] = identifierValue{pred: ctx.Predicate()}
			projection.predicate = ctx.Predicate()
		} else {
			// This should never happen since the parser should catch this. Just being safe
			return fail(fmt.Sprintf("Could not find expression or predicate preceding identifier \"%q\"", identifier), ctx.IDENTIFIER().GetSymbol())
		}
	}

	a.Projections = append(a.Projections, projection)
	return ok()
}

func (a *Analyzer) VisitWhereClause(ctx *parser.WhereClauseContext) any {
	for _, predicate := range ctx.AllPredicate() {
		result := a.Visit(predicate).(result)
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

func (a *Analyzer) VisitPredicate(ctx *parser.PredicateContext) any {
	for _, expr := range ctx.AllExpr() {
		result := a.Visit(expr).(result)
		if result.err != nil {
			return result
		}
	}

	for _, predicate := range ctx.AllPredicate() {
		result := a.Visit(predicate).(result)
		if result.err != nil {
			return result
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
	if ctx.Specificity() != nil {
		result := a.Visit(ctx.Specificity()).(result)
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

func (a *Analyzer) VisitSpecificity(ctx *parser.SpecificityContext) any {
	identifier := strings.ToLower(ctx.IDENTIFIER().GetText())
	identifierToken := ctx.IDENTIFIER().GetSymbol()

	switch a.Grain.baseDimension {
	case PlayerDimension:
		if identifier != "g" {
			return fail(fmt.Sprintf("Invalid specificity type used in a statement with \"FOR player\". Got %v. Expected \"g\"", identifier), identifierToken)
		}
		if a.currentClause == ShowClause {
			return fail("Specificity cannot be used in the SHOW clause of a statement with \"FOR player\"", ctx.GetStart())
		}
		if ctx.AggregateFunction() != nil {
			return fail("Cannot use aggregates in a specificity in a statement with \"FOR player\"!", ctx.AggregateFunction().GetStart())
		}
	case TeamDimension:
		if identifier != "g" && !slices.Contains(players, identifier) {
			message := fmt.Sprintf("Identifier in specificity cannot be mapped to a player! Got: %v. Expected one of %v", identifier, players)
			return fail(message, identifierToken)
		}
		if ctx.AggregateFunction() != nil && identifier != "g" && a.IsScalar {
			return fail("Cannot use an aggregate function within a scalar statement outside of the \"g:\" specificity!", ctx.AggregateFunction().GetStart())
		}
		team := teamPlayers[a.Grain.baseDimensionValue]
		if !slices.Contains(team, identifier) {
			return fail(fmt.Sprintf("Player \"%v\" is not on Team \"%v\"!", identifier, a.Grain.baseDimensionValue), identifierToken)
		}

		if a.currentClause == ShowClause && slices.Contains(a.Grain.refinements, GameDimension) {
			a.IsScalar = true
		}
	default:
		return fail("Fatal error! Somehow there is no dimension in the FOR!", ctx.GetStart())
	}

	return ok()
}
