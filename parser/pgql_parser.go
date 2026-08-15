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
		"", "';'", "'show'", "'and'", "','", "'as'", "'for'", "'by'", "'where'",
		"'sort by'", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'<'", "'>'",
		"'<='", "'>='", "'='", "'player'", "'team'", "'date'", "'type'", "'kills'",
		"'damage'", "'assists'", "'rescues'", "'recalls'", "'win'", "'game'",
		"'desc'", "'asc'", "'running'", "'average'", "'total'", "'min'", "'max'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "PLAYER", "TEAM", "DATE", "TYPE", "KILLS", "DAMAGE",
		"ASSISTS", "RESCUES", "RECALLS", "WIN", "GAME", "DESC", "ASC", "RUNNING",
		"AVERAGE", "TOTAL", "MIN", "MAX", "STRING", "NUMBER", "WS",
	}
	staticData.RuleNames = []string{
		"statement", "showClause", "showBody", "showFragment", "showFunction",
		"forClause", "byClause", "whereClause", "sortByClause", "dimension",
		"measure", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 131, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 3, 0, 28, 8, 0, 1, 0, 3, 0, 31, 8,
		0, 1, 0, 3, 0, 34, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 4, 1, 40, 8, 1, 11, 1,
		12, 1, 41, 1, 2, 1, 2, 1, 2, 5, 2, 47, 8, 2, 10, 2, 12, 2, 50, 9, 2, 1,
		3, 3, 3, 53, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 60, 8, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 71, 8, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 76, 8, 6, 5, 6, 78, 8, 6, 10, 6, 12, 6, 81, 9, 6, 1, 7,
		1, 7, 4, 7, 85, 8, 7, 11, 7, 12, 7, 86, 1, 8, 1, 8, 1, 8, 5, 8, 92, 8,
		8, 10, 8, 12, 8, 95, 9, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		4, 11, 104, 8, 11, 11, 11, 12, 11, 105, 1, 11, 1, 11, 1, 11, 1, 11, 3,
		11, 112, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 126, 8, 11, 10, 11, 12, 11, 129, 9,
		11, 1, 11, 0, 1, 22, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0,
		8, 1, 0, 3, 4, 1, 0, 34, 38, 1, 0, 32, 33, 2, 0, 21, 24, 30, 30, 1, 0,
		25, 30, 1, 0, 12, 13, 1, 0, 14, 15, 1, 0, 16, 19, 137, 0, 24, 1, 0, 0,
		0, 2, 37, 1, 0, 0, 0, 4, 43, 1, 0, 0, 0, 6, 52, 1, 0, 0, 0, 8, 61, 1, 0,
		0, 0, 10, 63, 1, 0, 0, 0, 12, 67, 1, 0, 0, 0, 14, 82, 1, 0, 0, 0, 16, 88,
		1, 0, 0, 0, 18, 96, 1, 0, 0, 0, 20, 98, 1, 0, 0, 0, 22, 111, 1, 0, 0, 0,
		24, 25, 3, 2, 1, 0, 25, 27, 3, 10, 5, 0, 26, 28, 3, 12, 6, 0, 27, 26, 1,
		0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 30, 1, 0, 0, 0, 29, 31, 3, 14, 7, 0, 30,
		29, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 33, 1, 0, 0, 0, 32, 34, 3, 16,
		8, 0, 33, 32, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36,
		5, 1, 0, 0, 36, 1, 1, 0, 0, 0, 37, 39, 5, 2, 0, 0, 38, 40, 3, 4, 2, 0,
		39, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1,
		0, 0, 0, 42, 3, 1, 0, 0, 0, 43, 48, 3, 6, 3, 0, 44, 45, 7, 0, 0, 0, 45,
		47, 3, 6, 3, 0, 46, 44, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0,
		0, 48, 49, 1, 0, 0, 0, 49, 5, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 53, 3,
		8, 4, 0, 52, 51, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 59, 1, 0, 0, 0, 54,
		60, 3, 20, 10, 0, 55, 56, 3, 22, 11, 0, 56, 57, 5, 5, 0, 0, 57, 58, 5,
		39, 0, 0, 58, 60, 1, 0, 0, 0, 59, 54, 1, 0, 0, 0, 59, 55, 1, 0, 0, 0, 60,
		7, 1, 0, 0, 0, 61, 62, 7, 1, 0, 0, 62, 9, 1, 0, 0, 0, 63, 64, 5, 6, 0,
		0, 64, 65, 3, 18, 9, 0, 65, 66, 5, 39, 0, 0, 66, 11, 1, 0, 0, 0, 67, 70,
		5, 7, 0, 0, 68, 71, 3, 18, 9, 0, 69, 71, 5, 31, 0, 0, 70, 68, 1, 0, 0,
		0, 70, 69, 1, 0, 0, 0, 71, 79, 1, 0, 0, 0, 72, 75, 7, 0, 0, 0, 73, 76,
		3, 18, 9, 0, 74, 76, 5, 31, 0, 0, 75, 73, 1, 0, 0, 0, 75, 74, 1, 0, 0,
		0, 76, 78, 1, 0, 0, 0, 77, 72, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 13, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		82, 84, 5, 8, 0, 0, 83, 85, 3, 22, 11, 0, 84, 83, 1, 0, 0, 0, 85, 86, 1,
		0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 15, 1, 0, 0, 0, 88,
		89, 5, 9, 0, 0, 89, 93, 3, 20, 10, 0, 90, 92, 7, 2, 0, 0, 91, 90, 1, 0,
		0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 17,
		1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 97, 7, 3, 0, 0, 97, 19, 1, 0, 0, 0,
		98, 99, 7, 4, 0, 0, 99, 21, 1, 0, 0, 0, 100, 101, 6, 11, -1, 0, 101, 103,
		5, 10, 0, 0, 102, 104, 3, 22, 11, 0, 103, 102, 1, 0, 0, 0, 104, 105, 1,
		0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 1, 0, 0,
		0, 107, 108, 5, 11, 0, 0, 108, 112, 1, 0, 0, 0, 109, 112, 3, 20, 10, 0,
		110, 112, 5, 40, 0, 0, 111, 100, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111,
		110, 1, 0, 0, 0, 112, 127, 1, 0, 0, 0, 113, 114, 10, 6, 0, 0, 114, 115,
		7, 5, 0, 0, 115, 126, 3, 22, 11, 7, 116, 117, 10, 5, 0, 0, 117, 118, 7,
		6, 0, 0, 118, 126, 3, 22, 11, 6, 119, 120, 10, 4, 0, 0, 120, 121, 7, 7,
		0, 0, 121, 126, 3, 22, 11, 5, 122, 123, 10, 3, 0, 0, 123, 124, 5, 20, 0,
		0, 124, 126, 3, 22, 11, 4, 125, 113, 1, 0, 0, 0, 125, 116, 1, 0, 0, 0,
		125, 119, 1, 0, 0, 0, 125, 122, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127,
		125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 23, 1, 0, 0, 0, 129, 127, 1,
		0, 0, 0, 16, 27, 30, 33, 41, 48, 52, 59, 70, 75, 79, 86, 93, 105, 111,
		125, 127,
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
	pgqlParserEOF     = antlr.TokenEOF
	pgqlParserT__0    = 1
	pgqlParserT__1    = 2
	pgqlParserT__2    = 3
	pgqlParserT__3    = 4
	pgqlParserT__4    = 5
	pgqlParserT__5    = 6
	pgqlParserT__6    = 7
	pgqlParserT__7    = 8
	pgqlParserT__8    = 9
	pgqlParserT__9    = 10
	pgqlParserT__10   = 11
	pgqlParserT__11   = 12
	pgqlParserT__12   = 13
	pgqlParserT__13   = 14
	pgqlParserT__14   = 15
	pgqlParserT__15   = 16
	pgqlParserT__16   = 17
	pgqlParserT__17   = 18
	pgqlParserT__18   = 19
	pgqlParserT__19   = 20
	pgqlParserPLAYER  = 21
	pgqlParserTEAM    = 22
	pgqlParserDATE    = 23
	pgqlParserTYPE    = 24
	pgqlParserKILLS   = 25
	pgqlParserDAMAGE  = 26
	pgqlParserASSISTS = 27
	pgqlParserRESCUES = 28
	pgqlParserRECALLS = 29
	pgqlParserWIN     = 30
	pgqlParserGAME    = 31
	pgqlParserDESC    = 32
	pgqlParserASC     = 33
	pgqlParserRUNNING = 34
	pgqlParserAVERAGE = 35
	pgqlParserTOTAL   = 36
	pgqlParserMIN     = 37
	pgqlParserMAX     = 38
	pgqlParserSTRING  = 39
	pgqlParserNUMBER  = 40
	pgqlParserWS      = 41
)

