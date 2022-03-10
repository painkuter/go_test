package examples

import (
	"testing"

	"github.com/golang/geo/r3"
	"github.com/golang/geo/s2"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	a := s2.Point{Vector: r3.Vector{}}
	b := s2.Point{Vector: r3.Vector{X: 0, Y: 1}}
	c := s2.Point{Vector: r3.Vector{X: 1, Y: 0}}
	l := s2.LoopFromPoints([]s2.Point{a, b, c})
	result := l.ContainsPoint(s2.Point{Vector: r3.Vector{X: 0.25, Y: 0.25}})
	assert.True(t, result)
}
