package main

import(
	"os"
)

func main() {
	lexer := NewLexer(os.Stdin)
	yyParse(lexer)
}