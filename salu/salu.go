package salu

import (
	"log"
)

type Salu struct {
	entities map[string]entity
	verbs    map[string][]verb
}

func NewSalu() *Salu {
	s := new(Salu)
	s.entities = make(map[string]entity)
	s.verbs = make(map[string][]verb)

	e := make(propertyset)
	e["name"] = "Nathan"
	s.entities["user"] = e

	var sa stringadder
	var v verb
	v.PatientType = "stringliteral"
	v.FocusType = "stringliteral"
	v.Handler = sa
	s.verbs["add"] = append(s.verbs["add"], v)

	return s
}

func (s *Salu) Eval(verb, patient, focus string) {
	p := s.interpretEntity(patient)
	log.Println("[patient]:", p)

	f := s.interpretEntity(focus)
	log.Println("[focus]:", f)

	handler, ok := s.interpretVerb(verb, p, f)
	if !ok {
		log.Println("[verb] appropriate", verb, "not found")
	} else {
		result, err := handler.HandleVerb(p, f)
		if err != nil {
			log.Println("[verb] error:", err)
		} else {
			log.Println("[verb]:", result)
		}
	}
}
