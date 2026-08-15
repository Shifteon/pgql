grammar pgql;

/* Parser rules */

statement : showClause forClause byClause? whereClause? sortByClause? ';' ;

/* Clauses */
showClause :  'show' showBody+ ;
/* Should I support things like max and min in the show clause and just in general? */
showBody : showFragment (('and' | ',') showFragment)* ;
showFragment : showFunction? (measure | (expr 'as' STRING)) ;
showFunction : RUNNING | AVERAGE | TOTAL | MIN | MAX ;

forClause : 'for' dimension STRING;
byClause : 'by' (dimension | GAME) (('and' | ',') (dimension | GAME))* ;
whereClause :  'where' ((AVERAGE | TOTAL | MIN | MAX)* expr) (('and' | ',') (AVERAGE | TOTAL | MIN | MAX)* expr)* ;
sortByClause : 'sort by' measure (DESC | ASC)* ;

dimension : PLAYER | TEAM | DATE | TYPE | WIN ;
measure : KILLS | DAMAGE | ASSISTS | RESCUES | RECALLS | WIN ;

expr : '(' expr+ ')'
     | expr ('*' | '/') expr
     | expr ('+' | '-') expr
     | expr ('<' | '>' | '<=' | '>=') expr
     | expr '=' expr
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

/* Modifiers */
DESC : 'desc' ;
ASC : 'asc' ;
RUNNING : 'running' ;
AVERAGE : 'average' ;
TOTAL : 'total' ;
MIN : 'min' ;
MAX : 'max' ;

/* primitives */
STRING : '"' ~["\r\n]* '"' ;
fragment INT : [0-9]+ ;
fragment FLOAT : [0-9]+ '.' [0-9]* 
      | '.' [0-9]+ 
      | [0-9]+ '.'? [0-9]* [eE] [+-]? [0-9]+ 
      | '.' [0-9]+ [eE] [+-]? [0-9]+
      ;
NUMBER : INT | FLOAT ;

/* Handle whitespace */
WS : [ \t\r\n]+ -> skip ;
