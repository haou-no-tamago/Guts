package problem

import (
	"github.com/haou-no-tamago/Guts/constraint"
	"github.com/haou-no-tamago/Guts/expression"
	"github.com/haou-no-tamago/Guts/variable"
)

func (p *Problem) SetDirection(direction string) *Problem {
	p.direction = directionCheck(direction)
	return p
}

func (p *Problem) AddVariable(variable *variable.Variable) *Problem {
	variableCheck(p, variable)
	p.variables[variable] = struct{}{}
	return p
}

func (p *Problem) AddVariables(variables []*variable.Variable) *Problem {
	for _, v := range variables {
		p.AddVariable(v)
	}
	return p
}

func (p *Problem) SetObjective(objective *expression.Expression) *Problem {
	p.objective = objective
	return p
}

func (p *Problem) AddConstraint(
	name string,
	leftHandSide *expression.Expression,
	sense string,
	rightHandSide float64,
) *Problem {
	constraintInUse := constraint.NewConstraint(
		name, leftHandSide, sense, rightHandSide,
	)
	constraintCheck(p, constraintInUse)
	p.constraints[constraintInUse] = struct{}{}
	return p
}

func (p *Problem) AddConstraints(constraints []*constraint.Constraint) *Problem {
	for _, c := range constraints {
		constraintCheck(p, c)
		p.constraints[c] = struct{}{}
	}
	return p
}
