package component

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Display(c Component) string {
	return fmt.Sprintf("%s(%p)", c.Name(), c)
}

func Abstract(nameAddress string) (string, uintptr) {
	fields := strings.Split(nameAddress, "(")
	name := fields[0]
	address, err := strconv.ParseUint(strings.Split(fields[1], ")")[0], 0, 64)
	if err != nil {
		log.Println(err)
		log.Fatalf("Pointer %v convert failed !\n", address)
	}
	ptr := uintptr(address)
	return name, ptr
}
