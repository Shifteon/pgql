package analyzer

import (
	"pubql/parser"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

type testCase struct {
	name            string
	input           string
	expectedMessage string
}

type testCases []testCase

func buildTree(input string) parser.IStatementContext {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewpgqlLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewpgqlParser(stream)
	return parser.Statement()
}

func runTests(t *testing.T, testCases testCases) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := Analyze(buildTree(test.input))

			// no error expected
			if test.expectedMessage == "" {
				if err != nil {
					t.Fatalf("Expected no error. Got %q", err)
				}
				return
			}

			// we expect an error at this point
			if err == nil {
				t.Fatalf("Expected error with message like: %q. Got \"nil\"", test.expectedMessage)
			}

			if !strings.Contains(err.Error(), test.expectedMessage) {
				t.Fatalf("Expected error with message like %q. Got %q", test.expectedMessage, err.Error())
			}

		})
	}
}

func TestForClause(t *testing.T) {
	validFor := `show kills
for team ic`
	invalidFor := `show kills
for win ic`
	invalidTeam := `show kills
for team i`
	invalidPlayer := `show kills
for player p`

	tests := testCases{
		{"valid", validFor, ""},
		{"invalid dimension", invalidFor, "Invalid dimension used in the FOR clause!"},
		{"invalid team", invalidTeam, "Unrecognized team i. Expected one of"},
		{"invalid player", invalidPlayer, "Unrecognized player p. Expected one of"},
	}

	runTests(t, tests)
}

func TestByClause(t *testing.T) {
	validBy := `show kills
for team ic
by player`
	validByTwoDimensions := `show kills
for team ic
by player, game`
	invalidDimension := `show kills
for team ic
by date`
	teamAlreadyInFor := `show kills
for team ic
by team`
	playerAlreadyInFor := `show kills
for player ben
by player`
	gameUsedTwice := `show kills
for team ic
by game, game`
	teamUsedTwice := `show kills
for player ben
by team, team`
	playerUsedTwice := `show kills
for team ic
by player, player`

	tests := testCases{
		{"valid", validBy, ""},
		{"valid with two dimensions", validByTwoDimensions, ""},
		{"invalid dimenstion", invalidDimension, "Invalid dimension used in the BY clause! Got date. Expected one of 'team' or 'player'"},
		{"team already in for", teamAlreadyInFor, "'team' dimension already used in FOR clause!"},
		{"player already in for", playerAlreadyInFor, "'player' dimension already used in FOR clause!"},
		{"game used twice", gameUsedTwice, "Used the same dimension 'game' more than once!"},
		{"team used twice", teamUsedTwice, "Used the same dimension 'team' more than once!"},
		{"player used twice", playerUsedTwice, "Used the same dimension 'player' more than once!"},
	}

	runTests(t, tests)
}

func TestShowClause(t *testing.T) {
	validShow := `show kills
for team ic`
	aggregateInScalar := `show average kills
for player ben
by game`
	aggregateInAggregate := `show max kills
for player ben`
	identifiers := `show (5 + 5) yep
for player ben`

	tests := testCases{
		{"valid show", validShow, ""},
		{"aggregate in scalar", aggregateInScalar, "Cannot use aggregate functions in a scalar statement!"},
		{"aggregate in aggregate", aggregateInAggregate, ""},
		// TODO: How can I test identifiers are set correctly?
		{"identifiers", identifiers, ""},
	}

	runTests(t, tests)
}
