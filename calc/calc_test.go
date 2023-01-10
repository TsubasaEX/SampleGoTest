//go:build unit
// +build unit

// go test -tags=unit -v
package calc

import (
	"testing"
)

func TestAddTwo(t *testing.T) {
	t.Log("Testing TestAddTwo...")

	t.Run("should return 4", func(t *testing.T) {
		if AddTwo(2) != 4 {
			t.Error("Expected 2 + 2 to equal 4")
		}
	})

	t.Run("should return 6", func(t *testing.T) {
		if AddTwo(4) != 6 {
			t.Error("Expected 4 + 2 to equal 6")
		}
	})

	t.Cleanup(func() {
		t.Log("Cleanup ...")
	})
}

func TestTableAddTwo(t *testing.T) {
	testingData := []struct {
		input    int
		expected int
	}{
		{2, 4},
		{10, 12},
		{100, 102},
	}

	t.Log("Testing TestTableAddTwo...")
	for _, data := range testingData {
		if output := AddTwo(data.input); output != (data.expected) {
			t.Error("Test Failed: {} inputed, {} expected, {} received", data.input, data.expected, output)
		}
	}
}

type AddResult struct {
	x        int
	y        int
	expected int
}

var addResults = []AddResult{
	{1, 1, 2},
	{2, 2, 4},
	{3, 3, 6},
}

func TestAdd(t *testing.T) {
	for _, test := range addResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected Result Not Given")
		}
	}
}
