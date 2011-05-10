package salu

import (
	"log"
	"os"
)

type verb struct {
	Handler     VerbHandler
	PatientType string
	FocusType   string
}

type VerbHandler interface {
	HandleVerb(patient, focus entity) (entity, os.Error)
}

func (s *Salu) interpretVerb(verb string, patient, focus entity) (VerbHandler, bool) {
	ptype := entityType(patient)
	log.Println("[verb] patient is a", ptype)
	ftype := entityType(focus)
	log.Println("[verb] focus is a", ftype)
	
	verbs := s.verbs[verb]
	for _, v := range verbs {
		if v.PatientType == ptype && v.FocusType == ftype {
			return v.Handler, true
		}
	}
	
	return nil, false
}
