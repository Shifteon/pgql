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
	DimensionUnkown Dimension = iota
	Team
	Player
)

// all of the valid team options
var teams = []string{"ictb", "itb", "icb", "ctb", "ict", "ib", "ic", "it"}

// all of the valid players
var players = []string{"isaac", "cody", "trenton", "ben"}

type Analyzer struct {
	parser.BasepgqlVisitor
	isScalar         bool
	forDimensionType Dimension
	identifiers      map[string]string
}

func Analyze(tree antlr.ParseTree) error {
	analyzer := Analyzer{}
	result := analyzer.Visit(tree).(result)
	return result.err
}

func (a *Analyzer) Visit(tree antlr.ParseTree) any {
	return tree.Accept(a)
}

func (a *Analyzer) VisitStatement(ctx *parser.StatementContext) any {
	forResult := a.Visit(ctx.ForClause()).(result)
	if forResult.err != nil {
		return forResult
	}
	return result{}
}

func (a *Analyzer) VisitForClause(ctx *parser.ForClauseContext) any {
	dimension := ctx.Dimension()
	identifier := strings.ToLower(ctx.IDENTIFIER().GetText())
	if dimension.PLAYER() != nil {
		a.forDimensionType = Player
		player := identifier
		if slices.Contains(players, player) {
			return ok()
		}
		lineNumber := ctx.IDENTIFIER().GetSymbol().GetLine()
		charPosition := ctx.IDENTIFIER().GetSymbol().GetTokenSource().GetCharPositionInLine()
		return fail(fmt.Sprintf("Unrecognized player %v. Expected one of %v", player, players), lineNumber, charPosition)
	}
	if dimension.TEAM() != nil {
		a.forDimensionType = Team
		team := identifier
		if slices.Contains(teams, team) {
			return ok()
		}
		lineNumber := ctx.IDENTIFIER().GetSymbol().GetLine()
		charPosition := ctx.IDENTIFIER().GetSymbol().GetTokenSource().GetCharPositionInLine()
		return fail(fmt.Sprintf("Unrecognized team %v. Expected one of %v", team, teams), lineNumber, charPosition)
	}

	lineNumber := dimension.GetStart().GetLine()
	charPosition := dimension.GetStart().GetTokenSource().GetCharPositionInLine()
	return fail(fmt.Sprintf("Invalid dimension used in the FOR clause! Got %v. Expected one of team or player", dimension.GetText()), lineNumber, charPosition)
}
