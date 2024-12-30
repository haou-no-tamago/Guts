package optimize

type Solver interface {
	Read(modelContent string) error
	Solve(tee bool, timeLimit float64) Solver
	Clear() error
	RunTime() float64
	Status() string
	Describe() string
	ObjectiveValue() float64
	VariableValue(name string) float64
}