// pgqlParser rules.
const (
	pgqlParserRULE_statement    = 0
	pgqlParserRULE_showClause   = 1
	pgqlParserRULE_showBody     = 2
	pgqlParserRULE_showFragment = 3
	pgqlParserRULE_showFunction = 4
	pgqlParserRULE_forClause    = 5
	pgqlParserRULE_byClause     = 6
	pgqlParserRULE_whereClause  = 7
	pgqlParserRULE_sortByClause = 8
	pgqlParserRULE_dimension    = 9
	pgqlParserRULE_measure      = 10
	pgqlParserRULE_expr         = 11
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ShowClause() IShowClauseContext
	ForClause() IForClauseContext
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
		p.SetState(24)
		p.ShowClause()
	}
	{
		p.SetState(25)
		p.ForClause()
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == pgqlParserT__6 {
		{
			p.SetState(26)
			p.ByClause()
		}

	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == pgqlParserT__7 {
		{
			p.SetState(29)
			p.WhereClause()
		}

	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == pgqlParserT__8 {
		{
			p.SetState(32)
			p.SortByClause()
		}

	}
	{
		p.SetState(35)
		p.Match(pgqlParserT__0)
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
	AllShowBody() []IShowBodyContext
	ShowBody(i int) IShowBodyContext

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

func (s *ShowClauseContext) AllShowBody() []IShowBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IShowBodyContext); ok {
			len++
		}
	}

	tst := make([]IShowBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IShowBodyContext); ok {
			tst[i] = t.(IShowBodyContext)
			i++
		}
	}

	return tst
}

