package main

import (
	"fmt"
)

type Token struct {
	ttype int
	text  string
}

const (
	WORD = iota
	STOP
	EOF
)

func TokenTypeName(ttype int) string {
	switch ttype {
	case WORD:
		return "WORD"
	case STOP:
		return "STOP"
	case EOF:
		return "EOF"
	}
	return "unknown token type"
}

func (t Token) String() string {
	return fmt.Sprintf("{%s %s}", TokenTypeName(t.ttype), t.text)
}
