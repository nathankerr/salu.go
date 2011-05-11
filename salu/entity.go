package salu

import (
	"fmt"
)

type Entity interface{
	String() string
}

type NumberLiteral int
type StringLiteral string
type PropertySet map[string]Entity

func EntityType(e Entity) (string) {
	switch _ := e.(type) {
	case NumberLiteral:
		return "NumberLiteral"
	case StringLiteral:
		return "StringLiteral"
	case PropertySet:
		return "PropertySet"
	}

	return "unknown"
}

func (n NumberLiteral) String() string {
	return fmt.Sprintf("%v", int(n))
}

func (s StringLiteral) String() string {
	return string(s)
}

func (p PropertySet) String() string {
	return "PropertySet"
//	return fmt.Sprintf("%v", p)
}
