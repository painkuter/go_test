package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type iCast interface {
	cast()
}

type caster struct {
}

func getCaster() interface{} {
	return caster{}
}

func TestCaster_NotOk(t *testing.T) {
	buf := getCaster()
	testCaster, ok := buf.(*iCast)
	assert.True(t, !ok)
	assert.Nil(t, testCaster)
}
