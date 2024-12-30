package expression

import "github.com/haou-no-tamago/Guts/variable"

func (e *Expression) Add(v *variable.Variable, c float64) *Expression {
	e.polynomial[v] = e.polynomial[v] + c
	if e.polynomial[v] == 0.0 {
		delete(e.polynomial, v)
	}
	return e
}

func (e *Expression) Multiply(coefficient float64) *Expression {
	for v, c := range e.polynomial {
		e.polynomial[v] = c * coefficient
		if e.polynomial[v] == 0.0 {
			delete(e.polynomial, v)
		}
	}
	return e
}

func (e *Expression) AddExpression(expression *Expression) *Expression {
	for v, c := range expression.polynomial {
		e.Add(v, c)
	}
	return e
}

func (e *Expression) SetCoefficient(v *variable.Variable, c float64) *Expression {
	e.Add(v, c-e.polynomial[v])
	return e
}
