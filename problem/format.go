package problem

import (
	"fmt"
	"log"

	"github.com/haou-no-tamago/Guts/component"
	"github.com/haou-no-tamago/Guts/constant"
	"github.com/haou-no-tamago/Guts/variable"
)

func (p *Problem) Format() string {
	format := fmt.Sprintf("\\* %s *\\\n", p.name)
	if p.objective != nil {
		format += fmt.Sprintf(
			"%s\n %s: %s\n",
			p.direction,
			component.Display(p.objective),
			p.objective.Format(),
		)
	} else {
		log.Fatalf("Empty problem %s !\n", p.name)
	}

	format += "Subject To\n"
	for c := range p.constraints {
		format += fmt.Sprintf("%s: %s\n", component.Display(c), c.Format())
	}

	var integers []*variable.Variable
	var binaries []*variable.Variable

	format += "Bounds\n"
	for v := range p.variables {

		switch v.Category() {
		case constant.BinaryVariable:
			binaries = append(binaries, v)
		case constant.IntegerVariable:
			integers = append(integers, v)
		}

		if v.Category() == constant.BinaryVariable && v.LowerBound() == constant.DefaultBinaryLowerBound && v.UpperBound() == constant.DefaultBinaryUpperBound {
			continue
		} else if v.UpperBound() != constant.DefaultUpperBound {
			format += fmt.Sprintf(
				"%v <= %s <= %v \n",
				v.LowerBound(),
				component.Display(v),
				v.UpperBound(),
			)
		} else {
			format += fmt.Sprintf(
				"%v <= %s \n",
				v.LowerBound(),
				component.Display(v),
			)
		}
	}

	if len(integers) > 0 {
		format += fmt.Sprintf("%s\n", constant.IntegerVariable)
		for _, v := range integers {
			format += fmt.Sprintf(" %s", component.Display(v))
		}
		format += "\n"
	}

	if len(binaries) > 0 {
		format += fmt.Sprintf("%s\n", constant.BinaryVariable)
		for _, v := range binaries {
			format += fmt.Sprintf(" %s", component.Display(v))
		}
		format += "\n"
	}

	format += "End\n"
	return format
}
