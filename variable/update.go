package variable

func (v *Variable) SetCategory(category string) *Variable {
	variableCategoryCheck(category)
	return v
}

func (v *Variable) SetValue(value float64) *Variable {
	v.value = value
	return v
}

func (v *Variable) SetBound(lowerBound float64, upperBound float64) *Variable {
	variableBoundCheck(lowerBound, upperBound)
	v.lowerBound = lowerBound
	v.upperBound = upperBound
	return v
}

func (v *Variable) SetFixed(value float64) *Variable {
	v.SetValue(value)
	v.SetBound(value, value)
	return v
}

func (v *Variable) SetFree() *Variable {
	return variableBoundReset(v)
}
