package variable

import (
	"math"

	"github.com/haou-no-tamago/Guts/constant"
)

func variableCategoryCheck(category string) bool {
	switch category {
	case constant.RealVariable, constant.IntegerVariable, constant.BinaryVariable:
		return true
	default:
		return false
	}
}

func variableBoundCheck(lowerBound float64, upperBound float64) bool {
	return lowerBound <= upperBound
}

func variableValueCheck(value float64, category string, lowerBound float64, upperBound float64) bool {
	if value > upperBound + constant.ErrorTolerance || value < lowerBound - constant.ErrorTolerance {
		return false
	}

	switch category {
	case constant.RealVariable:
	case constant.IntegerVariable:
		error := math.Min(math.Abs(value-float64(int(value))), 1.0-math.Abs(value-float64(int(value))))
		if error > constant.ErrorTolerance {
			return false
		}
	case constant.BinaryVariable:
		error := math.Min(math.Abs(value-1.0), math.Abs(value-0.0))
		if error > constant.ErrorTolerance {
			return false
		}
	default:
		return false
	}

	return true
}

func variableBoundReset(v *Variable) *Variable {
	switch v.category {
	case constant.BinaryVariable:
		v.lowerBound = constant.DefaultBinaryLowerBound
		v.upperBound = constant.DefaultBinaryUpperBound
	case constant.IntegerVariable, constant.RealVariable:
		v.lowerBound = constant.DefaultLowerBound
		v.upperBound = constant.DefaultUpperBound
	}
	return v
}
