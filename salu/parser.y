%{
package salu

import (
	"fmt"
	"io"
)

type Parser struct {
	lexer *Lexer
}

%}

%union
{
	text string
}

%token <text> WORD
%token <text> STOP
%token EOF

%%
sen: WORD STOP
{
	fmt.Println("sentence", $1, $2)
	return 1
}
%%

func NewParser(input io.Reader) *Parser {
	p := new(Parser)
	p.lexer = NewLexer(input)

	return p
}

func (p *Parser) Parse() *Sentence {
	sen := new(Sentence)
	yyParse(p.lexer)
	return sen
}