package expression

import (
	"log"

	"github.com/haou-no-tamago/Guts/constant"
	"github.com/haou-no-tamago/Guts/variable"
)

type Expression struct {
	name       string
	polynomial map[*variable.Variable]float64
}

func NewExpressionEmpty(name string) *Expression {
	return &Expression{
		name:       name,
		polynomial: make(map[*variable.Variable]float64),
	}
}

func NewExpressionNobody() *Expression {
	return &Expression{
		name:       constant.DefaultExpressionName,
		polynomial: make(map[*variable.Variable]float64),
	}
}

func NewExpression(
	name string,
	variables []*variable.Variable,
	coefficients []float64,
) *Expression {
	if len(variables) != len(coefficients) || len(variables) == 0 {
		log.Fatalln("Error length during initialize expression !")
	}
	e := NewExpressionEmpty(name)
	for i, v := range variables {
		e.Add(v, coefficients[i])
	}
	return e
}
