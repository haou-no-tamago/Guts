package test

import (
	"log"
	"testing"

	"github.com/haou-no-tamago/Guts/constant"
	"github.com/haou-no-tamago/Guts/variable"
)

func testVariablesInUse() []*variable.Variable {
	var variables []*variable.Variable
	variables = append(variables, variable.NewVariable("v0", constant.RealVariable, 0.0, 5.0))
	variables = append(variables, variable.NewVariableSlim("v1"))
	variables = append(variables, variable.NewIntegerVariable("v2", 0.0, 3.0))
	variables = append(variables, variable.NewBinaryVariable("v3"))
	return variables
}

func TestVariable(t *testing.T) {
	variables := testVariablesInUse()
	for _, v := range variables {
		log.Println(v.Detail())
	}
}
