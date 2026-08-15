// Generated from pgql.g4 by ANTLR 4.13.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class pgqlParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, PLAYER=21, TEAM=22, DATE=23, TYPE=24, KILLS=25, 
		DAMAGE=26, ASSISTS=27, RESCUES=28, RECALLS=29, WIN=30, GAME=31, DESC=32, 
		ASC=33, RUNNING=34, AVERAGE=35, TOTAL=36, MIN=37, MAX=38, STRING=39, NUMBER=40, 
		WS=41;
	public static final int
		RULE_statement = 0, RULE_showClause = 1, RULE_showBody = 2, RULE_showFragment = 3, 
		RULE_showFunction = 4, RULE_forClause = 5, RULE_byClause = 6, RULE_whereClause = 7, 
		RULE_sortByClause = 8, RULE_dimension = 9, RULE_measure = 10, RULE_expr = 11;
	private static String[] makeRuleNames() {
		return new String[] {
			"statement", "showClause", "showBody", "showFragment", "showFunction", 
			"forClause", "byClause", "whereClause", "sortByClause", "dimension", 
			"measure", "expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'show'", "'and'", "','", "'as'", "'for'", "'by'", "'where'", 
			"'sort by'", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'<'", "'>'", 
			"'<='", "'>='", "'='", "'player'", "'team'", "'date'", "'type'", "'kills'", 
			"'damage'", "'assists'", "'rescues'", "'recalls'", "'win'", "'game'", 
			"'desc'", "'asc'", "'running'", "'average'", "'total'", "'min'", "'max'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, "PLAYER", "TEAM", 
			"DATE", "TYPE", "KILLS", "DAMAGE", "ASSISTS", "RESCUES", "RECALLS", "WIN", 
			"GAME", "DESC", "ASC", "RUNNING", "AVERAGE", "TOTAL", "MIN", "MAX", "STRING", 
			"NUMBER", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "pgql.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public pgqlParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public ShowClauseContext showClause() {
			return getRuleContext(ShowClauseContext.class,0);
		}
		public ForClauseContext forClause() {
			return getRuleContext(ForClauseContext.class,0);
		}
		public ByClauseContext byClause() {
			return getRuleContext(ByClauseContext.class,0);
		}
		public WhereClauseContext whereClause() {
			return getRuleContext(WhereClauseContext.class,0);
		}
		public SortByClauseContext sortByClause() {
			return getRuleContext(SortByClauseContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitStatement(this);
		}
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24);
			showClause();
			setState(25);
			forClause();
			setState(27);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(26);
				byClause();
				}
			}

			setState(30);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(29);
				whereClause();
				}
			}

			setState(33);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__8) {
				{
				setState(32);
				sortByClause();
				}
			}

			setState(35);
			match(T__0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ShowClauseContext extends ParserRuleContext {
		public List<ShowBodyContext> showBody() {
			return getRuleContexts(ShowBodyContext.class);
		}
		public ShowBodyContext showBody(int i) {
			return getRuleContext(ShowBodyContext.class,i);
		}
		public ShowClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_showClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterShowClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitShowClause(this);
		}
	}

	public final ShowClauseContext showClause() throws RecognitionException {
		ShowClauseContext _localctx = new ShowClauseContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_showClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(37);
			match(T__1);
			setState(39); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(38);
				showBody();
				}
				}
				setState(41); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 1634201502720L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ShowBodyContext extends ParserRuleContext {
		public List<ShowFragmentContext> showFragment() {
			return getRuleContexts(ShowFragmentContext.class);
		}
		public ShowFragmentContext showFragment(int i) {
			return getRuleContext(ShowFragmentContext.class,i);
		}
		public ShowBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_showBody; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterShowBody(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitShowBody(this);
		}
	}

	public final ShowBodyContext showBody() throws RecognitionException {
		ShowBodyContext _localctx = new ShowBodyContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_showBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			showFragment();
			setState(48);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2 || _la==T__3) {
				{
				{
				setState(44);
				_la = _input.LA(1);
				if ( !(_la==T__2 || _la==T__3) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(45);
				showFragment();
				}
				}
				setState(50);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ShowFragmentContext extends ParserRuleContext {
		public MeasureContext measure() {
			return getRuleContext(MeasureContext.class,0);
		}
		public ShowFunctionContext showFunction() {
			return getRuleContext(ShowFunctionContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode STRING() { return getToken(pgqlParser.STRING, 0); }
		public ShowFragmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_showFragment; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterShowFragment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitShowFragment(this);
		}
	}

	public final ShowFragmentContext showFragment() throws RecognitionException {
		ShowFragmentContext _localctx = new ShowFragmentContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_showFragment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 532575944704L) != 0)) {
				{
				setState(51);
				showFunction();
				}
			}

			setState(59);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				setState(54);
				measure();
				}
				break;
			case 2:
				{
				{
				setState(55);
				expr(0);
				setState(56);
				match(T__4);
				setState(57);
				match(STRING);
				}
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ShowFunctionContext extends ParserRuleContext {
		public TerminalNode RUNNING() { return getToken(pgqlParser.RUNNING, 0); }
		public TerminalNode AVERAGE() { return getToken(pgqlParser.AVERAGE, 0); }
		public TerminalNode TOTAL() { return getToken(pgqlParser.TOTAL, 0); }
		public TerminalNode MIN() { return getToken(pgqlParser.MIN, 0); }
		public TerminalNode MAX() { return getToken(pgqlParser.MAX, 0); }
		public ShowFunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_showFunction; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterShowFunction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitShowFunction(this);
		}
	}

	public final ShowFunctionContext showFunction() throws RecognitionException {
		ShowFunctionContext _localctx = new ShowFunctionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_showFunction);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(61);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 532575944704L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ForClauseContext extends ParserRuleContext {
		public DimensionContext dimension() {
			return getRuleContext(DimensionContext.class,0);
		}
		public TerminalNode STRING() { return getToken(pgqlParser.STRING, 0); }
		public ForClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterForClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitForClause(this);
		}
	}

	public final ForClauseContext forClause() throws RecognitionException {
		ForClauseContext _localctx = new ForClauseContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_forClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(63);
			match(T__5);
			setState(64);
			dimension();
			setState(65);
			match(STRING);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ByClauseContext extends ParserRuleContext {
		public List<DimensionContext> dimension() {
			return getRuleContexts(DimensionContext.class);
		}
		public DimensionContext dimension(int i) {
			return getRuleContext(DimensionContext.class,i);
		}
		public List<TerminalNode> GAME() { return getTokens(pgqlParser.GAME); }
		public TerminalNode GAME(int i) {
			return getToken(pgqlParser.GAME, i);
		}
		public ByClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_byClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterByClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitByClause(this);
		}
	}

	public final ByClauseContext byClause() throws RecognitionException {
		ByClauseContext _localctx = new ByClauseContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_byClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(67);
			match(T__6);
			setState(70);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PLAYER:
			case TEAM:
			case DATE:
			case TYPE:
			case WIN:
				{
				setState(68);
				dimension();
				}
				break;
			case GAME:
				{
				setState(69);
				match(GAME);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(79);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2 || _la==T__3) {
				{
				{
				setState(72);
				_la = _input.LA(1);
				if ( !(_la==T__2 || _la==T__3) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(75);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case PLAYER:
				case TEAM:
				case DATE:
				case TYPE:
				case WIN:
					{
					setState(73);
					dimension();
					}
					break;
				case GAME:
					{
					setState(74);
					match(GAME);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				}
				setState(81);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class WhereClauseContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> AVERAGE() { return getTokens(pgqlParser.AVERAGE); }
		public TerminalNode AVERAGE(int i) {
			return getToken(pgqlParser.AVERAGE, i);
		}
		public List<TerminalNode> TOTAL() { return getTokens(pgqlParser.TOTAL); }
		public TerminalNode TOTAL(int i) {
			return getToken(pgqlParser.TOTAL, i);
		}
		public List<TerminalNode> MIN() { return getTokens(pgqlParser.MIN); }
		public TerminalNode MIN(int i) {
			return getToken(pgqlParser.MIN, i);
		}
		public List<TerminalNode> MAX() { return getTokens(pgqlParser.MAX); }
		public TerminalNode MAX(int i) {
			return getToken(pgqlParser.MAX, i);
		}
		public WhereClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whereClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterWhereClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitWhereClause(this);
		}
	}

	public final WhereClauseContext whereClause() throws RecognitionException {
		WhereClauseContext _localctx = new WhereClauseContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_whereClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
			match(T__7);
			{
			setState(86);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 515396075520L) != 0)) {
				{
				{
				setState(83);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 515396075520L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				}
				setState(88);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(89);
			expr(0);
			}
			setState(101);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2 || _la==T__3) {
				{
				{
				setState(91);
				_la = _input.LA(1);
				if ( !(_la==T__2 || _la==T__3) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(95);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 515396075520L) != 0)) {
					{
					{
					setState(92);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 515396075520L) != 0)) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					}
					setState(97);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(98);
				expr(0);
				}
				}
				setState(103);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SortByClauseContext extends ParserRuleContext {
		public MeasureContext measure() {
			return getRuleContext(MeasureContext.class,0);
		}
		public List<TerminalNode> DESC() { return getTokens(pgqlParser.DESC); }
		public TerminalNode DESC(int i) {
			return getToken(pgqlParser.DESC, i);
		}
		public List<TerminalNode> ASC() { return getTokens(pgqlParser.ASC); }
		public TerminalNode ASC(int i) {
			return getToken(pgqlParser.ASC, i);
		}
		public SortByClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sortByClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterSortByClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitSortByClause(this);
		}
	}

	public final SortByClauseContext sortByClause() throws RecognitionException {
		SortByClauseContext _localctx = new SortByClauseContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_sortByClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(104);
			match(T__8);
			setState(105);
			measure();
			setState(109);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DESC || _la==ASC) {
				{
				{
				setState(106);
				_la = _input.LA(1);
				if ( !(_la==DESC || _la==ASC) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				}
				setState(111);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DimensionContext extends ParserRuleContext {
		public TerminalNode PLAYER() { return getToken(pgqlParser.PLAYER, 0); }
		public TerminalNode TEAM() { return getToken(pgqlParser.TEAM, 0); }
		public TerminalNode DATE() { return getToken(pgqlParser.DATE, 0); }
		public TerminalNode TYPE() { return getToken(pgqlParser.TYPE, 0); }
		public TerminalNode WIN() { return getToken(pgqlParser.WIN, 0); }
		public DimensionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dimension; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterDimension(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitDimension(this);
		}
	}

	public final DimensionContext dimension() throws RecognitionException {
		DimensionContext _localctx = new DimensionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_dimension);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1105199104L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MeasureContext extends ParserRuleContext {
		public TerminalNode KILLS() { return getToken(pgqlParser.KILLS, 0); }
		public TerminalNode DAMAGE() { return getToken(pgqlParser.DAMAGE, 0); }
		public TerminalNode ASSISTS() { return getToken(pgqlParser.ASSISTS, 0); }
		public TerminalNode RESCUES() { return getToken(pgqlParser.RESCUES, 0); }
		public TerminalNode RECALLS() { return getToken(pgqlParser.RECALLS, 0); }
		public TerminalNode WIN() { return getToken(pgqlParser.WIN, 0); }
		public MeasureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_measure; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterMeasure(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitMeasure(this);
		}
	}

	public final MeasureContext measure() throws RecognitionException {
		MeasureContext _localctx = new MeasureContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_measure);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2113929216L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public MeasureContext measure() {
			return getRuleContext(MeasureContext.class,0);
		}
		public TerminalNode NUMBER() { return getToken(pgqlParser.NUMBER, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).enterExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof pgqlListener ) ((pgqlListener)listener).exitExpr(this);
		}
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 22;
		enterRecursionRule(_localctx, 22, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(127);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__9:
				{
				setState(117);
				match(T__9);
				setState(119); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(118);
					expr(0);
					}
					}
					setState(121); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 1101625558016L) != 0) );
				setState(123);
				match(T__10);
				}
				break;
			case KILLS:
			case DAMAGE:
			case ASSISTS:
			case RESCUES:
			case RECALLS:
			case WIN:
				{
				setState(125);
				measure();
				}
				break;
			case NUMBER:
				{
				setState(126);
				match(NUMBER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(143);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(141);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(129);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(130);
						_la = _input.LA(1);
						if ( !(_la==T__11 || _la==T__12) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(131);
						expr(7);
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(132);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(133);
						_la = _input.LA(1);
						if ( !(_la==T__13 || _la==T__14) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(134);
						expr(6);
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(135);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(136);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 983040L) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(137);
						expr(5);
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(138);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(139);
						match(T__19);
						setState(140);
						expr(4);
						}
						break;
					}
					} 
				}
				setState(145);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 11:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		case 2:
			return precpred(_ctx, 4);
		case 3:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001)\u0093\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0003\u0000\u001c\b\u0000\u0001\u0000\u0003"+
		"\u0000\u001f\b\u0000\u0001\u0000\u0003\u0000\"\b\u0000\u0001\u0000\u0001"+
		"\u0000\u0001\u0001\u0001\u0001\u0004\u0001(\b\u0001\u000b\u0001\f\u0001"+
		")\u0001\u0002\u0001\u0002\u0001\u0002\u0005\u0002/\b\u0002\n\u0002\f\u0002"+
		"2\t\u0002\u0001\u0003\u0003\u00035\b\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0003\u0003<\b\u0003\u0001\u0004\u0001"+
		"\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0003\u0006G\b\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0003\u0006L\b\u0006\u0005\u0006N\b\u0006\n\u0006\f\u0006Q\t\u0006"+
		"\u0001\u0007\u0001\u0007\u0005\u0007U\b\u0007\n\u0007\f\u0007X\t\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007^\b\u0007"+
		"\n\u0007\f\u0007a\t\u0007\u0001\u0007\u0005\u0007d\b\u0007\n\u0007\f\u0007"+
		"g\t\u0007\u0001\b\u0001\b\u0001\b\u0005\bl\b\b\n\b\f\bo\t\b\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0004\u000bx\b"+
		"\u000b\u000b\u000b\f\u000by\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0003\u000b\u0080\b\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u008e\b\u000b\n\u000b\f\u000b"+
		"\u0091\t\u000b\u0001\u000b\u0000\u0001\u0016\f\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0000\t\u0001\u0000\u0003\u0004\u0001"+
		"\u0000\"&\u0001\u0000#&\u0001\u0000 !\u0002\u0000\u0015\u0018\u001e\u001e"+
		"\u0001\u0000\u0019\u001e\u0001\u0000\f\r\u0001\u0000\u000e\u000f\u0001"+
		"\u0000\u0010\u0013\u009b\u0000\u0018\u0001\u0000\u0000\u0000\u0002%\u0001"+
		"\u0000\u0000\u0000\u0004+\u0001\u0000\u0000\u0000\u00064\u0001\u0000\u0000"+
		"\u0000\b=\u0001\u0000\u0000\u0000\n?\u0001\u0000\u0000\u0000\fC\u0001"+
		"\u0000\u0000\u0000\u000eR\u0001\u0000\u0000\u0000\u0010h\u0001\u0000\u0000"+
		"\u0000\u0012p\u0001\u0000\u0000\u0000\u0014r\u0001\u0000\u0000\u0000\u0016"+
		"\u007f\u0001\u0000\u0000\u0000\u0018\u0019\u0003\u0002\u0001\u0000\u0019"+
		"\u001b\u0003\n\u0005\u0000\u001a\u001c\u0003\f\u0006\u0000\u001b\u001a"+
		"\u0001\u0000\u0000\u0000\u001b\u001c\u0001\u0000\u0000\u0000\u001c\u001e"+
		"\u0001\u0000\u0000\u0000\u001d\u001f\u0003\u000e\u0007\u0000\u001e\u001d"+
		"\u0001\u0000\u0000\u0000\u001e\u001f\u0001\u0000\u0000\u0000\u001f!\u0001"+
		"\u0000\u0000\u0000 \"\u0003\u0010\b\u0000! \u0001\u0000\u0000\u0000!\""+
		"\u0001\u0000\u0000\u0000\"#\u0001\u0000\u0000\u0000#$\u0005\u0001\u0000"+
		"\u0000$\u0001\u0001\u0000\u0000\u0000%\'\u0005\u0002\u0000\u0000&(\u0003"+
		"\u0004\u0002\u0000\'&\u0001\u0000\u0000\u0000()\u0001\u0000\u0000\u0000"+
		")\'\u0001\u0000\u0000\u0000)*\u0001\u0000\u0000\u0000*\u0003\u0001\u0000"+
		"\u0000\u0000+0\u0003\u0006\u0003\u0000,-\u0007\u0000\u0000\u0000-/\u0003"+
		"\u0006\u0003\u0000.,\u0001\u0000\u0000\u0000/2\u0001\u0000\u0000\u0000"+
		"0.\u0001\u0000\u0000\u000001\u0001\u0000\u0000\u00001\u0005\u0001\u0000"+
		"\u0000\u000020\u0001\u0000\u0000\u000035\u0003\b\u0004\u000043\u0001\u0000"+
		"\u0000\u000045\u0001\u0000\u0000\u00005;\u0001\u0000\u0000\u00006<\u0003"+
		"\u0014\n\u000078\u0003\u0016\u000b\u000089\u0005\u0005\u0000\u00009:\u0005"+
		"\'\u0000\u0000:<\u0001\u0000\u0000\u0000;6\u0001\u0000\u0000\u0000;7\u0001"+
		"\u0000\u0000\u0000<\u0007\u0001\u0000\u0000\u0000=>\u0007\u0001\u0000"+
		"\u0000>\t\u0001\u0000\u0000\u0000?@\u0005\u0006\u0000\u0000@A\u0003\u0012"+
		"\t\u0000AB\u0005\'\u0000\u0000B\u000b\u0001\u0000\u0000\u0000CF\u0005"+
		"\u0007\u0000\u0000DG\u0003\u0012\t\u0000EG\u0005\u001f\u0000\u0000FD\u0001"+
		"\u0000\u0000\u0000FE\u0001\u0000\u0000\u0000GO\u0001\u0000\u0000\u0000"+
		"HK\u0007\u0000\u0000\u0000IL\u0003\u0012\t\u0000JL\u0005\u001f\u0000\u0000"+
		"KI\u0001\u0000\u0000\u0000KJ\u0001\u0000\u0000\u0000LN\u0001\u0000\u0000"+
		"\u0000MH\u0001\u0000\u0000\u0000NQ\u0001\u0000\u0000\u0000OM\u0001\u0000"+
		"\u0000\u0000OP\u0001\u0000\u0000\u0000P\r\u0001\u0000\u0000\u0000QO\u0001"+
		"\u0000\u0000\u0000RV\u0005\b\u0000\u0000SU\u0007\u0002\u0000\u0000TS\u0001"+
		"\u0000\u0000\u0000UX\u0001\u0000\u0000\u0000VT\u0001\u0000\u0000\u0000"+
		"VW\u0001\u0000\u0000\u0000WY\u0001\u0000\u0000\u0000XV\u0001\u0000\u0000"+
		"\u0000YZ\u0003\u0016\u000b\u0000Ze\u0001\u0000\u0000\u0000[_\u0007\u0000"+
		"\u0000\u0000\\^\u0007\u0002\u0000\u0000]\\\u0001\u0000\u0000\u0000^a\u0001"+
		"\u0000\u0000\u0000_]\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000"+
		"`b\u0001\u0000\u0000\u0000a_\u0001\u0000\u0000\u0000bd\u0003\u0016\u000b"+
		"\u0000c[\u0001\u0000\u0000\u0000dg\u0001\u0000\u0000\u0000ec\u0001\u0000"+
		"\u0000\u0000ef\u0001\u0000\u0000\u0000f\u000f\u0001\u0000\u0000\u0000"+
		"ge\u0001\u0000\u0000\u0000hi\u0005\t\u0000\u0000im\u0003\u0014\n\u0000"+
		"jl\u0007\u0003\u0000\u0000kj\u0001\u0000\u0000\u0000lo\u0001\u0000\u0000"+
		"\u0000mk\u0001\u0000\u0000\u0000mn\u0001\u0000\u0000\u0000n\u0011\u0001"+
		"\u0000\u0000\u0000om\u0001\u0000\u0000\u0000pq\u0007\u0004\u0000\u0000"+
		"q\u0013\u0001\u0000\u0000\u0000rs\u0007\u0005\u0000\u0000s\u0015\u0001"+
		"\u0000\u0000\u0000tu\u0006\u000b\uffff\uffff\u0000uw\u0005\n\u0000\u0000"+
		"vx\u0003\u0016\u000b\u0000wv\u0001\u0000\u0000\u0000xy\u0001\u0000\u0000"+
		"\u0000yw\u0001\u0000\u0000\u0000yz\u0001\u0000\u0000\u0000z{\u0001\u0000"+
		"\u0000\u0000{|\u0005\u000b\u0000\u0000|\u0080\u0001\u0000\u0000\u0000"+
		"}\u0080\u0003\u0014\n\u0000~\u0080\u0005(\u0000\u0000\u007ft\u0001\u0000"+
		"\u0000\u0000\u007f}\u0001\u0000\u0000\u0000\u007f~\u0001\u0000\u0000\u0000"+
		"\u0080\u008f\u0001\u0000\u0000\u0000\u0081\u0082\n\u0006\u0000\u0000\u0082"+
		"\u0083\u0007\u0006\u0000\u0000\u0083\u008e\u0003\u0016\u000b\u0007\u0084"+
		"\u0085\n\u0005\u0000\u0000\u0085\u0086\u0007\u0007\u0000\u0000\u0086\u008e"+
		"\u0003\u0016\u000b\u0006\u0087\u0088\n\u0004\u0000\u0000\u0088\u0089\u0007"+
		"\b\u0000\u0000\u0089\u008e\u0003\u0016\u000b\u0005\u008a\u008b\n\u0003"+
		"\u0000\u0000\u008b\u008c\u0005\u0014\u0000\u0000\u008c\u008e\u0003\u0016"+
		"\u000b\u0004\u008d\u0081\u0001\u0000\u0000\u0000\u008d\u0084\u0001\u0000"+
		"\u0000\u0000\u008d\u0087\u0001\u0000\u0000\u0000\u008d\u008a\u0001\u0000"+
		"\u0000\u0000\u008e\u0091\u0001\u0000\u0000\u0000\u008f\u008d\u0001\u0000"+
		"\u0000\u0000\u008f\u0090\u0001\u0000\u0000\u0000\u0090\u0017\u0001\u0000"+
		"\u0000\u0000\u0091\u008f\u0001\u0000\u0000\u0000\u0012\u001b\u001e!)0"+
		"4;FKOV_emy\u007f\u008d\u008f";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}