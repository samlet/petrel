package common

import (
	"hash/fnv"
	"strconv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func ParseId(idval string) int {
	i, err := strconv.Atoi(idval)
	if err != nil {
		//log.Fatalf("parse id value %s fail: %v", idval, err)
		return int(hash(idval))
	}
	return i
}

