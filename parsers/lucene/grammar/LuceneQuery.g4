grammar LuceneQuery;

query: expression EOF;

expression: 
    expression AND expression    
    | expression OR expression   
    | '(' expression ')'        
    | fieldExpression           
    | messageTerm               
    ;

// Tag:value expressions
fieldExpression: field ':' value;

// Message terms (searches _cardinalhq.message)
messageTerm: 
    PHRASE      
    | WORD      
    ;

// Field names (tags): log.service_entity_name, _cardinalhq.level, etc.
field: IDENTIFIER ('.' IDENTIFIER)*;

// Values for tags
value: 
    PHRASE      
    | WORD      
    ;

// Lexer rules
AND: 'AND';
OR: 'OR';

IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_-]*;
PHRASE: '"' (~["\\] | '\\' .)* '"';
WORD: [a-zA-Z0-9_-]+;
WS: [ \t\r\n]+ -> skip; 