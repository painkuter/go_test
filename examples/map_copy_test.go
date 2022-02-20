package examples

import (
	"fmt"
	"testing"
)

type str struct {
	cnt int
}

func TestMainMap(t *testing.T) {
	withPointer()
	fmt.Println("")
	withoutPointer()
}

func withPointer() {
	v := &str{}

	firstMap := make(map[int]*str)
	firstMap[1] = v

	v.cnt = 2

	secondMap := make(map[int]*str)
	secondMap[2] = v

	fmt.Println("with pointer")
	fmt.Println(firstMap[1].cnt)
	fmt.Println(secondMap[2].cnt)
}

func withoutPointer() {
	v := str{}

	firstMap := make(map[int]str)
	firstMap[1] = v

	v.cnt = 2

	secondMap := make(map[int]str)
	secondMap[2] = v

	fmt.Println("without pointer")
	fmt.Println(firstMap[1].cnt)
	fmt.Println(secondMap[2].cnt)
}

func TestMapLen(t *testing.T) {
	m := make(map[int]struct{}, 10)
	m[1] = struct{}{}
	m[2] = struct{}{}
	m[3] = struct{}{}

	fmt.Println(len(m))
}
