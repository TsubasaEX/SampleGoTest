package trandnum

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockRand struct{ mock.Mock }

func newMockRand() *mockRand { return &mockRand{} }

func (m *mockRand) randomInt(max int) int {
	args := m.Called(max)
	return args.Int(0)
}

func TestDivByRand(t *testing.T) {
	// get our mockRand
	m := new(mockRand)
	// specify our return value. Since the code in divByRand
	// passes 10 into randomInt, we pass 10 in as the argument
	// to go with randomInt, and specify that we want the
	// method to return 6.
	m.On("randomInt", 10).Return(5)

	// now run divByRand and assert that we got back the
	// return value we expected, just like in a Go test that
	// doesn't use Testify Mock.
	quotient := divByRand(30, m)
	if quotient != 5 {
		t.Errorf("expected quotient to be 5, got %d", quotient)
	}

	// check that randomInt was called with the number 10;
	// if not then the test fails
	m.AssertCalled(t, "randomInt", 10)
}

func TestDivByRandCantDivideByZero(t *testing.T) {
	m := new(mockRand)
	m.On("randomInt", 10).Return(0)

	quotient := divByRand(30, m)
	if quotient != 30 {
		t.Errorf("expected quotient to be 30, got %d", quotient)
	}
}
