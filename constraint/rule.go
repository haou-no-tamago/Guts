package constraint

import (
	"log"

	"github.com/haou-no-tamago/Guts/constant"
)

func senseCheck(sense string) string {
	switch sense {
	case ">=", ">":
		return constant.GreaterSense
	case "=", "==":
		return constant.EqualSense
	case "<=", "<":
		return constant.LessSense
	default:
		log.Fatalln("Error in constraint sense !")
	}
	return sense
}
