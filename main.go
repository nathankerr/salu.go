package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"salu"
)

var s *salu.Salu

func main() {
	setup()
	console(os.Stdin, os.Stdout)
	//telnetlistener("localhost:3000")
}

func setup() {
	s = salu.NewSalu()

	// register a stringadder
	var sa stringadder
	s.RegisterVerb("malu", sa, "StringLiteral", "StringLiteral")

	// register a numberadder
	var na numberadder
	s.RegisterVerb("malu", na, "NumberLiteral", "NumberLiteral")
}

func console(input io.Reader, output io.Writer) {
	parser := salu.NewParser(input)

	for {
		fmt.Fprintf(output, "> ")
		sen := parser.Parse()
		log.Println(sen)

		result := s.Eval(sen)
		fmt.Fprintln(output, result)
	}
}

func telnetlistener (addr string) {
	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go console(c, c)
	}
}
