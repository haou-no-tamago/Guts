package problem

import (
	"github.com/haou-no-tamago/Guts/constraint"
	"github.com/haou-no-tamago/Guts/expression"
	"github.com/haou-no-tamago/Guts/variable"
)

func (p *Problem) Name() string {
	return p.name
}

func (p *Problem) Direction() string {
	return p.direction
}

func (p *Problem) Variables() []*variable.Variable {
	var variables []*variable.Variable
	for v := range p.variables {
		variables = append(variables, v)
	}
	return variables
}

func (p *Problem) Objective() *expression.Expression {
	return p.objective
}

func (p *Problem) Constraints() []*constraint.Constraint {
	var constraints []*constraint.Constraint
	for c := range p.constraints {
		constraints = append(constraints, c)
	}
	return constraints
}

func (p *Problem) Feasible() bool {
	for v := range p.variables {
		if !v.Feasible() {
			return false
		}
	}
	for c := range p.constraints {
		if !c.Feasible() {
			return false
		}
	}
	return true
}
