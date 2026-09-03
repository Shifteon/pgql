grammar pgql;

/* Parser rules */

statement : showClause forClause byClause? whereClause? sortByClause? EOF ;

/* Clauses */
showClause :  'show' sequentialFunction? showBody ;
showBody : showFragment (COMMA showFragment)* ;
showFragment : measure | aggregateFunction | specificity | (LPAREN (expr | predicate) RPAREN IDENTIFIER) ;

forClause : 'for' dimension IDENTIFIER;
byClause : 'by' (dimension | GAME) (COMMA (dimension | GAME))? ;
whereClause : 'where' predicate ;
sortByClause : 'sort by' sortBody ;
sortBody : expr (DESC | ASC)? (COMMA sortBody)* ;

dimension : PLAYER | TEAM | DATE | TYPE | WIN ;
measure : KILLS | DAMAGE | ASSISTS | RESCUES | RECALLS | WIN | TOTALLOSES | TOTALWINS | GAMECOUNT ;
aggregateFunction : (AVERAGE | TOTAL | MAX | MIN) measure ;
sequentialFunction : RUNNING ;
logicalOperator : LOGICALAND | LOGICALOR ;
comparisonOperator : LESSER | GREATER | LESSEREQUAL | GREATEREQUAL | EQUAL | NOTEQUAL ;

predicate : LPAREN predicate+ RPAREN
          | predicate logicalOperator predicate
          | expr comparisonOperator expr ;
specificity : IDENTIFIER':'(measure | aggregateFunction) ;

expr : LPAREN expr+ RPAREN
     | expr (MULTIPLY | DIVIDE) expr
     | expr (SUM | DIFFERENCE) expr
     | aggregateFunction
     | specificity
     | IDENTIFIER
     | measure
     | NUMBER
     ;

/* Lexer rules */

/* Keywords */

/* Dimensions and measures */
PLAYER : 'player' ;
TEAM : 'team' ;
DATE : 'date' ;
TYPE : 'type' ;
KILLS : 'kills';
DAMAGE : 'damage' ;
ASSISTS : 'assists' ;
RESCUES : 'rescues' ;
RECALLS : 'recalls' ;
WIN : 'win' ;
GAME : 'game' ;
TOTALWINS : 'totalWins' ;
TOTALLOSES : 'totalLoses' ;
GAMECOUNT : 'gameCount' ;

/* Modifiers */
DESC : 'desc' ;
ASC : 'asc' ;
RUNNING : 'running' ;
AVERAGE : 'average' ;
TOTAL : 'total' ;
MIN : 'min' ;
MAX : 'max' ;

/* Operators */
LESSEREQUAL : '<=' ;
GREATEREQUAL : '>=' ;
GREATER : '>' ;
LESSER : '<' ;
EQUAL : '=' ;
NOTEQUAL : '!=' ;
SUM : '+' ;
DIFFERENCE : '-' ;
MULTIPLY : '*' ;
DIVIDE : '/' ;

LOGICALAND : 'and' ;
LOGICALOR : 'or' ;

LPAREN : '(' ;
RPAREN : ')' ;
COMMA : ',' ;

/* primitives */
STRING : '"' ~["\r\n]* '"' ;
fragment INT : [0-9]+ ;
fragment FLOAT : [0-9]+ '.' [0-9]* 
      | '.' [0-9]+ 
      | [0-9]+ '.'? [0-9]* [eE] [+-]? [0-9]+ 
      | '.' [0-9]+ [eE] [+-]? [0-9]+
      ;
NUMBER : INT | FLOAT ;
IDENTIFIER : [a-zA-Z]+ ;

/* Handle whitespace */
WS : [ \t\r\n]+ -> skip ;
