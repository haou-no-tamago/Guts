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
	if value > upperBound || value < lowerBound {
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
