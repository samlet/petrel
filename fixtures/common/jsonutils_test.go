package common

import (
	"encoding/json"
	"testing"
)

func TestDump(t *testing.T) {
	var jsonStr = `
	{
	  "array": [
		1,
		2,
		3
	  ],
	  "boolean": true,
	  "null": null,
	  "number": 123,
	  "object": {
		"a": "b",
		"c": "d",
		"e": "f"
	  },
	  "string": "Hello World"
	}
	`

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}
	DumpMap("", jsonMap)
}
