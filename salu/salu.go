package salu

import (
	"fmt"
	"io"
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

	registerBuiltins(s)

	e := make(PropertySet)
	e["name"] = StringLiteral("Nathan")
	s.entities["user"] = e

	return s
}

func (s *Salu) Console(input io.Reader, output io.Writer) {
	parser := NewParser(input)

	for {
		sen := parser.Parse()

		result := s.eval(sen)
		fmt.Fprintf(output, "> %v\n", result)
	}
}

func (s *Salu) eval(sen *sentence) string {
	p := s.interpretEntity(sen.patient)
	f := s.interpretEntity(sen.focus)

	for _, v := range s.verbs[sen.verb] {
		// log.Println("Looking at verbs")
		p, ok := s.getEntityAs(p, v.PatientType)
		if !ok {
			continue
		}

		f, ok := s.getEntityAs(f, v.FocusType)
		if !ok {
			continue
		}

		result, err := v.Handler.HandleVerb(p, f)
		if err != nil {
			return "ERROR (salu.Eval): " + err.String()
		}
		return result.String()
	}
	return "ERROR (salu.Eval): no suitable verb implementation found for " + sen.verb
}

func (s *Salu) getEntityAs(e Entity, etype string) (Entity, bool) {
	ok := false

	switch etype {
	case "NumberLiteral":
		e, ok = e.(NumberLiteral)
	case "StringLiteral":
		e, ok = e.(StringLiteral)
	case "PropertySet":
		e, ok = e.(PropertySet)
	case "EntityList":
		e, ok = e.(EntityList)
	}
	return e, ok
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
