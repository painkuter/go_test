package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceCopy(t *testing.T) {
	foo := []string{"a", "b", "c", "d"}
	bar := []string{"e", "f", "g", "h"}
	copy(bar, foo[2:])

	assert.Equal(t, []string{"c", "d", "g", "h"}, bar)
}

func TestDeleteFromSlice(t *testing.T) {
	data := []string{"a", "b", "c"}
	data = data[1:]
	assert.Equal(t, []string{"b", "c"}, data)
}
