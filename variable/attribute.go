package variable

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
