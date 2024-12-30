package test

import (
	"log"
	"testing"

	"github.com/haou-no-tamago/Guts/expression"
	"github.com/haou-no-tamago/Guts/problem"
)

func TestProblem(t *testing.T) {
	v := testVariablesInUse()
	e0 := expression.NewExpressionNobody().Add(v[0], 1).Add(v[1], 2)
	e1 := expression.NewExpressionNobody().Add(v[2], 2).Add(v[3], 3)
	model := problem.NewProblem("test").SetDirection("max")
	model.AddVariables(v)
	model.SetObjective(e0)
	model.AddConstraint("c", e1, ">", 2)
	log.Println(model.Format())

	solver := NewMySolver()
	model.Solve(solver, true, 0.5)
}
