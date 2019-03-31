package examples

import (
	"runtime"
	"strconv"
	"strings"
	"testing"
)

const (
	size = 10000
)

func TestJustCopyWithoutGC(t *testing.T) {
	m1 := newTestMap()
	m2 := make(map[int]string)
	_ = m2
	var i int
	for k := range m1 {
		i++

		m2[k] = m1[k]
		if i%100 == 0 {
			PrintMemUsage()
		}
	}
	PrintMemUsage()
}

func TestJustCopy(t *testing.T) {
	m1 := newTestMap()
	m2 := make(map[int]string)
	_ = m2
	var i int
	for k := range m1 {
		i++

		m2[k] = m1[k]
		if i%100 == 0 {
			eatATwix()
		}
	}
	eatATwix()
}

func TestReplaceEmptyString(t *testing.T) {
	m1 := newTestMap()
	m2 := make(map[int]string)
	_ = m2
	var i int
	for k := range m1 {
		i++

		m2[k] = m1[k]
		m1[k] = ""
		if i%100 == 0 {
			eatATwix()
		}
	}
	eatATwix()
}

func TestDelete(t *testing.T) {
	m1 := newTestMap()
	m2 := make(map[int]string)
	_ = m2
	var i int
	for k := range m1 {
		i++

		m2[k] = m1[k]
		delete(m1, k)
		if i%100 == 0 {
			eatATwix()
		}
	}
	eatATwix()
}

func TestCopyByPointerWithoutGC(t *testing.T) {
	m1 := newPointerTestMap()
	m2 := make(map[int]*string)
	_ = m2
	var i int
	for k := range m1 {
		i++

		m2[k] = m1[k]
		if i%100 == 0 {
			PrintMemUsage()
		}
	}
	PrintMemUsage()
}

func TestJustCopyByPointer(t *testing.T) {
	m1 := newPointerTestMap()
	m2 := make(map[int]*string)
	_ = m2
	var i int
	for k := range m1 {
		i++

		m2[k] = m1[k]
		if i%100 == 0 {
			eatATwix()
		}
	}
	eatATwix()
}

func TestReplaceByNil(t *testing.T) {
	m1 := newPointerTestMap()
	m2 := make(map[int]*string)
	_ = m2
	var i int
	for k := range m1 {
		i++

		m2[k] = m1[k]
		m1[k] = nil
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
	for i := 0; i < 10000; i++ {
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

func newPointerTestMap() map[int]*string {
	m1 := make(map[int]*string, size)
	for i := 0; i < size; i++ {
		s := generateString()
		m1[i] = &s
	}
	return m1
}

func eatATwix() {
	runtime.GC()
	PrintMemUsage()
}
