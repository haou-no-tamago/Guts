package problem

import (
	"log"

	"github.com/haou-no-tamago/Guts/component"
	"github.com/haou-no-tamago/Guts/constant"
	"github.com/haou-no-tamago/Guts/optimize"
	"github.com/haou-no-tamago/Guts/variable"
)

func (p *Problem) Solve(solver optimize.Solver, tee bool, timeLimit float64) *Problem {
	if solver.Read(p.Format()) != nil {
		log.Fatalln("Error in reading model")
	}
	switch solver.Solve(tee, timeLimit).Status() {
	case constant.OptimalStatus:
		for v := range p.variables {
			v.SetValue(solver.VariableValue(component.Display(v)))
		}
	case constant.AbortedStatus:
		variableStore := make(map[*variable.Variable]float64)
		for v := range p.variables {
			variableStore[v] = v.Value()
		}
		for v := range p.variables {
			v.SetValue(solver.VariableValue(component.Display(v)))
		}
		if !p.Feasible() {
			for v := range p.variables {
				v.SetValue(variableStore[v])
			}
		}
	default:
		log.Println("Status: ", solver.Status())
	}
	return p
}
