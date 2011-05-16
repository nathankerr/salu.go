%{
package salu

import (
	"fmt"
	"io"
	"os"
)

type Parser struct {
	lexer Lexer
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
%token <text> NUMBER

%%
sen: /*empty*/ {
	os.Exit(0)
}
| patient verb focus STOP
{
	sen = &sentence{$1, $2, $3}
	return 1
}
patient: WORD | NUMBER
verb: WORD
focus: WORD | NUMBER
%%

func NewParser(input io.Reader) *Parser {
	p := new(Parser)
	p.lexer = NewLexer(input)

	return p
}

func (p *Parser) Parse() *sentence {
	sen = new(sentence)
	yyParse(p.lexer)
	return sen
}