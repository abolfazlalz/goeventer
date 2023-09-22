grammar GoEventer;

parse
    : block EOF
    ;

block
    : stat*
    ;

stat
    : methodCallStat
    | defineListener
    | assignVariable
    | forStat
    | whileStat
    | defineFunction
    | returnStat
    | notifyChanStat
    ;

assignVariable
    : ID op=(ASSIGN|UPDATE) expr SCOL
    ;

notifyChanStat
    : ID NOTIFY expr SCOL
    ;

forStat
 : FOR ID ASSIGN expr COL expr statBlock
 ;

whileStat
 : WHILE expr statBlock
 ;

defineListener
    : ON methodCall statBlock
    ;

defineFunction
    : FUNC ID PARAN_OPEN functionDefineArguments PARAN_CLOSE statBlock
    ;

statBlock
 : BARACE_OPEN block BARACE_CLOSE
 | stat
 ;

returnStat
 : RETURN expr SCOL
 ;

methodCallStat
 : methodCall SCOL
 ;

methodCall
 : ID PARAN_OPEN methodCallArguments PARAN_CLOSE
 ;

methodCallArguments
 : // No arguments
 | expr (',' expr)*  // Some arguments
 ;

functionDefineArguments
 : // No arguments
 | ID (',' ID)*
 ;

expr
 : methodCall                          #methodCallExpr
 | <assoc=right> expr POW expr           #powExpr
 | MINUS expr                           #unaryMinusExpr
 | NOT expr                             #notExpr
 | expr op=(MULT | DIV | MOD) expr      #multiplicationExpr
 | expr op=(PLUS | MINUS) expr          #additiveExpr
 | expr op=(LTEQ | GTEQ | LT | GT) expr #relationalExpr
 | expr op=(EQ | NEQ) expr              #equalityExpr
 | expr AND expr                        #andExpr
 | expr OR expr                         #orExpr
 | atom                                 #atomExpr
 ;

atom
 : PARAN_OPEN expr PARAN_CLOSE #exprAtom
 | (INT | FLOAT)  #numberAtom
 | (TRUE | FALSE) #booleanAtom
 | ID             #idAtom
 | STRING         #stringAtom
 | NIL            #nilAtom
 | CHAN           #makeChanAtom
 ;

ASSIGN : ':=';
UPDATE : '=';

SCOL : ';';
COL : ':';

FOR : 'for';
WHILE : 'while';
RANGE : 'range';
FUNC : 'func';
ON : 'on';
RETURN : 'return';

CHAN : 'chan';
NOTIFY : 'notify';

STRING_TYPE : 'string';
INTEGER_TYPE : 'int';

OR : '||';
AND : '&&';
EQ : '==';
NEQ : '!=';
GT : '>';
LT : '<';
GTEQ : '>=';
LTEQ : '<=';
PLUS : '+';
MINUS : '-';
MULT : '*';
DIV : '/';
MOD : '%';
POW : '^';
NOT : '!';

TRUE : 'true';
FALSE : 'false';
NIL : 'nil';

PARAN_OPEN : '(';
PARAN_CLOSE : ')';
BARACE_OPEN : '{';
BARACE_CLOSE : '}';

ID
 : [a-zA-Z_] [a-zA-Z_0-9]*
 ;

INT
 : [0-9]+
 ;

FLOAT
 : [0-9]+ '.' [0-9]*
 | '.' [0-9]+
 ;

STRING
 : '"' (~["\r\n] | '""')* '"'
 ;


SPACE
 : [ \t\r\n] -> skip
 ;