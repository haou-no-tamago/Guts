package variable

import "fmt"

func (v *Variable) Copy(name string) *Variable {
	return &Variable{
		name:       name,
		category:   v.category,
		value:      v.value,
		lowerBound: v.lowerBound,
		upperBound: v.upperBound,
	}
}

func (v *Variable) Name() string {
	return v.name
}

func (v *Variable) Category() string {
	return v.category
}

func (v *Variable) Value() float64 {
	return v.value
}

func (v *Variable) LowerBound() float64 {
	return v.lowerBound
}

func (v *Variable) UpperBound() float64 {
	return v.upperBound
}

func (v *Variable) Feasible() bool {
	return variableValueCheck(v.value, v.category, v.lowerBound, v.upperBound)
}

func (v *Variable) Detail() string {
	return fmt.Sprintf(
		"%s (%s: %v, %v) = %v",
		v.name,
		v.category, v.lowerBound, v.upperBound,
		v.value,
	)
}
