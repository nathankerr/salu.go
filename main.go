package main

import (
	"log"
	"net"
	"os"
	"salu"
)

var s *salu.Salu

func main() {	
	s = salu.NewSalu()

	// register a stringadder
	var sa stringadder
	s.RegisterVerb("malu", sa, "StringLiteral", "StringLiteral")

	// register a numberadder
	var na numberadder
	s.RegisterVerb("malu", na, "NumberLiteral", "NumberLiteral")

	s.Console(os.Stdin, os.Stdout)
	//telnetlistener(s, "localhost:3000")
}

func telnetlistener (s *salu.Salu, addr string) {
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

		go s.Console(c, c)
	}
}
