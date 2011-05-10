package main

import (
	"salu"
)

func main() {
	s := salu.NewSalu()
	//s.Eval("verb", "patient", "focus")
	s.Eval("add", "user", "nathan")
	// s.Eval("add", "2", "1")
	// s.Eval("add", "test", "string")
}