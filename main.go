package main

import (
	"salu"
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
	s.Eval("add", "nathan", "kerr")
	s.Eval("add", "2", "1")
}
