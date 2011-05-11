package salu

import (
	"fmt"
	"strconv"
)

type Salu struct {
	entities map[string]Entity
	verbs    map[string][]Verb
}

func NewSalu() *Salu {
	s := new(Salu)
	s.entities = make(map[string]Entity)
	s.verbs = make(map[string][]Verb)

	e := make(PropertySet)
	e["name"] = "Nathan"
	s.entities["user"] = e

	return s
}

func (s *Salu) Eval(verb, patient, focus string) {
	p := s.interpretEntity(patient)
	f := s.interpretEntity(focus)

	handler, ok := s.interpretVerb(verb, p, f)
	if ok {
		result, err := handler.HandleVerb(p, f)
		if err != nil {
			fmt.Println("ERROR:", err)
		} else {
			fmt.Println(result)
		}
	}
}

// If an existing Entity is not found, then create Entity as a literal
func (s *Salu) interpretEntity(estring string) Entity {
	e, ok := s.entities[estring]
	if !ok {
		// not a known Entity, is it a literal?
		if n, err := strconv.Atoi(estring); err == nil {
			return NumberLiteral(n)
		}
		return StringLiteral(estring)
	}
	return e
}

func (s *Salu) interpretVerb(verb string, patient, focus Entity) (VerbHandler, bool) {
	ptype := EntityType(patient)
	ftype := EntityType(focus)

	verbs := s.verbs[verb]
	for _, v := range verbs {
		if v.PatientType == ptype && v.FocusType == ftype {
			return v.Handler, true
		}
	}

	return nil, false
}

func (s *Salu) RegisterVerb(verb string, handler VerbHandler, patientType string, focusType string) {
	var v Verb
	v.PatientType = patientType
	v.FocusType = focusType
	v.Handler = handler
	s.verbs[verb] = append(s.verbs[verb], v)
}
