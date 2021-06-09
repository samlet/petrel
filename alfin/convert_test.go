package alfin

import (
	"strconv"
	"testing"
)

func TestConvertInt(t *testing.T) {
	i, err := strconv.Atoi("00001")
	if err != nil {
		panic(err)
	}
	println(i)
}