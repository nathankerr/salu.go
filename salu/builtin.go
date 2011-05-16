// Builtin verbs
package salu

import (
	"os"
)

func registerBuiltins(s *Salu) {
	var is isa
	s.RegisterVerb("isa", is, "NumberLiteral", "StringLiteral")
	s.RegisterVerb("isa", is, "StringLiteral", "StringLiteral")
	s.RegisterVerb("isa", is, "PropertySet", "StringLiteral")
	s.RegisterVerb("isa", is, "EntityList", "StringLiteral")
}

// Promotes an entity to the requested type
type isa int

func (is isa) HandleVerb(patient, focus Entity) (Entity, os.Error) {
	return nil, nil
}
