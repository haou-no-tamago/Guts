package variable

import "github.com/haou-no-tamago/Guts/constant"

func NewIntegerVariable(name string, lowerBound float64, upperBound float64) *Variable {
	return &Variable{
		name:       name,
		category:   constant.IntegerVariable,
		value:      constant.DefaultLowerBound,
		lowerBound: lowerBound,
		upperBound: upperBound,
	}
}
