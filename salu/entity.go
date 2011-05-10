package salu

import (
	"strconv"
)

type entity interface{}

type numberliteral int
type stringliteral string
type propertyset map[string]entity

// If an existing entity is not found, then create entity as a literal
func (s *Salu) interpretEntity(estring string) entity {
	e, ok := s.entities[estring]
	if !ok {
		// not a known entity, is it a literal?
		if n, err := strconv.Atoi(estring); err == nil {
			return n
		}
		return stringliteral(estring)
	}
	return e
}

func entityType(e entity) (string) {
	switch _ := e.(type) {
	case numberliteral:
		return "numberliteral"
	case stringliteral:
		return "stringliteral"
	case propertyset:
		return "propertyset"
	}
	
	return "unknown"
}
