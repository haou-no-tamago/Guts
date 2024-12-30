package problem

import (
	"log"

	"github.com/haou-no-tamago/Guts/constant"
	"github.com/haou-no-tamago/Guts/constraint"
	"github.com/haou-no-tamago/Guts/variable"
)

func directionCheck(direction string) string {
	directionInUse := ""
	switch direction {
	case "max", "Max", "MAX", "Maximize":
		directionInUse = constant.MaximizedProblem
	case "min", "Min", "MIN", "Minimize":
		directionInUse = constant.MinimizedProblem
	default:
		log.Fatalln("Unknown model direction")
	}
	return directionInUse
}

func variableCheck(p *Problem, variable *variable.Variable) {
	_, exists := p.variables[variable]
	if exists {
		log.Printf("Varibale %s is already in problem %s", variable.Name(), p.Name())
	}
}

func constraintCheck(p *Problem, constraint *constraint.Constraint) {
	_, exists := p.constraints[constraint]
	if exists {
		log.Printf("Constraint %s is already in problem %s", constraint.Name(), p.Name())
	}
}
