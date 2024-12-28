package variable

import (
	"github.com/haou-no-tamago/Guts/constant"
)

type Variable struct {
	name       string
	category   string
	value      float64
	lowerBound float64
	upperBound float64
}

func NewVariable(
	name string,
	category string,
	lowerBound float64,
	upperBound float64) *Variable {

	variableCategoryCheck(category)

	variableBoundCheck(lowerBound, upperBound)

	return &Variable{
		name:       name,
		category:   category,
		value:      constant.DefaultLowerBound,
		lowerBound: lowerBound,
		upperBound: upperBound,
	}
}

func NewVariableSlim(name string) *Variable {
	return &Variable{
		name:       name,
		category:   constant.RealVariable,
		value:      constant.DefaultLowerBound,
		lowerBound: constant.DefaultLowerBound,
		upperBound: constant.DefaultUpperBound,
	}
}
