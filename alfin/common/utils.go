package common

import (
	"log"
	"strconv"
)

func ParseId(idval string) int {
	i, err := strconv.Atoi(idval)
	if err != nil {
		log.Fatalf("parse id value %s fail: %v", idval, err)
	}
	return i
}

