package calc

import "testing"

func TestAddTwo(t *testing.T) {
	t.Log("Testing TestAddTwo...")
	if AddTwo(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
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
