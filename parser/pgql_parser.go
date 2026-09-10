// Code generated from pgql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // pgql

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type pgqlParser struct {
	*antlr.BaseParser
}

var PgqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pgqlParserInit() {
	staticData := &PgqlParserStaticData
	staticData.LiteralNames = []string{
		"", "'show'", "'for'", "'by'", "'where'", "'sort by'", "':'", "'player'",
		"'team'", "'date'", "'type'", "'kills'", "'damage'", "'assists'", "'rescues'",
		"'recalls'", "'win'", "'game'", "'totalWins'", "'totalLoses'", "'gameCount'",
		"'desc'", "'asc'", "'running'", "'average'", "'total'", "'min'", "'max'",
		"'<='", "'>='", "'>'", "'<'", "'='", "'!='", "'+'", "'-'", "'*'", "'/'",
		"'and'", "'or'", "'not'", "'('", "')'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "PLAYER", "TEAM", "DATE", "TYPE", "KILLS",
		"DAMAGE", "ASSISTS", "RESCUES", "RECALLS", "WIN", "GAME", "TOTALWINS",
		"TOTALLOSES", "GAMECOUNT", "DESC", "ASC", "RUNNING", "AVERAGE", "TOTAL",
		"MIN", "MAX", "LESSEREQUAL", "GREATEREQUAL", "GREATER", "LESSER", "EQUAL",
		"NOTEQUAL", "SUM", "DIFFERENCE", "MULTIPLY", "DIVIDE", "LOGICALAND",
		"LOGICALOR", "LOGICALNOT", "LPAREN", "RPAREN", "COMMA", "STRING", "NUMBER",
		"IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"statement", "showClause", "showBody", "showFragment", "forClause",
		"byClause", "whereClause", "sortByClause", "sortBody", "dimension",
		"measure", "aggregateFunction", "sequentialFunction", "comparisonOperator",
		"predicate", "scope", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 176, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 3, 0, 38, 8, 0, 1, 0, 3, 0, 41, 8, 0, 1,
		0, 3, 0, 44, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 50, 8, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 5, 2, 57, 8, 2, 10, 2, 12, 2, 60, 9, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 68, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 73, 8, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 82, 8, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 87, 8, 5, 3, 5, 89, 8, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 3, 8, 99, 8, 8, 1, 8, 1, 8, 5, 8, 103, 8, 8, 10, 8, 12, 8,
		106, 9, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 3, 14, 130, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 5, 14, 138, 8, 14, 10, 14, 12, 14, 141, 9, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 147, 8, 15, 1, 16, 1, 16, 1, 16, 4, 16, 152, 8, 16, 11,
		16, 12, 16, 153, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16,
		163, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 171, 8, 16,
		10, 16, 12, 16, 174, 9, 16, 1, 16, 0, 2, 28, 32, 17, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 7, 1, 0, 21, 22, 2, 0, 7,
		10, 16, 16, 2, 0, 11, 16, 18, 20, 1, 0, 24, 27, 1, 0, 28, 33, 1, 0, 36,
		37, 1, 0, 34, 35, 185, 0, 34, 1, 0, 0, 0, 2, 47, 1, 0, 0, 0, 4, 53, 1,
		0, 0, 0, 6, 72, 1, 0, 0, 0, 8, 74, 1, 0, 0, 0, 10, 78, 1, 0, 0, 0, 12,
		90, 1, 0, 0, 0, 14, 93, 1, 0, 0, 0, 16, 96, 1, 0, 0, 0, 18, 107, 1, 0,
		0, 0, 20, 109, 1, 0, 0, 0, 22, 111, 1, 0, 0, 0, 24, 114, 1, 0, 0, 0, 26,
		116, 1, 0, 0, 0, 28, 129, 1, 0, 0, 0, 30, 142, 1, 0, 0, 0, 32, 162, 1,
		0, 0, 0, 34, 35, 3, 2, 1, 0, 35, 37, 3, 8, 4, 0, 36, 38, 3, 10, 5, 0, 37,
		36, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 41, 3, 12,
		6, 0, 40, 39, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 43, 1, 0, 0, 0, 42, 44,
		3, 14, 7, 0, 43, 42, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0,
		45, 46, 5, 0, 0, 1, 46, 1, 1, 0, 0, 0, 47, 49, 5, 1, 0, 0, 48, 50, 3, 24,
		12, 0, 49, 48, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51,
		52, 3, 4, 2, 0, 52, 3, 1, 0, 0, 0, 53, 58, 3, 6, 3, 0, 54, 55, 5, 43, 0,
		0, 55, 57, 3, 6, 3, 0, 56, 54, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56,
		1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 5, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0,
		61, 73, 3, 20, 10, 0, 62, 73, 3, 22, 11, 0, 63, 73, 3, 30, 15, 0, 64, 67,
		5, 41, 0, 0, 65, 68, 3, 32, 16, 0, 66, 68, 3, 28, 14, 0, 67, 65, 1, 0,
		0, 0, 67, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 5, 42, 0, 0, 70,
		71, 5, 46, 0, 0, 71, 73, 1, 0, 0, 0, 72, 61, 1, 0, 0, 0, 72, 62, 1, 0,
		0, 0, 72, 63, 1, 0, 0, 0, 72, 64, 1, 0, 0, 0, 73, 7, 1, 0, 0, 0, 74, 75,
		5, 2, 0, 0, 75, 76, 3, 18, 9, 0, 76, 77, 5, 46, 0, 0, 77, 9, 1, 0, 0, 0,
		78, 81, 5, 3, 0, 0, 79, 82, 3, 18, 9, 0, 80, 82, 5, 17, 0, 0, 81, 79, 1,
		0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 88, 1, 0, 0, 0, 83, 86, 5, 43, 0, 0, 84,
		87, 3, 18, 9, 0, 85, 87, 5, 17, 0, 0, 86, 84, 1, 0, 0, 0, 86, 85, 1, 0,
		0, 0, 87, 89, 1, 0, 0, 0, 88, 83, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 11,
		1, 0, 0, 0, 90, 91, 5, 4, 0, 0, 91, 92, 3, 28, 14, 0, 92, 13, 1, 0, 0,
		0, 93, 94, 5, 5, 0, 0, 94, 95, 3, 16, 8, 0, 95, 15, 1, 0, 0, 0, 96, 98,
		3, 32, 16, 0, 97, 99, 7, 0, 0, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0, 0,
		0, 99, 104, 1, 0, 0, 0, 100, 101, 5, 43, 0, 0, 101, 103, 3, 16, 8, 0, 102,
		100, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105,
		1, 0, 0, 0, 105, 17, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 107, 108, 7, 1,
		0, 0, 108, 19, 1, 0, 0, 0, 109, 110, 7, 2, 0, 0, 110, 21, 1, 0, 0, 0, 111,
		112, 7, 3, 0, 0, 112, 113, 3, 20, 10, 0, 113, 23, 1, 0, 0, 0, 114, 115,
		5, 23, 0, 0, 115, 25, 1, 0, 0, 0, 116, 117, 7, 4, 0, 0, 117, 27, 1, 0,
		0, 0, 118, 119, 6, 14, -1, 0, 119, 120, 5, 41, 0, 0, 120, 121, 3, 28, 14,
		0, 121, 122, 5, 42, 0, 0, 122, 130, 1, 0, 0, 0, 123, 124, 5, 40, 0, 0,
		124, 130, 3, 28, 14, 4, 125, 126, 3, 32, 16, 0, 126, 127, 3, 26, 13, 0,
		127, 128, 3, 32, 16, 0, 128, 130, 1, 0, 0, 0, 129, 118, 1, 0, 0, 0, 129,
		123, 1, 0, 0, 0, 129, 125, 1, 0, 0, 0, 130, 139, 1, 0, 0, 0, 131, 132,
		10, 3, 0, 0, 132, 133, 5, 38, 0, 0, 133, 138, 3, 28, 14, 4, 134, 135, 10,
		2, 0, 0, 135, 136, 5, 39, 0, 0, 136, 138, 3, 28, 14, 3, 137, 131, 1, 0,
		0, 0, 137, 134, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0,
		139, 140, 1, 0, 0, 0, 140, 29, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 143,
		5, 46, 0, 0, 143, 146, 5, 6, 0, 0, 144, 147, 3, 20, 10, 0, 145, 147, 3,
		22, 11, 0, 146, 144, 1, 0, 0, 0, 146, 145, 1, 0, 0, 0, 147, 31, 1, 0, 0,
		0, 148, 149, 6, 16, -1, 0, 149, 151, 5, 41, 0, 0, 150, 152, 3, 32, 16,
		0, 151, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 5, 42, 0, 0, 156, 163,
		1, 0, 0, 0, 157, 163, 3, 22, 11, 0, 158, 163, 3, 30, 15, 0, 159, 163, 5,
		46, 0, 0, 160, 163, 3, 20, 10, 0, 161, 163, 5, 45, 0, 0, 162, 148, 1, 0,
		0, 0, 162, 157, 1, 0, 0, 0, 162, 158, 1, 0, 0, 0, 162, 159, 1, 0, 0, 0,
		162, 160, 1, 0, 0, 0, 162, 161, 1, 0, 0, 0, 163, 172, 1, 0, 0, 0, 164,
		165, 10, 7, 0, 0, 165, 166, 7, 5, 0, 0, 166, 171, 3, 32, 16, 8, 167, 168,
		10, 6, 0, 0, 168, 169, 7, 6, 0, 0, 169, 171, 3, 32, 16, 7, 170, 164, 1,
		0, 0, 0, 170, 167, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0,
		0, 172, 173, 1, 0, 0, 0, 173, 33, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 20,
		37, 40, 43, 49, 58, 67, 72, 81, 86, 88, 98, 104, 129, 137, 139, 146, 153,
		162, 170, 172,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// pgqlParserInit initializes any static state used to implement pgqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewpgqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PgqlParserInit() {
	staticData := &PgqlParserStaticData
	staticData.once.Do(pgqlParserInit)
}

// NewpgqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewpgqlParser(input antlr.TokenStream) *pgqlParser {
	PgqlParserInit()
	this := new(pgqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PgqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "pgql.g4"

	return this
}

// pgqlParser tokens.
const (
	pgqlParserEOF          = antlr.TokenEOF
	pgqlParserT__0         = 1
	pgqlParserT__1         = 2
	pgqlParserT__2         = 3
	pgqlParserT__3         = 4
	pgqlParserT__4         = 5
	pgqlParserT__5         = 6
	pgqlParserPLAYER       = 7
	pgqlParserTEAM         = 8
	pgqlParserDATE         = 9
	pgqlParserTYPE         = 10
	pgqlParserKILLS        = 11
	pgqlParserDAMAGE       = 12
	pgqlParserASSISTS      = 13
	pgqlParserRESCUES      = 14
	pgqlParserRECALLS      = 15
	pgqlParserWIN          = 16
	pgqlParserGAME         = 17
	pgqlParserTOTALWINS    = 18
	pgqlParserTOTALLOSES   = 19
	pgqlParserGAMECOUNT    = 20
	pgqlParserDESC         = 21
	pgqlParserASC          = 22
	pgqlParserRUNNING      = 23
	pgqlParserAVERAGE      = 24
	pgqlParserTOTAL        = 25
	pgqlParserMIN          = 26
	pgqlParserMAX          = 27
	pgqlParserLESSEREQUAL  = 28
	pgqlParserGREATEREQUAL = 29
	pgqlParserGREATER      = 30
	pgqlParserLESSER       = 31
	pgqlParserEQUAL        = 32
	pgqlParserNOTEQUAL     = 33
	pgqlParserSUM          = 34
	pgqlParserDIFFERENCE   = 35
	pgqlParserMULTIPLY     = 36
	pgqlParserDIVIDE       = 37
	pgqlParserLOGICALAND   = 38
	pgqlParserLOGICALOR    = 39
	pgqlParserLOGICALNOT   = 40
	pgqlParserLPAREN       = 41
	pgqlParserRPAREN       = 42
	pgqlParserCOMMA        = 43
	pgqlParserSTRING       = 44
	pgqlParserNUMBER       = 45
	pgqlParserIDENTIFIER   = 46
	pgqlParserWS           = 47
)

// pgqlParser rules.
const (
	pgqlParserRULE_statement          = 0
	pgqlParserRULE_showClause         = 1
	pgqlParserRULE_showBody           = 2
	pgqlParserRULE_showFragment       = 3
	pgqlParserRULE_forClause          = 4
	pgqlParserRULE_byClause           = 5
	pgqlParserRULE_whereClause        = 6
	pgqlParserRULE_sortByClause       = 7
	pgqlParserRULE_sortBody           = 8
	pgqlParserRULE_dimension          = 9
	pgqlParserRULE_measure            = 10
	pgqlParserRULE_aggregateFunction  = 11
	pgqlParserRULE_sequentialFunction = 12
	pgqlParserRULE_comparisonOperator = 13
	pgqlParserRULE_predicate          = 14
	pgqlParserRULE_scope              = 15
	pgqlParserRULE_expr               = 16
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ShowClause() IShowClauseContext
	ForClause() IForClauseContext
	EOF() antlr.TerminalNode
	ByClause() IByClauseContext
	WhereClause() IWhereClauseContext
	SortByClause() ISortByClauseContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ShowClause() IShowClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShowClauseContext)
}

