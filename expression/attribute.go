package expression

import (
	"fmt"

	"github.com/haou-no-tamago/Guts/component"
	"github.com/haou-no-tamago/Guts/variable"
)

func (e *Expression) Copy(name string) *Expression {
	return &Expression{
		name:       name,
		polynomial: e.polynomial,
	}
}

func (e *Expression) Name() string {
	return e.name
}

func (e *Expression) Polynomial() map[*variable.Variable]float64 {
	return e.polynomial
}

func (e *Expression) Coefficient(v *variable.Variable) float64 {
	return e.polynomial[v]
}

func (e *Expression) Value() float64 {
	value := 0.0
	for v, c := range e.polynomial {
		value += v.Value() * c
	}
	return value
}

func (e *Expression) Format() string {
	format := ""
	cnt := 0
	for v, c := range e.polynomial {
		if c > 0 {
			if cnt == 0 {
				format += fmt.Sprintf("%v %s ", c, component.Display(v))
			} else {
				format += fmt.Sprintf("+%v %s ", c, component.Display(v))
			}
		} else if c < 0 {
			format += fmt.Sprintf("%v %s ", c, component.Display(v))
		}
		cnt++
	}
	return format
}

func (e *Expression) Detail() string {
	return fmt.Sprintf("%s: %s, value = %v\n", e.name, e.Format(), e.Value())
}
