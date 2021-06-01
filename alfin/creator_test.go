package alfin

import (
	"testing"
)

func TestPyGen(t *testing.T) {
	ent := "ExampleItem"
	// $ python service_meta.py entity_abi Example
	out, err := PyGen("service_meta.py", "entity_abi", ent)
	if err != nil {
		panic(err)
	}
	println("PyGen OUT => ", out)
}

func TestMainEntConf(t *testing.T) {
	configFile := "conf/maint_common.yml"
	maintConf := ReadMaintConf(configFile)
	Pretty(maintConf)
}
