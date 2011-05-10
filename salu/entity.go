package salu

type Entity interface{}

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
