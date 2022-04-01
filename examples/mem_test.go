package examples

import (
	"testing"
	"unsafe"
)

func TestName(t *testing.T) {
	t.Log(unsafe.Sizeof("123"))
	t.Log(unsafe.Sizeof([...]byte{'1', '2', '3'}))
}
