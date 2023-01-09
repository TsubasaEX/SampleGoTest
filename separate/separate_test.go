package separate_test

import (
	"testing"

	. "github.com/TsubasaEX/SampleGoModTest/separate"
)

func TestAddTwo(t *testing.T) {
	if AddTwo(2) != 4 {
		t.Error("Error in Testing")
	}
}
