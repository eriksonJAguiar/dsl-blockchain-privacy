// Translator.g4
grammar Translator;

WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : select1 EOF;

select1 : 'SELECT' set_list 'FROM' from_list 'WHERE' condition;

set_list : function '(' attribute ')' ;

function : 'AVG' | 'COUNT' ;

from_list : relation;

condition : attribute '=' STRING_LITERAL ;

attribute : IDENTIFIER ;

relation : IDENTIFIER ;

IDENTIFIER :  (NUM_LETTER|'_'|('\\"'))+;


STRING_LITERAL : '"' (NUM_LETTER|'_'|('\\"'))* '"';

fragment
NUM_LETTER : ('a'..'z'|'A'..'Z'|'0'..'9');


