package tdevelop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilerUnique(t *testing.T) {
	input := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Elliot"},
		Developer{Name: "David"},
		Developer{Name: "Alexander"},
		Developer{Name: "Eva"},
		Developer{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"David",
		"Alexander",
		"Eva",
		"Alan",
	}

	assert.Equal(t, expected, FilerUinque(input))
}

func TestNegativeFilerUnique(t *testing.T) {
	input := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Elliot"},
		Developer{Name: "David"},
		Developer{Name: "Alexander"},
		Developer{Name: "Eva"},
		Developer{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"Eva",
		"Alan",
	}

	assert.NotEqual(t, expected, FilerUinque(input))
}
