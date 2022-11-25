package examples

import (
	"fmt"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	var m []int
	for i := 0; i < 10; i++ {
		m = append(m, i)
	}
	fmt.Println(m)
}

func TestSlice(t *testing.T) {
	s := []string{"a", "b", "c"}
	fmt.Println(s[:len(s)-1])
}

func TestSlice2(t *testing.T) {
	s := []string{"a", "b", "c"}
	fmt.Println(s[1:2:3])
	fmt.Println(s[1:3])
}
