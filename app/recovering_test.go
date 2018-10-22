package app

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var result string

func TestRecovering(t *testing.T) {
	g()
	assert.Equal(t, "recovered from  test panic\n", result)
}

func g() {
	defer testRecover()
	f()
	return
}

func testRecover() {
	if r := recover(); r != nil {
		result = fmt.Sprintln("recovered from ", r)
	}
}

func f() {
	panic("test panic")
}
