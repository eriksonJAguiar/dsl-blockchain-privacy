// Translator.g4
grammar Translator;

WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : select2 EOF
    | insert1 EOF;

insert1 : 'INSERT INTO'  table_name '(' columns_list ')' 'VALUES' '(' values_list ')';

columns_list : column ( ',' column)*;

values_list : value ( ',' value)*;

select2 : select1 ('EPSILON' '=' NUM_FLOAT)?;

select1 : 'SELECT' set_list 'FROM' from_list 'WHERE' condition;

set_list : pair ( ',' pair )* ;

pair : function '(' attribute ')' ;

function : 'AVG' | 'HISTOGRAM' ;

from_list : relation;

condition : attribute '=' STRING_LITERAL ;

column : IDENTIFIER;

value : STRING_LITERAL;

table_name : IDENTIFIER;

attribute : IDENTIFIER ;

relation : IDENTIFIER ;

IDENTIFIER :  (NUM_LETTER|'_'|('\\"'))+;


STRING_LITERAL : '"' (NUM_LETTER|'_'|('\\"'))* '"';


NUM_FLOAT : ('0'..'9')+ ('.' ('0'..'9')*)?;


fragment
NUM_LETTER : ('a'..'z'|'A'..'Z'|'0'..'9');

