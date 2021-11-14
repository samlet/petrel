package misc

import (
	"github.com/samlet/petrel/alfin/modules/catalog/seedcreators"
	"testing"
)

func TestLoadSeeds(t *testing.T) {
	seedcreators.LoadSeeds(true)
}
