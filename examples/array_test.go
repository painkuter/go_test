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