func (s *StatementContext) ForClause() IForClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForClauseContext)
}

func (s *StatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(pgqlParserEOF, 0)
}

func (s *StatementContext) ByClause() IByClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IByClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IByClauseContext)
}

func (s *StatementContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *StatementContext) SortByClause() ISortByClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortByClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortByClauseContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, pgqlParserRULE_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.ShowClause()
	}
	{
		p.SetState(35)
		p.ForClause()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == pgqlParserT__2 {
		{
			p.SetState(36)
			p.ByClause()
		}

	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == pgqlParserT__3 {
		{
			p.SetState(39)
			p.WhereClause()
		}

	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == pgqlParserT__4 {
		{
			p.SetState(42)
			p.SortByClause()
		}

	}
	{
		p.SetState(45)
		p.Match(pgqlParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IShowClauseContext is an interface to support dynamic dispatch.
type IShowClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ShowBody() IShowBodyContext
	SequentialFunction() ISequentialFunctionContext

	// IsShowClauseContext differentiates from other interfaces.
	IsShowClauseContext()
}

type ShowClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowClauseContext() *ShowClauseContext {
	var p = new(ShowClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_showClause
	return p
}

func InitEmptyShowClauseContext(p *ShowClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_showClause
}

func (*ShowClauseContext) IsShowClauseContext() {}

func NewShowClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowClauseContext {
	var p = new(ShowClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_showClause

	return p
}

func (s *ShowClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowClauseContext) ShowBody() IShowBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShowBodyContext)
}

func (s *ShowClauseContext) SequentialFunction() ISequentialFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequentialFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequentialFunctionContext)
}

func (s *ShowClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitShowClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) ShowClause() (localctx IShowClauseContext) {
	localctx = NewShowClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, pgqlParserRULE_showClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(pgqlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == pgqlParserRUNNING {
		{
			p.SetState(48)
			p.SequentialFunction()
		}

	}
	{
		p.SetState(51)
		p.ShowBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IShowBodyContext is an interface to support dynamic dispatch.
type IShowBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllShowFragment() []IShowFragmentContext
	ShowFragment(i int) IShowFragmentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsShowBodyContext differentiates from other interfaces.
	IsShowBodyContext()
}

type ShowBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowBodyContext() *ShowBodyContext {
	var p = new(ShowBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_showBody
	return p
}

func InitEmptyShowBodyContext(p *ShowBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_showBody
}

func (*ShowBodyContext) IsShowBodyContext() {}

func NewShowBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowBodyContext {
	var p = new(ShowBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_showBody

	return p
}

func (s *ShowBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowBodyContext) AllShowFragment() []IShowFragmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IShowFragmentContext); ok {
			len++
		}
	}

	tst := make([]IShowFragmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IShowFragmentContext); ok {
			tst[i] = t.(IShowFragmentContext)
			i++
		}
	}

	return tst
}

func (s *ShowBodyContext) ShowFragment(i int) IShowFragmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowFragmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShowFragmentContext)
}