func (s *ShowClauseContext) ShowBody(i int) IShowBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowBodyContext); ok {
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

	return t.(IShowBodyContext)
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
		p.SetState(37)
		p.Match(pgqlParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1634201502720) != 0) {
		{
			p.SetState(38)
			p.ShowBody()
		}

		p.SetState(41)
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

// IShowBodyContext is an interface to support dynamic dispatch.
type IShowBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllShowFragment() []IShowFragmentContext
	ShowFragment(i int) IShowFragmentContext

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
		p.SetState(43)
		p.ShowFragment()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == pgqlParserT__2 || _la == pgqlParserT__3 {
		{
			p.SetState(44)
			_la = p.GetTokenStream().LA(1)

			if !(_la == pgqlParserT__2 || _la == pgqlParserT__3) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(45)
			p.ShowFragment()
		}

		p.SetState(50)
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
	ShowFunction() IShowFunctionContext
	Expr() IExprContext
	STRING() antlr.TerminalNode

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

func (s *ShowFragmentContext) ShowFunction() IShowFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShowFunctionContext)
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

func (s *ShowFragmentContext) STRING() antlr.TerminalNode {
	return s.GetToken(pgqlParserSTRING, 0)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&532575944704) != 0 {
		{
			p.SetState(51)
			p.ShowFunction()
		}

	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(54)
			p.Measure()
		}

	case 2:
		{
			p.SetState(55)
			p.expr(0)
		}
		{
			p.SetState(56)
			p.Match(pgqlParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(pgqlParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IShowFunctionContext is an interface to support dynamic dispatch.
type IShowFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RUNNING() antlr.TerminalNode
	AVERAGE() antlr.TerminalNode
	TOTAL() antlr.TerminalNode
	MIN() antlr.TerminalNode
	MAX() antlr.TerminalNode

	// IsShowFunctionContext differentiates from other interfaces.
	IsShowFunctionContext()
}

type ShowFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowFunctionContext() *ShowFunctionContext {
	var p = new(ShowFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_showFunction
	return p
}

func InitEmptyShowFunctionContext(p *ShowFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pgqlParserRULE_showFunction
}

func (*ShowFunctionContext) IsShowFunctionContext() {}

func NewShowFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowFunctionContext {
	var p = new(ShowFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pgqlParserRULE_showFunction

	return p
}

func (s *ShowFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowFunctionContext) RUNNING() antlr.TerminalNode {
	return s.GetToken(pgqlParserRUNNING, 0)
}

func (s *ShowFunctionContext) AVERAGE() antlr.TerminalNode {
	return s.GetToken(pgqlParserAVERAGE, 0)
}

func (s *ShowFunctionContext) TOTAL() antlr.TerminalNode {
	return s.GetToken(pgqlParserTOTAL, 0)
}

func (s *ShowFunctionContext) MIN() antlr.TerminalNode {
	return s.GetToken(pgqlParserMIN, 0)
}

func (s *ShowFunctionContext) MAX() antlr.TerminalNode {
	return s.GetToken(pgqlParserMAX, 0)
}

func (s *ShowFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pgqlVisitor:
		return t.VisitShowFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pgqlParser) ShowFunction() (localctx IShowFunctionContext) {
	localctx = NewShowFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, pgqlParserRULE_showFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&532575944704) != 0) {
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

// IForClauseContext is an interface to support dynamic dispatch.
type IForClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Dimension() IDimensionContext
	STRING() antlr.TerminalNode

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

func (s *ForClauseContext) STRING() antlr.TerminalNode {
	return s.GetToken(pgqlParserSTRING, 0)
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
	p.EnterRule(localctx, 10, pgqlParserRULE_forClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(pgqlParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Dimension()
	}
	{
		p.SetState(65)
		p.Match(pgqlParserSTRING)
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
	p.EnterRule(localctx, 12, pgqlParserRULE_byClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(pgqlParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case pgqlParserPLAYER, pgqlParserTEAM, pgqlParserDATE, pgqlParserTYPE, pgqlParserWIN:
		{
			p.SetState(68)
			p.Dimension()
		}

	case pgqlParserGAME:
		{
			p.SetState(69)
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == pgqlParserT__2 || _la == pgqlParserT__3 {
		{
			p.SetState(72)
			_la = p.GetTokenStream().LA(1)

			if !(_la == pgqlParserT__2 || _la == pgqlParserT__3) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case pgqlParserPLAYER, pgqlParserTEAM, pgqlParserDATE, pgqlParserTYPE, pgqlParserWIN:
			{
				p.SetState(73)
				p.Dimension()
			}

		case pgqlParserGAME:
			{
				p.SetState(74)
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

		p.SetState(81)
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

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

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

func (s *WhereClauseContext) AllExpr() []IExprContext {
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

func (s *WhereClauseContext) Expr(i int) IExprContext {
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
	p.EnterRule(localctx, 14, pgqlParserRULE_whereClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(pgqlParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1101625558016) != 0) {
		{
			p.SetState(83)
			p.expr(0)
		}

		p.SetState(86)
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

// ISortByClauseContext is an interface to support dynamic dispatch.
type ISortByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Measure() IMeasureContext
	AllDESC() []antlr.TerminalNode
	DESC(i int) antlr.TerminalNode
	AllASC() []antlr.TerminalNode
	ASC(i int) antlr.TerminalNode

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

func (s *SortByClauseContext) Measure() IMeasureContext {
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

func (s *SortByClauseContext) AllDESC() []antlr.TerminalNode {
	return s.GetTokens(pgqlParserDESC)
}

func (s *SortByClauseContext) DESC(i int) antlr.TerminalNode {
	return s.GetToken(pgqlParserDESC, i)
}

func (s *SortByClauseContext) AllASC() []antlr.TerminalNode {
	return s.GetTokens(pgqlParserASC)
}

func (s *SortByClauseContext) ASC(i int) antlr.TerminalNode {
	return s.GetToken(pgqlParserASC, i)
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
	p.EnterRule(localctx, 16, pgqlParserRULE_sortByClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(pgqlParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Measure()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == pgqlParserDESC || _la == pgqlParserASC {
		{
			p.SetState(90)
			_la = p.GetTokenStream().LA(1)

			if !(_la == pgqlParserDESC || _la == pgqlParserASC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(95)
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
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1105199104) != 0) {
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
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2113929216) != 0) {
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Measure() IMeasureContext
	NUMBER() antlr.TerminalNode

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
	_startState := 22
	p.EnterRecursionRule(localctx, 22, pgqlParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case pgqlParserT__9:
		{
			p.SetState(101)
			p.Match(pgqlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1101625558016) != 0) {
			{
				p.SetState(102)
				p.expr(0)
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(107)
			p.Match(pgqlParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case pgqlParserKILLS, pgqlParserDAMAGE, pgqlParserASSISTS, pgqlParserRESCUES, pgqlParserRECALLS, pgqlParserWIN:
		{
			p.SetState(109)
			p.Measure()
		}

	case pgqlParserNUMBER:
		{
			p.SetState(110)
			p.Match(pgqlParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pgqlParserRULE_expr)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(114)
					_la = p.GetTokenStream().LA(1)

					if !(_la == pgqlParserT__11 || _la == pgqlParserT__12) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(115)
					p.expr(7)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pgqlParserRULE_expr)
				p.SetState(116)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(117)
					_la = p.GetTokenStream().LA(1)

					if !(_la == pgqlParserT__13 || _la == pgqlParserT__14) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(118)
					p.expr(6)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pgqlParserRULE_expr)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(120)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&983040) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(121)
					p.expr(5)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pgqlParserRULE_expr)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(123)
					p.Match(pgqlParserT__19)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(124)
					p.expr(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	case 11:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *pgqlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
