package examples

import (
	"fmt"
	"testing"
)

func TestOrAnd(t *testing.T) {
	if 1 == 0 || 1 == 1 && 2 == 3 {
		println(true)
		return
	}
	println(false)
}

func TestAnd(t *testing.T) {
	var arr1 []int
	if len(arr1) > 0 && arr1[0] == 1 {
		println(1)
	}
	println(0)
}

func TestErr(t *testing.T) {
	var err error
	fmt.Println(err.Error())
}