func (s *ShowBodyContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(pgqlParserCOMMA)
}

func (s *ShowBodyContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(pgqlParserCOMMA, i)
}

func (s *ShowBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitShowBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) ShowBody() (localctx IShowBodyContext) {
	localctx = NewShowBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, pgqlParserRULE_showBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.ShowFragment()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == pgqlParserCOMMA {
		{
			p.SetState(54)
			p.Match(pgqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.ShowFragment()
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IShowFragmentContext is an interface to support dynamic dispatch.
type IShowFragmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Measure() IMeasureContext
	AggregateFunction() IAggregateFunctionContext
	Scope() IScopeContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Expr() IExprContext
	Predicate() IPredicateContext

	// IsShowFragmentContext differentiates from other interfaces.
	IsShowFragmentContext()
}

type ShowFragmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowFragmentContext() *ShowFragmentContext {
	var p = new(ShowFragmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_showFragment
	return p
}

func InitEmptyShowFragmentContext(p *ShowFragmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_showFragment
}

func (*ShowFragmentContext) IsShowFragmentContext() {}

func NewShowFragmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowFragmentContext {
	var p = new(ShowFragmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_showFragment

	return p
}

func (s *ShowFragmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowFragmentContext) Measure() IMeasureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeasureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeasureContext)
}

func (s *ShowFragmentContext) AggregateFunction() IAggregateFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFunctionContext)
}

func (s *ShowFragmentContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *ShowFragmentContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(pgqlParserLPAREN, 0)
}

func (s *ShowFragmentContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(pgqlParserRPAREN, 0)
}

func (s *ShowFragmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pgqlParserIDENTIFIER, 0)
}

