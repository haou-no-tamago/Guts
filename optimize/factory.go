package optimize

import (
	"fmt"
	"time"

	"github.com/haou-no-tamago/Guts/constant"
)

type Factory struct {
	modelContent   string
	runTime        float64
	status         string
	objectiveValue float64
	variableValue  map[string]float64
}

func (f *Factory) Read(modelContent string) error {
	f.modelContent = modelContent
	return nil
}

// customized solve

func (f *Factory) Clear() error {
	f.modelContent = ""
	f.status = constant.UnsolvedStatus
	for k := range f.variableValue {
		delete(f.variableValue, k)
	}
	return nil
}

func (f *Factory) RunTime() float64 {
	return f.runTime
}

func (f *Factory) Status() string {
	return f.status
}

func (f *Factory) ObjectiveValue() float64 {
	return f.objectiveValue
}

func (f *Factory) VariableValue(name string) float64 {
	return f.variableValue[name]
}

func (f *Factory) ModelContent() string {
	return f.modelContent
}

func (f *Factory) Describe() string {
	return fmt.Sprintf(
		"Status: %s\nRun time: %v\nObj: %v\n",
		f.status,
		f.runTime,
		f.objectiveValue,
	)
}

func (f *Factory) StartUp() {
	f.Clear()
	f.runTime = float64(time.Now().UnixMilli())
}

func (f *Factory) TimeOut() {
	f.runTime = (float64(time.Now().UnixMilli()) - f.runTime) / 1e3
}

func (f *Factory) SetModelContent(modelContent string) *Factory {
	f.modelContent = modelContent
	return f
}

func (f *Factory) SetSolveStatus(status string) *Factory {
	f.status = Status
	return f
}

func (f *Factory) SetObjectiveValue(objectiveValue float64) *Factory {
	f.objectiveValue = objectiveValue
	return f
}

func (f *Factory) SetVariableValue(name string, value float64) *Factory {
	f.variableValue[name] = value
	return f
}
