package tcalc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	testInput  []int
	testOutput []int
}

func (s *Suite) SetupSuite() {
	fmt.Println("SetupSuite()...")
}

func (s *Suite) TearDownSuite() {
	fmt.Printf("TearDowmnSuite()...\n")
}

func (s *Suite) SetupTest() {
	fmt.Printf("SetupTest()... \n")
	s.testInput = []int{1, 2, 3, 4, 5}
	s.testOutput = []int{3, 4, 5, 6, 7}
}

func (s *Suite) TearDownTest() {
	fmt.Printf("TearDownTest()... \n")
}

func (s *Suite) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest=> suiteName: %s, testName: %s\n", suiteName, testName)
	if testName == "TestAddTwoSecond" {
		s.testInput = []int{10, 20, 30, 40, 50}
		s.testOutput = []int{12, 22, 32, 42, 52}
	}
}

func (s *Suite) AfterTest(suiteName, testName string) {
	fmt.Printf("AferTest=> suiteName: %s, testName: %s\n", suiteName, testName)
}

func (s *Suite) TestAddTwo() {
	for i, val := range s.testInput {
		result := AddTwo(val)
		assert.Equal(s.T(), s.testOutput[i], result)
	}
}

func (s *Suite) TestAddTwoSecond() {
	for i, val := range s.testInput {
		result := AddTwo(val)
		assert.Equal(s.T(), s.testOutput[i], result)
	}
}

func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}
