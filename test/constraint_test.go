package test

import (
	"log"
	"testing"

	"github.com/haou-no-tamago/Guts/constraint"
	"github.com/haou-no-tamago/Guts/expression"
)

func TestConstraint(t *testing.T) {
	v := testVariablesInUse()
	e := expression.NewExpressionNobody().Add(v[0], 1.0).Add(v[1], 2.0)
	c := constraint.NewConstraint("c1", e, ">", 3.0)
	log.Println(c.Format())
}
