package constant

// Problem
const (
	MaximizedProblem = "Maximize"
	MinimizedProblem = "Minimize"
)

// Sense
const (
	GreaterSense = ">="
	LessSense    = "<="
	EqualSense   = "="
)

// Variable
const (
	RealVariable    = "Real"
	IntegerVariable = "Integer"
	BinaryVariable  = "Binary"
)

// Solver
const (
	UnsolvedStatus   = "Unsolved"
	OptimalStatus    = "Optimal"
	AbortedStatus    = "Aborted"
	InfeasibleStatus = "Infeasible"
)
