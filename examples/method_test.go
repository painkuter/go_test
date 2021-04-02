package examples

import (
	"fmt"
	"testing"
)

type obj struct {
	field int
}

func (o obj) setValue(value int) {
	o.field = value
}

func TestMethod(t *testing.T) {
	o := obj{}
	o.setValue(1)
	fmt.Println(o.field)
}
