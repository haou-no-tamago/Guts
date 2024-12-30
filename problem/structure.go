package problem

import (
	"github.com/haou-no-tamago/Guts/constraint"
	"github.com/haou-no-tamago/Guts/expression"
	"github.com/haou-no-tamago/Guts/variable"
)

type Problem struct {
	name        string
	direction   string
	variables   map[*variable.Variable]struct{}
	objective   *expression.Expression
	constraints map[*constraint.Constraint]struct{}
}

func NewProblem(name string) *Problem {
	return &Problem{
		name:        name,
		variables:   make(map[*variable.Variable]struct{}),
		constraints: make(map[*constraint.Constraint]struct{}),
	}
}
