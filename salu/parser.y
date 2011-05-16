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

%token <text> WORD STOP NUMBER MODIFIER
%left MODIFIER

%%
sen: /*empty*/ {
	os.Exit(0)
}
| patient verb focus STOP
{
	sen = &sentence{$1, $2, $3}
	return 1
}
patient: WORD
| patient MODIFIER WORD {
	// FIXME: Pass all the data up
	$$ = $1
}
| NUMBER

verb: WORD

focus: WORD
| focus MODIFIER WORD {
	//FIXME: Pass all the data up
	$$ = $1
}
| NUMBER
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