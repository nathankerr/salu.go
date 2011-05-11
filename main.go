package main

import (
	"bufio"
	"fmt"
	"os"
	"salu"
	"strings"
)

func main() {
	s := salu.NewSalu()

	// register a stringadder
	var sa stringadder
	s.RegisterVerb("add", sa, "StringLiteral", "StringLiteral")

	// register a numberadder
	var na numberadder
	s.RegisterVerb("add", na, "NumberLiteral", "NumberLiteral")

	//s.Eval("verb", "patient", "focus")
	//s.Eval("add", "nathan", "kerr")
	//s.Eval("add", "2", "1")

	input := bufio.NewReader(os.Stdout)

	for {
		fmt.Printf("> ")
		line, _, err := input.ReadLine()
		if err != nil {
			if err == os.EOF {
				fmt.Println()
				os.Exit(0)
			}
			panic(err)
		}

		args := strings.Split(string(line), " ", -1)
		if len(args) != 3 {
			fmt.Println("requires 3 tokens: patient verb focus")
			continue
		}
		patient := args[0]
		verb := args[1]
		focus := args[2]

		s.Eval(verb, patient, focus)
	}
}
