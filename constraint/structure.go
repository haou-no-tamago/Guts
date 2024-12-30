package constraint

import "github.com/haou-no-tamago/Guts/expression"

type Constraint struct {
	name          string
	leftHandSide  *expression.Expression
	sense         string
	rightHandSide float64
}

func NewConstraint(
	name string,
	leftHandSide *expression.Expression,
	sense string,
	rightHandSide float64,
) *Constraint {
	senseInUse := senseCheck(sense)
	return &Constraint{
		name:          name,
		leftHandSide:  leftHandSide,
		sense:         senseInUse,
		rightHandSide: rightHandSide,
	}
}