func (s *ShowFragmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ShowFragmentContext) Predicate() IPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *ShowFragmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowFragmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowFragmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitShowFragment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) ShowFragment() (localctx IShowFragmentContext) {
	localctx = NewShowFragmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, pgqlParserRULE_showFragment)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case pgqlParserKILLS, pgqlParserDAMAGE, pgqlParserASSISTS, pgqlParserRESCUES, pgqlParserRECALLS, pgqlParserWIN, pgqlParserTOTALWINS, pgqlParserTOTALLOSES, pgqlParserGAMECOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Measure()
		}

	case pgqlParserAVERAGE, pgqlParserTOTAL, pgqlParserMIN, pgqlParserMAX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.AggregateFunction()
		}

	case pgqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.Scope()
		}

	case pgqlParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)
			p.Match(pgqlParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(65)
				p.expr(0)
			}

		case 2:
			{
				p.SetState(66)
				p.predicate(0)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(69)
			p.Match(pgqlParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(pgqlParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForClauseContext is an interface to support dynamic dispatch.
type IForClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Dimension() IDimensionContext
	IDENTIFIER() antlr.TerminalNode

	// IsForClauseContext differentiates from other interfaces.
	IsForClauseContext()
}

type ForClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForClauseContext() *ForClauseContext {
	var p = new(ForClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_forClause
	return p
}

func InitEmptyForClauseContext(p *ForClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_forClause
}

func (*ForClauseContext) IsForClauseContext() {}

func NewForClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForClauseContext {
	var p = new(ForClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_forClause

	return p
}

func (s *ForClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ForClauseContext) Dimension() IDimensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimensionContext)
}

func (s *ForClauseContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pgqlParserIDENTIFIER, 0)
}

func (s *ForClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitForClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) ForClause() (localctx IForClauseContext) {
	localctx = NewForClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, pgqlParserRULE_forClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(pgqlParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Dimension()
	}
	{
		p.SetState(76)
		p.Match(pgqlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IByClauseContext is an interface to support dynamic dispatch.
type IByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDimension() []IDimensionContext
	Dimension(i int) IDimensionContext
	AllGAME() []antlr.TerminalNode
	GAME(i int) antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsByClauseContext differentiates from other interfaces.
	IsByClauseContext()
}

type ByClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyByClauseContext() *ByClauseContext {
	var p = new(ByClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_byClause
	return p
}

func InitEmptyByClauseContext(p *ByClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_byClause
}

func (*ByClauseContext) IsByClauseContext() {}

func NewByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByClauseContext {
	var p = new(ByClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_byClause

	return p
}

func (s *ByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ByClauseContext) AllDimension() []IDimensionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDimensionContext); ok {
			len++
		}
	}

	tst := make([]IDimensionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDimensionContext); ok {
			tst[i] = t.(IDimensionContext)
			i++
		}
	}

	return tst
}

func (s *ByClauseContext) Dimension(i int) IDimensionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimensionContext)
}

