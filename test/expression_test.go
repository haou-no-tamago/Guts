package test

import (
	"log"
	"math"
	"testing"

	"github.com/haou-no-tamago/Guts/expression"
)

func TestExpression(t *testing.T) {
	e := expression.NewExpressionEmpty("e")
	variables := testVariablesInUse()
	for i, v := range variables {
		v.SetValue(float64(4 - i))
		log.Println(v.Detail())
		e.Add(v, -1.0*float64(i)*math.Pow(-1.0, float64(i)))
	}
	log.Println(e.Detail())
}
