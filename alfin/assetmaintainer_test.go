package alfin

import (
	rice "github.com/GeertJohan/go.rice"
	"log"
	"testing"
)

func TestAssets(t *testing.T) {
	// find a rice.Box
	assertBox, err := rice.FindBox("assets")
	if err != nil {
		log.Fatal(err)
	}
	printSample(assertBox, "Ballot.json")
	printSample(assertBox, "person.json")
}

func printSample(assertBox *rice.Box, asset string) {
	// get file contents as string
	asset, err := assertBox.String(asset)
	if err != nil {
		log.Fatal(err)
	}

	var bodySample string
	if len(asset) > 500 {
		bodySample = asset[0:500] + " ..."
	} else {
		bodySample = asset
	}

	println(bodySample)
}
