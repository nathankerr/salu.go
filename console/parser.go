package main

import (
	"io"
	"log"
)

type Parser struct {
	lexer *Lexer
	token Token
}

func NewParser(input io.Reader) *Parser {
	p := new(Parser)
	p.lexer = NewLexer(input)
	
	p.consume() //Prime the parser
	return p
}

func (p *Parser) Parse() *Sentence{
	sen := new(Sentence)
	p.patient(sen)
	p.verb(sen)
	p.focus(sen)
	p.match(STOP)
	
	return sen
}

func (p *Parser) match(ttype int) Token{
	token := p.token
	if p.token.ttype == ttype {
		p.consume()
	} else {
		log.Fatalln("ERROR: Expecting", TokenTypeName(ttype), ", got", TokenTypeName(p.token.ttype))
	}
	return token
}

func (p *Parser) consume() {
	p.token = p.lexer.NextToken()
}

func (p *Parser) patient(sen *Sentence) {
	token := p.match(WORD)
	sen.patient = token.text
}

func (p *Parser) verb(sen *Sentence) {
	token := p.match(WORD)
	sen.verb = token.text
}

func (p *Parser) focus(sen *Sentence) {
	token := p.match(WORD)
	sen.focus = token.text
}