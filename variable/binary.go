package variable

import "github.com/haou-no-tamago/Guts/constant"

func NewBinaryVariable(name string) *Variable {
	return &Variable{
		name:       name,
		category:   constant.BinaryVariable,
		value:      constant.DefaultLowerBound,
		lowerBound: 0.0,
		upperBound: 1.0,
	}
}
