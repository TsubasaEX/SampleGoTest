package trandnum

import "math/rand"

type randNumberGenerator interface {
	randomInt(max int) int
}

// our standard-library implementation is an empty
// struct whose randomInt method calls math/rand.Intn
type standardRand struct{}

func (s standardRand) randomInt(max int) int {
	return rand.Intn(max)
}

func divByRand(numerator int, r randNumberGenerator) int {
	denominator := 1 + r.randomInt(10)
	return numerator / denominator
}
