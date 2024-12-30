package constraint

import "github.com/haou-no-tamago/Guts/variable"

func (c *Constraint) SetCoefficient(v *variable.Variable, coefficient float64) *Constraint {
	c.leftHandSide.SetCoefficient(v, coefficient)
	return c
}

func (c *Constraint) SetSense(sense string) *Constraint {
	c.sense = senseCheck(sense)
	return c
}

func (c *Constraint) SetRightHandSide(rightHandSide float64) *Constraint {
	c.rightHandSide = rightHandSide
	return c
}
