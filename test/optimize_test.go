package test

import (
	"log"
	"testing"
	"time"

	"github.com/haou-no-tamago/Guts/optimize"
)

type MySolver struct {
	optimize.Factory
}

func NewMySolver() *MySolver {
	ms := &MySolver{}
	ms.Clear()
	return ms
}

func (ms *MySolver) Solve(tee bool, timeLimit float64) optimize.Solver {
	ms.StartUp()
	defer ms.TimeOut()

	time.Sleep(time.Duration(timeLimit * 1e3 * float64(time.Millisecond)))
	ms.SetObjectiveValue(3.1415926)
	return ms
}

func useOpt(solver optimize.Solver) string {
	solver.Read("abcd")
	solver.Solve(true, 1.5)
	return solver.Describe()
}

func TestOptimize(t *testing.T) {
	opt := NewMySolver()
	s := useOpt(opt)
	log.Println(s)
}
