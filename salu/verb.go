package salu

import (
	"os"
)

type Verb struct {
	Handler     VerbHandler
	PatientType string
	FocusType   string
}

type VerbHandler interface {
	HandleVerb(patient, focus Entity) (Entity, os.Error)
}
