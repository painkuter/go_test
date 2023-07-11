package examples

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestInline(t *testing.T) {
	funcWithInnerCall(1)
	funcWithInnerCall(2)
}

func funcWithInnerCall(i int) {
	f := func() {
		fmt.Println("test1", i)
	}
	fmt.Println(reflect.TypeOf(f))
	spew.Dump(f)
	f()
}
