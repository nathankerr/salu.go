%{
package salu

import (
	"fmt"
	"io"
)

type Parser struct {
	lexer *lexer
}

var sen *sentence

%}

%union
{
	text string
}

%type <text> patient verb focus

%token <text> WORD
%token <text> STOP
%token EOF

%%
sen: patient verb focus STOP
{
	sen = &sentence{$1, $2, $3}
	return 1
}
patient: WORD
verb: WORD
focus: WORD
%%

func NewParser(input io.Reader) *Parser {
	p := new(Parser)
	p.lexer = newLexer(input)

	return p
}

func (p *Parser) Parse() *sentence {
	sen = new(sentence)
	yyParse(p.lexer)
	return sen
}