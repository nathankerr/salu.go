package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Lexer struct {
	input io.Reader
	c     byte
	err   os.Error
	text  []byte
}

func NewLexer(input io.Reader) *Lexer {
	l := new(Lexer)
	l.input = input

	l.consume() //prime NextToken
	return l
}

// C: [tksmljnpdb]
// V: [uioe]
// N: [a]
// WS: [ \n\r\t]
// WORD: (CN)*CV
// STOP: '.'
func (l *Lexer) NextToken() Token {
	for l.err == nil {
		switch l.c {
		default:
			l.Error("WS, C, or STOP")
		case ' ', '\n', '\r', '\t': //WS
			l.consume()
		case 't', 'k', 's', 'm', 'l', 'j', 'n', 'p', 'd', 'b': //C
			return l.word()
		case '.':
			return l.tokenize(STOP)
		}
	}
	if l.err == os.EOF {
		l.c = ' '
		return l.tokenize(EOF)
	}
	panic("l.NextToken: unhandled error " + l.err.String())
}

func (l *Lexer) consume() {
	var buf [1]byte

	_, l.err = l.input.Read(buf[:])
	if l.err == nil {
		l.c = buf[0]
	}
	// fmt.Printf("%s", string(l.c))
}

// (CN)*CV
// stops at WS or STOP
func (l *Lexer) word() Token {
	l.text = append(l.text, l.c)
	l.consume()

	for l.err == nil {
		switch l.c {
		default:
			l.Error("N or V")
		case 'a': //N
			return l.wordcontinues()
		case 'u', 'i', 'o', 'e': //V
			return l.tokenize(WORD)
		}
	}
	panic("l.word: invalid token")
}

func (l *Lexer) wordcontinues() Token {
	l.text = append(l.text, l.c)
	l.consume()

	for l.err == nil {
		switch l.c {
		default:
			l.Error("C")
		case 't', 'k', 's', 'm', 'l', 'j', 'n', 'p', 'd', 'b': //C
			return l.word()
		}
	}
	panic("l.wordcontinues: invalid token")
}

func (l *Lexer) tokenize(ttype int) Token {
	l.text = append(l.text, l.c)
	t := Token{ttype, string(l.text)}
	l.text = make([]byte, 1)
	l.consume()
	log.Println(t)
	return t
}

func (l *Lexer) String() string {
	return fmt.Sprintf("%c %s", l.c, string(l.text))
}

func (l *Lexer) Error(error string) {
	log.Fatalf("ERROR: [%s][%s] %s\n", string(l.text), string(l.c), error)
}
