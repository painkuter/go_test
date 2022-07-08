package examples

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestName(t *testing.T) {
	t.Log(unsafe.Sizeof("123"))
	t.Log(unsafe.Sizeof([...]byte{'1', '2', '3'}))
}

func TestMem(t *testing.T) {
	type foo struct {
		from time.Time
	}

	x := &foo{}
	y := foo{}
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(y))
}
