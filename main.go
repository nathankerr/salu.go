package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"salu"
	"strings"
)

var s *salu.Salu

func main() {
	setup()
	//console(os.Stdin, os.Stdout)
	telnetlistener("localhost:3000")
}

func setup() {
	s = salu.NewSalu()

	// register a stringadder
	var sa stringadder
	s.RegisterVerb("add", sa, "StringLiteral", "StringLiteral")

	// register a numberadder
	var na numberadder
	s.RegisterVerb("add", na, "NumberLiteral", "NumberLiteral")
}

func console(input io.Reader, output io.Writer) {
	buffer := bufio.NewReader(input)

	for {
		fmt.Fprintf(output, "> ")
		line, _, err := buffer.ReadLine()
		if err != nil {
			if err == os.EOF {
				fmt.Println()
				os.Exit(0)
			}
			panic(err)
		}

		args := strings.Split(string(line), " ", -1)
		if len(args) != 3 {
			fmt.Fprintln(output, "requires 3 tokens: patient verb focus")
			continue
		}
		patient := args[0]
		verb := args[1]
		focus := args[2]

		result := s.Eval(verb, patient, focus)
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
