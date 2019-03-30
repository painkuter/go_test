package examples

import (
	"runtime"
	"strconv"
	"strings"
	"testing"
)

const (
	size = 1000
)

func TestDelete(t *testing.T) {
	m1 := newTestMap()
	m2 := make(map[int]string, size)
	_ = m2
	var i int
	for k := range m1 {
		i++

		m1[k] = ""
		if i%100 == 0 {
			eatATwix()
		}
	}
	eatATwix()
}

func getMaxValue(m map[int]int) int {
	var result int
	for _, v := range m {
		if result < v {
			result = v
		}
	}
	return result
}

func generateString() string {
	var sb strings.Builder
	for i := 0; i < 100000; i++ {
		sb.WriteString(strconv.Itoa(i))
	}
	result := sb.String()
	return result
}

func newTestMap() map[int]string {
	m1 := make(map[int]string, size)
	for i := 0; i < size; i++ {
		m1[i] = generateString()
	}
	return m1
}

func eatATwix() {
	runtime.GC()
	PrintMemUsage()
}
