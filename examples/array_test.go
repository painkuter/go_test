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