func (s *ByClauseContext) AllGAME() []antlr.TerminalNode {
	return s.GetTokens(pgqlParserGAME)
}

func (s *ByClauseContext) GAME(i int) antlr.TerminalNode {
	return s.GetToken(pgqlParserGAME, i)
}

func (s *ByClauseContext) COMMA() antlr.TerminalNode {
	return s.GetToken(pgqlParserCOMMA, 0)
}

func (s *ByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitByClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) ByClause() (localctx IByClauseContext) {
	localctx = NewByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, pgqlParserRULE_byClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(pgqlParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case pgqlParserPLAYER, pgqlParserTEAM, pgqlParserDATE, pgqlParserTYPE, pgqlParserWIN:
		{
			p.SetState(79)
			p.Dimension()
		}

	case pgqlParserGAME:
		{
			p.SetState(80)
			p.Match(pgqlParserGAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == pgqlParserCOMMA {
		{
			p.SetState(83)
			p.Match(pgqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case pgqlParserPLAYER, pgqlParserTEAM, pgqlParserDATE, pgqlParserTYPE, pgqlParserWIN:
			{
				p.SetState(84)
				p.Dimension()
			}

		case pgqlParserGAME:
			{
				p.SetState(85)
				p.Match(pgqlParserGAME)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Predicate() IPredicateContext

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_whereClause
	return p
}

func InitEmptyWhereClauseContext(p *WhereClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_whereClause
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) Predicate() IPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitWhereClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, pgqlParserRULE_whereClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(pgqlParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.predicate(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISortByClauseContext is an interface to support dynamic dispatch.
type ISortByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SortBody() ISortBodyContext

	// IsSortByClauseContext differentiates from other interfaces.
	IsSortByClauseContext()
}

type SortByClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortByClauseContext() *SortByClauseContext {
	var p = new(SortByClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_sortByClause
	return p
}

func InitEmptySortByClauseContext(p *SortByClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_sortByClause
}

func (*SortByClauseContext) IsSortByClauseContext() {}

func NewSortByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortByClauseContext {
	var p = new(SortByClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_sortByClause

	return p
}

func (s *SortByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SortByClauseContext) SortBody() ISortBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortBodyContext)
}

func (s *SortByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitSortByClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) SortByClause() (localctx ISortByClauseContext) {
	localctx = NewSortByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, pgqlParserRULE_sortByClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(pgqlParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.SortBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISortBodyContext is an interface to support dynamic dispatch.
type ISortBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllSortBody() []ISortBodyContext
	SortBody(i int) ISortBodyContext
	DESC() antlr.TerminalNode
	ASC() antlr.TerminalNode

	// IsSortBodyContext differentiates from other interfaces.
	IsSortBodyContext()
}

type SortBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortBodyContext() *SortBodyContext {
	var p = new(SortBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_sortBody
	return p
}

func InitEmptySortBodyContext(p *SortBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_sortBody
}

func (*SortBodyContext) IsSortBodyContext() {}

func NewSortBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortBodyContext {
	var p = new(SortBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_sortBody

	return p
}

func (s *SortBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SortBodyContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SortBodyContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(pgqlParserCOMMA)
}

func (s *SortBodyContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(pgqlParserCOMMA, i)
}

func (s *SortBodyContext) AllSortBody() []ISortBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortBodyContext); ok {
			len++
		}
	}

	tst := make([]ISortBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortBodyContext); ok {
			tst[i] = t.(ISortBodyContext)
			i++
		}
	}

	return tst
}

func (s *SortBodyContext) SortBody(i int) ISortBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortBodyContext)
}

func (s *SortBodyContext) DESC() antlr.TerminalNode {
	return s.GetToken(pgqlParserDESC, 0)
}

func (s *SortBodyContext) ASC() antlr.TerminalNode {
	return s.GetToken(pgqlParserASC, 0)
}

func (s *SortBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitSortBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) SortBody() (localctx ISortBodyContext) {
	localctx = NewSortBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, pgqlParserRULE_sortBody)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.expr(0)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == pgqlParserDESC || _la == pgqlParserASC {
		{
			p.SetState(97)
			_la = p.GetTokenStream().LA(1)

			if !(_la == pgqlParserDESC || _la == pgqlParserASC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(100)
				p.Match(pgqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(101)
				p.SortBody()
			}

		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDimensionContext is an interface to support dynamic dispatch.
type IDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLAYER() antlr.TerminalNode
	TEAM() antlr.TerminalNode
	DATE() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	WIN() antlr.TerminalNode

	// IsDimensionContext differentiates from other interfaces.
	IsDimensionContext()
}

type DimensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDimensionContext() *DimensionContext {
	var p = new(DimensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_dimension
	return p
}

func InitEmptyDimensionContext(p *DimensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_dimension
}

func (*DimensionContext) IsDimensionContext() {}

func NewDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DimensionContext {
	var p = new(DimensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_dimension

	return p
}

func (s *DimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *DimensionContext) PLAYER() antlr.TerminalNode {
	return s.GetToken(pgqlParserPLAYER, 0)
}

func (s *DimensionContext) TEAM() antlr.TerminalNode {
	return s.GetToken(pgqlParserTEAM, 0)
}

func (s *DimensionContext) DATE() antlr.TerminalNode {
	return s.GetToken(pgqlParserDATE, 0)
}

func (s *DimensionContext) TYPE() antlr.TerminalNode {
	return s.GetToken(pgqlParserTYPE, 0)
}

func (s *DimensionContext) WIN() antlr.TerminalNode {
	return s.GetToken(pgqlParserWIN, 0)
}

func (s *DimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DimensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitDimension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) Dimension() (localctx IDimensionContext) {
	localctx = NewDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, pgqlParserRULE_dimension)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67456) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMeasureContext is an interface to support dynamic dispatch.
type IMeasureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KILLS() antlr.TerminalNode
	DAMAGE() antlr.TerminalNode
	ASSISTS() antlr.TerminalNode
	RESCUES() antlr.TerminalNode
	RECALLS() antlr.TerminalNode
	WIN() antlr.TerminalNode
	TOTALLOSES() antlr.TerminalNode
	TOTALWINS() antlr.TerminalNode
	GAMECOUNT() antlr.TerminalNode

	// IsMeasureContext differentiates from other interfaces.
	IsMeasureContext()
}

type MeasureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeasureContext() *MeasureContext {
	var p = new(MeasureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_measure
	return p
}

func InitEmptyMeasureContext(p *MeasureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_measure
}

func (*MeasureContext) IsMeasureContext() {}

func NewMeasureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MeasureContext {
	var p = new(MeasureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_measure

	return p
}

func (s *MeasureContext) GetParser() antlr.Parser { return s.parser }

func (s *MeasureContext) KILLS() antlr.TerminalNode {
	return s.GetToken(pgqlParserKILLS, 0)
}

func (s *MeasureContext) DAMAGE() antlr.TerminalNode {
	return s.GetToken(pgqlParserDAMAGE, 0)
}

func (s *MeasureContext) ASSISTS() antlr.TerminalNode {
	return s.GetToken(pgqlParserASSISTS, 0)
}

func (s *MeasureContext) RESCUES() antlr.TerminalNode {
	return s.GetToken(pgqlParserRESCUES, 0)
}

func (s *MeasureContext) RECALLS() antlr.TerminalNode {
	return s.GetToken(pgqlParserRECALLS, 0)
}

func (s *MeasureContext) WIN() antlr.TerminalNode {
	return s.GetToken(pgqlParserWIN, 0)
}

func (s *MeasureContext) TOTALLOSES() antlr.TerminalNode {
	return s.GetToken(pgqlParserTOTALLOSES, 0)
}

func (s *MeasureContext) TOTALWINS() antlr.TerminalNode {
	return s.GetToken(pgqlParserTOTALWINS, 0)
}

func (s *MeasureContext) GAMECOUNT() antlr.TerminalNode {
	return s.GetToken(pgqlParserGAMECOUNT, 0)
}

func (s *MeasureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeasureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MeasureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitMeasure(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) Measure() (localctx IMeasureContext) {
	localctx = NewMeasureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, pgqlParserRULE_measure)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1964032) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggregateFunctionContext is an interface to support dynamic dispatch.
type IAggregateFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Measure() IMeasureContext
	AVERAGE() antlr.TerminalNode
	TOTAL() antlr.TerminalNode
	MAX() antlr.TerminalNode
	MIN() antlr.TerminalNode

	// IsAggregateFunctionContext differentiates from other interfaces.
	IsAggregateFunctionContext()
}

type AggregateFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateFunctionContext() *AggregateFunctionContext {
	var p = new(AggregateFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_aggregateFunction
	return p
}

func InitEmptyAggregateFunctionContext(p *AggregateFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_aggregateFunction
}

func (*AggregateFunctionContext) IsAggregateFunctionContext() {}

func NewAggregateFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateFunctionContext {
	var p = new(AggregateFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_aggregateFunction

	return p
}

func (s *AggregateFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateFunctionContext) Measure() IMeasureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeasureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeasureContext)
}

func (s *AggregateFunctionContext) AVERAGE() antlr.TerminalNode {
	return s.GetToken(pgqlParserAVERAGE, 0)
}

func (s *AggregateFunctionContext) TOTAL() antlr.TerminalNode {
	return s.GetToken(pgqlParserTOTAL, 0)
}

func (s *AggregateFunctionContext) MAX() antlr.TerminalNode {
	return s.GetToken(pgqlParserMAX, 0)
}

func (s *AggregateFunctionContext) MIN() antlr.TerminalNode {
	return s.GetToken(pgqlParserMIN, 0)
}

func (s *AggregateFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitAggregateFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) AggregateFunction() (localctx IAggregateFunctionContext) {
	localctx = NewAggregateFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, pgqlParserRULE_aggregateFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&251658240) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(112)
		p.Measure()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISequentialFunctionContext is an interface to support dynamic dispatch.
type ISequentialFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RUNNING() antlr.TerminalNode

	// IsSequentialFunctionContext differentiates from other interfaces.
	IsSequentialFunctionContext()
}

type SequentialFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequentialFunctionContext() *SequentialFunctionContext {
	var p = new(SequentialFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_sequentialFunction
	return p
}

func InitEmptySequentialFunctionContext(p *SequentialFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_sequentialFunction
}

func (*SequentialFunctionContext) IsSequentialFunctionContext() {}

func NewSequentialFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequentialFunctionContext {
	var p = new(SequentialFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_sequentialFunction

	return p
}

func (s *SequentialFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *SequentialFunctionContext) RUNNING() antlr.TerminalNode {
	return s.GetToken(pgqlParserRUNNING, 0)
}

func (s *SequentialFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequentialFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequentialFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitSequentialFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) SequentialFunction() (localctx ISequentialFunctionContext) {
	localctx = NewSequentialFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, pgqlParserRULE_sequentialFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(pgqlParserRUNNING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LESSER() antlr.TerminalNode
	GREATER() antlr.TerminalNode
	LESSEREQUAL() antlr.TerminalNode
	GREATEREQUAL() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NOTEQUAL() antlr.TerminalNode

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_comparisonOperator
	return p
}

func InitEmptyComparisonOperatorContext(p *ComparisonOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_comparisonOperator
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) LESSER() antlr.TerminalNode {
	return s.GetToken(pgqlParserLESSER, 0)
}

func (s *ComparisonOperatorContext) GREATER() antlr.TerminalNode {
	return s.GetToken(pgqlParserGREATER, 0)
}

func (s *ComparisonOperatorContext) LESSEREQUAL() antlr.TerminalNode {
	return s.GetToken(pgqlParserLESSEREQUAL, 0)
}

func (s *ComparisonOperatorContext) GREATEREQUAL() antlr.TerminalNode {
	return s.GetToken(pgqlParserGREATEREQUAL, 0)
}

func (s *ComparisonOperatorContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(pgqlParserEQUAL, 0)
}

func (s *ComparisonOperatorContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(pgqlParserNOTEQUAL, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, pgqlParserRULE_comparisonOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16911433728) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllPredicate() []IPredicateContext
	Predicate(i int) IPredicateContext
	RPAREN() antlr.TerminalNode
	LOGICALNOT() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	ComparisonOperator() IComparisonOperatorContext
	LOGICALAND() antlr.TerminalNode
	LOGICALOR() antlr.TerminalNode

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_predicate
	return p
}

func InitEmptyPredicateContext(p *PredicateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_predicate
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(pgqlParserLPAREN, 0)
}

func (s *PredicateContext) AllPredicate() []IPredicateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPredicateContext); ok {
			len++
		}
	}

	tst := make([]IPredicateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPredicateContext); ok {
			tst[i] = t.(IPredicateContext)
			i++
		}
	}

	return tst
}

func (s *PredicateContext) Predicate(i int) IPredicateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredicateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *PredicateContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(pgqlParserRPAREN, 0)
}

func (s *PredicateContext) LOGICALNOT() antlr.TerminalNode {
	return s.GetToken(pgqlParserLOGICALNOT, 0)
}

func (s *PredicateContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PredicateContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PredicateContext) ComparisonOperator() IComparisonOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *PredicateContext) LOGICALAND() antlr.TerminalNode {
	return s.GetToken(pgqlParserLOGICALAND, 0)
}

func (s *PredicateContext) LOGICALOR() antlr.TerminalNode {
	return s.GetToken(pgqlParserLOGICALOR, 0)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) Predicate() (localctx IPredicateContext) {
	return p.predicate(0)
}

func (p *pgqlParser) predicate(_p int) (localctx IPredicateContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPredicateContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, pgqlParserRULE_predicate, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(119)
			p.Match(pgqlParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.predicate(0)
		}
		{
			p.SetState(121)
			p.Match(pgqlParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(123)
			p.Match(pgqlParserLOGICALNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.predicate(4)
		}

	case 3:
		{
			p.SetState(125)
			p.expr(0)
		}
		{
			p.SetState(126)
			p.ComparisonOperator()
		}
		{
			p.SetState(127)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPredicateContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pgqlParserRULE_predicate)
				p.SetState(131)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(132)
					p.Match(pgqlParserLOGICALAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(133)
					p.predicate(4)
				}

			case 2:
				localctx = NewPredicateContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pgqlParserRULE_predicate)
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(135)
					p.Match(pgqlParserLOGICALOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(136)
					p.predicate(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Measure() IMeasureContext
	AggregateFunction() IAggregateFunctionContext

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_scope
	return p
}

func InitEmptyScopeContext(p *ScopeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_scope
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pgqlParserIDENTIFIER, 0)
}

func (s *ScopeContext) Measure() IMeasureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeasureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeasureContext)
}

func (s *ScopeContext) AggregateFunction() IAggregateFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFunctionContext)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) Scope() (localctx IScopeContext) {
	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, pgqlParserRULE_scope)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(pgqlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(pgqlParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case pgqlParserKILLS, pgqlParserDAMAGE, pgqlParserASSISTS, pgqlParserRESCUES, pgqlParserRECALLS, pgqlParserWIN, pgqlParserTOTALWINS, pgqlParserTOTALLOSES, pgqlParserGAMECOUNT:
		{
			p.SetState(144)
			p.Measure()
		}

	case pgqlParserAVERAGE, pgqlParserTOTAL, pgqlParserMIN, pgqlParserMAX:
		{
			p.SetState(145)
			p.AggregateFunction()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AggregateFunction() IAggregateFunctionContext
	Scope() IScopeContext
	IDENTIFIER() antlr.TerminalNode
	Measure() IMeasureContext
	NUMBER() antlr.TerminalNode
	MULTIPLY() antlr.TerminalNode
	DIVIDE() antlr.TerminalNode
	SUM() antlr.TerminalNode
	DIFFERENCE() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(pgqlParserLPAREN, 0)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(pgqlParserRPAREN, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) AggregateFunction() IAggregateFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFunctionContext)
}

func (s *ExprContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *ExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pgqlParserIDENTIFIER, 0)
}

func (s *ExprContext) Measure() IMeasureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeasureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeasureContext)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(pgqlParserNUMBER, 0)
}

