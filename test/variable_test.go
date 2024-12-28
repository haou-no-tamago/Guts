package test

import (
	"fmt"
	"testing"

	"github.com/haou-no-tamago/Guts/variable"
)

func TestVariable(t *testing.T) {
	fmt.Println("hello world !")
	v := variable.NewVariableSlim("v")

	fmt.Printf("%s = %v\n", v.Name(), v.Value())
}
