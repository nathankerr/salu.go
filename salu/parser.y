%{
package salu

import (
	"fmt"
	"io"
)

type Parser struct {
	lexer *Lexer
}

var sentence *Sentence

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
	fmt.Printf("sentence %#v %#v %#v\n", $1, $2, $3)
	sentence = &Sentence{$1, $2, $3}
	return 1
}
patient: WORD
verb: WORD
focus: WORD
%%

func NewParser(input io.Reader) *Parser {
	p := new(Parser)
	p.lexer = NewLexer(input)

	return p
}

func (p *Parser) Parse() *Sentence {
	sentence = new(Sentence)
	yyParse(p.lexer)
	return sentence
}