package salu

import (
	"log"
	"os"
)

type stringadder int

func (s stringadder) HandleVerb(patient, focus entity) (entity, os.Error) {
	log.Println("[verb] Using stringadder")
	p, ok := patient.(stringliteral)
	if !ok {
		return nil, os.NewError("patient is not a string")
	}

	f, ok := focus.(stringliteral)
	if !ok {
		return nil, os.NewError("focus is not a string")
	}

	return stringliteral(p + " " + f), nil
}
