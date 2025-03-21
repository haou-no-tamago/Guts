package constraint

import (
	"fmt"

	"github.com/haou-no-tamago/Guts/constant"
	"github.com/haou-no-tamago/Guts/expression"
)

func (c *Constraint) Name() string {
	return c.name
}

func (c *Constraint) LeftHandSide() *expression.Expression {
	return c.leftHandSide
}

func (c *Constraint) Sense() string {
	return c.sense
}

func (c *Constraint) RightHandSide() float64 {
	return c.rightHandSide
}

func (c *Constraint) Feasible() bool {
	switch c.sense {
	case constant.GreaterSense:
		return c.leftHandSide.Value() >= c.rightHandSide - constant.ErrorTolerance
	case constant.EqualSense:
		return c.leftHandSide.Value() <= c.rightHandSide + constant.ErrorTolerance && c.leftHandSide.Value() >= c.rightHandSide - constant.ErrorTolerance
	case constant.LessSense:
		return c.leftHandSide.Value() <= c.rightHandSide + constant.ErrorTolerance
	}
	return false
}

func (c *Constraint) Format() string {
	return fmt.Sprintf(
		"%s %s %v",
		c.leftHandSide.Format(), c.sense, c.rightHandSide,
	)
}

func (c *Constraint) Detail() string {
	return fmt.Sprintf("%s (Feasible: %v)", c.Format(), c.Feasible())
}
