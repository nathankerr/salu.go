package salu

import (
	"log"
	"strconv"
)

type Salu struct {
	entities map[string]entity
}

func NewSalu() *Salu {
	s := new(Salu)
	s.entities = make(map[string]entity)
	
 	e := make(propertyset)
	e["name"] = "Nathan"
	s.entities["user"] = e

	return s
}

func (s *Salu) Eval(verb, patient, focus string) {
	p := s.interpretEntity(patient)
	log.Println("[patient]:", p)
	
	f := s.interpretEntity(focus)
	log.Println("[focus]:", f)
}

// If an existing entity is not found, then create entity as a literal
func (s *Salu) interpretEntity(estring string) entity {
	e, ok := s.entities[estring]
	if !ok {
		// not a known entity, is it a literal?
		if n, err := strconv.Atoi(estring); err == nil {
			return n
		}
		return estring
	}
	return e
}
