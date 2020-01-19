package primitive

import (
	"log"
	"strconv"
)

func MustAtoI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return i
}