package common

import (
	"encoding/json"
	"fmt"
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

type Flight struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Price       int    `json:"price"`
}

func TestPrettyPrint(t *testing.T) {
	flight := Flight{
		Origin:      "GLA",
		Destination: "JFK",
		Price:       2000,
	}

	b, err := json.MarshalIndent(flight, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
	println()
}