func (s *ExprContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(pgqlParserMULTIPLY, 0)
}

func (s *ExprContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(pgqlParserDIVIDE, 0)
}

func (s *ExprContext) SUM() antlr.TerminalNode {
	return s.GetToken(pgqlParserSUM, 0)
}

func (s *ExprContext) DIFFERENCE() antlr.TerminalNode {
	return s.GetToken(pgqlParserDIFFERENCE, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *pgqlParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, pgqlParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(149)
			p.Match(pgqlParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&107752393144320) != 0) {
			{
				p.SetState(150)
				p.expr(0)
			}

			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(155)
			p.Match(pgqlParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(157)
			p.AggregateFunction()
		}

	case 3:
		{
			p.SetState(158)
			p.Scope()
		}

	case 4:
		{
			p.SetState(159)
			p.Match(pgqlParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(160)
			p.Measure()
		}

	case 6:
		{
			p.SetState(161)
			p.Match(pgqlParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pgqlParserRULE_expr)
				p.SetState(164)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(165)
					_la = p.GetTokenStream().LA(1)

					if !(_la == pgqlParserMULTIPLY || _la == pgqlParserDIVIDE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(166)
					p.expr(8)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pgqlParserRULE_expr)
				p.SetState(167)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(168)
					_la = p.GetTokenStream().LA(1)

					if !(_la == pgqlParserSUM || _la == pgqlParserDIFFERENCE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(169)
					p.expr(7)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *pgqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *PredicateContext = nil
		if localctx != nil {
			t = localctx.(*PredicateContext)
		}
		return p.Predicate_Sempred(t, predIndex)

	case 16:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *pgqlParser) Predicate_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *pgqlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
