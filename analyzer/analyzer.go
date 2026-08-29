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
	if ctx.ByClause() != nil {
		byResult := a.Visit(ctx.ByClause()).(result)
		if byResult.err != nil {
			return byResult
		}
	}
	return result{}
}

func (a *Analyzer) VisitForClause(ctx *parser.ForClauseContext) any {
	dimension := ctx.Dimension()
	identifier := strings.ToLower(ctx.IDENTIFIER().GetText())
	if dimension.PLAYER() != nil {
		a.forDimensionType = PlayerDimension
		player := identifier
		if slices.Contains(players, player) {
			return ok()
		}
		return fail(fmt.Sprintf("Unrecognized player %v. Expected one of %v", player, players), ctx.IDENTIFIER().GetSymbol())
	}
	if dimension.TEAM() != nil {
		a.forDimensionType = TeamDimension
		team := identifier
		if slices.Contains(teams, team) {
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

	var prevDimension Dimension
	for i, dimension := range ctx.AllDimension() {
		if dimension.PLAYER() != nil {
			player := dimension.PLAYER()
			if prevDimension == PlayerDimension {
				return fail("Used the same dimension 'player' more than once!", player.GetSymbol())
			}
			if a.forDimensionType == PlayerDimension {
				return fail("'player' dimension already used in FOR clause!", player.GetSymbol())
			}
			isPlayerUsed = true
			prevDimension = PlayerDimension
			continue
		}
		if dimension.TEAM() != nil {
			team := dimension.TEAM()
			if prevDimension == TeamDimension {
				return fail("Used the same dimension 'team' more than once!", team.GetSymbol())
			}
			if a.forDimensionType == TeamDimension {
				return fail("'team' dimension already used in FOR clause!", team.GetSymbol())
			}
			prevDimension = TeamDimension
			continue
		}

		invalidDimension := ctx.Dimension(i).GetStart()
		return fail(fmt.Sprintf("Invalid dimension used in the BY clause! Got %v. Expected one of 'team' or 'player'", invalidDimension.GetText()), invalidDimension)
	}

	a.isScalar = a.forDimensionType == PlayerDimension && isGameUsed || isGameUsed && isPlayerUsed
	return ok()
}
