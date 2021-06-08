package alfin

import (
	"fmt"
	"testing"
)

func TestMapify(t *testing.T) {
	var myarr = []string{"key1", "val1", "key2", "val2"}
	fmt.Println(myarr)
	fmt.Println(Mapify(myarr))
}
