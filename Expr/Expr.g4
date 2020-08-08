grammar Expr;
//import Array;

program
	: stat + EOF
	;

stat:   expr                #printExpr
    |   ID '=' expr         #assign
    ;

//stat: expr                #printExpr
//	| ifStatement
//	| blockStatement
//	| assignStatement
//	| declareStatement
////	| printStatement
//	| whileStatement
//	| doWhileStatement
//	| forStatement
//	;
//ifStatement
//	: 'if' quoteExpr blockStatement
//	| 'if' quoteExpr blockStatement 'else' blockStatement
//	;
//quoteExpr
//	: '(' expr ')'
//	;
//blockStatement
//	: '{' (stat)* '}'
//	;
//assignStatement
//	: assign ';'
//	;
//declareStatement
//	: declareExpression ';'
//	;
////printStatement
////	: 'print' '(' expr ')' ';'
////	;
//whileStatement
//	: 'while' '(' expr ')' blockStatement
//	;
//doWhileStatement
//	: 'do' blockStatement 'while' '(' expr ')' ';'
//	;
//forStatement
//	: 'for' '(' commonExpression ';' expr ';' assignAbleStatement ')' blockStatement
//	;
//commonExpression
//	: declareExpression
//	| assignAbleStatement
//	;
//assignAbleStatement
//	: assign
//	| expr
//	;
//
//basicType
//	: 'int'
//	| 'string'
//	| 'bool'
//	| 'list'
//	| 'var'
//	;
//
//declarators
//	: assign (',' assign)*
//	;
//assign
//	: ID '=' expr
//	;
//primaryExpression
//	: variableExpression
//	| IntLiteral
//	| '(' expr ')'
//	;
//variableExpression
//	: ID
//	| BooleanLiteral
//	| StringLiteral
//	;
//declareExpression
//	: basicType declarators
//	;

expr:   expr 'in' expr              #listIn
    |   expr 're' expr              #re    |expr op=('*' | '/') expr    #mulDiv
    |   expr op=('+' | '-') expr    #addSub
    |   expr op=( '<'| '<='| '>='| '>') expr    #logic
    |   expr op=('==' | '!=' ) expr #equal
    |   expr op=('&&' | '||' ) expr #impact
    |   ID                          #id
    |   IntLiteral                  #intLiteral
    |   BooleanLiteral              #booleanLiteral
    |   StringLiteral               #stringLiteral
    |   List                        #list
    |   CIDR                        #cidr
    |   IP                          #ip
    |   Reg                         #reg
    |   '(' expr ')'                #parens
    ;

// 关键字
BOOLEAN : 'boolean';
INT : 'int';
STRING : 'string';
LIST : 'list';

MUL :   '*';
DIV :   '/';
ADD :   '+';
SUB :   '-';

GT : '>';
LT : '<';

LE : '<=';
GE : '>=';

EQUAL : '==';
NOTEQUAL : '!=';

AND : '&&';
OR : '||';

IN :'in';
RE :'re';

//ID  : [a-zA-Z]+;
//INT : [0-9]+;

//常量


BooleanLiteral:   'true'|'false';  //布尔值常量
IP    : IntLiteral '.' IntLiteral '.' IntLiteral '.' IntLiteral ;//
CIDR    : IntLiteral '.' IntLiteral '.' IntLiteral '.' IntLiteral '/'IntLiteral;//
IntLiteral:        [0-9]+;         //整形常量
//StringLiteral:     [a-zA-Z]+;
StringLiteral: '"' .*? '"' ;                    // 字符串常量
ID: [_a-zA-Z,.]+ ;//标识符

//List : '[' IntLiteral+ ']' ;    // [1], [1 2], [1 2 3], ...

List : '[' .*? (',' .*?)* ']' ; // must match at least one value
Reg  : '`' .*?  '`';


WS  : [ \r\n\t]+ -> skip;