package main

import (
	"log"
)

type Lexer int

func (Lexer) Lex(yylval *yySymType) int {
	log.Println(yylval)
	return VAR
}

func (Lexer) Error(s string) {
	log.Println("syntax error, last name: %v", s)
}