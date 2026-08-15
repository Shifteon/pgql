// Generated from pgql.g4 by ANTLR 4.13.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link pgqlParser}.
 */
public interface pgqlListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link pgqlParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(pgqlParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(pgqlParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#showClause}.
	 * @param ctx the parse tree
	 */
	void enterShowClause(pgqlParser.ShowClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#showClause}.
	 * @param ctx the parse tree
	 */
	void exitShowClause(pgqlParser.ShowClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#showBody}.
	 * @param ctx the parse tree
	 */
	void enterShowBody(pgqlParser.ShowBodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#showBody}.
	 * @param ctx the parse tree
	 */
	void exitShowBody(pgqlParser.ShowBodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#showFragment}.
	 * @param ctx the parse tree
	 */
	void enterShowFragment(pgqlParser.ShowFragmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#showFragment}.
	 * @param ctx the parse tree
	 */
	void exitShowFragment(pgqlParser.ShowFragmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#showFunction}.
	 * @param ctx the parse tree
	 */
	void enterShowFunction(pgqlParser.ShowFunctionContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#showFunction}.
	 * @param ctx the parse tree
	 */
	void exitShowFunction(pgqlParser.ShowFunctionContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#forClause}.
	 * @param ctx the parse tree
	 */
	void enterForClause(pgqlParser.ForClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#forClause}.
	 * @param ctx the parse tree
	 */
	void exitForClause(pgqlParser.ForClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#byClause}.
	 * @param ctx the parse tree
	 */
	void enterByClause(pgqlParser.ByClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#byClause}.
	 * @param ctx the parse tree
	 */
	void exitByClause(pgqlParser.ByClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#whereClause}.
	 * @param ctx the parse tree
	 */
	void enterWhereClause(pgqlParser.WhereClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#whereClause}.
	 * @param ctx the parse tree
	 */
	void exitWhereClause(pgqlParser.WhereClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#sortByClause}.
	 * @param ctx the parse tree
	 */
	void enterSortByClause(pgqlParser.SortByClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#sortByClause}.
	 * @param ctx the parse tree
	 */
	void exitSortByClause(pgqlParser.SortByClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#dimension}.
	 * @param ctx the parse tree
	 */
	void enterDimension(pgqlParser.DimensionContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#dimension}.
	 * @param ctx the parse tree
	 */
	void exitDimension(pgqlParser.DimensionContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#measure}.
	 * @param ctx the parse tree
	 */
	void enterMeasure(pgqlParser.MeasureContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#measure}.
	 * @param ctx the parse tree
	 */
	void exitMeasure(pgqlParser.MeasureContext ctx);
	/**
	 * Enter a parse tree produced by {@link pgqlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(pgqlParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link pgqlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(pgqlParser.ExprContext ctx);
}