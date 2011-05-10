package main

import (
	"os"
	"salu"
)

type numberadder int

func (s numberadder) HandleVerb(patient, focus salu.Entity) (salu.Entity, os.Error) {
	p, ok := patient.(salu.NumberLiteral)
	if !ok {
		return nil, os.NewError("patient is not a NumberLiteral")
	}

	f, ok := focus.(salu.NumberLiteral)
	if !ok {
		return nil, os.NewError("focus is not a NumberLiteral")
	}

	return salu.NumberLiteral(p + f), nil
}
