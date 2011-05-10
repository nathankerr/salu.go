package main

import (
	"os"
	"salu"
)

type stringadder int

func (s stringadder) HandleVerb(patient, focus salu.Entity) (salu.Entity, os.Error) {
	p, ok := patient.(salu.StringLiteral)
	if !ok {
		return nil, os.NewError("patient is not a StringLiteral")
	}

	f, ok := focus.(salu.StringLiteral)
	if !ok {
		return nil, os.NewError("focus is not a StringLiteral")
	}

	return salu.StringLiteral(p + " " + f), nil
}