package benchmarks

import (
	"runtime"
	"strconv"
	"testing"
)

type void struct{}

func ToMap1(v []string) map[string]void {
	out := map[string]void{}
	for _, i := range v {
		out[i] = void{}
	}
	return out
}

func ToMap2(v []int64) map[int64]void {
	m := make(map[int64]void, len(v))
	for _, i := range v {
		m[i] = void{}
	}
	return m
}

var benchmarkV = func() []int64 {
	v := make([]int64, 2000000)
	for i := range v {
		v[i] = int64(i)
	}
	return v
}()

var benchmarkVStr = func() []string {
	v := make([]string, 2000000)
	for i := range v {
		v[i] = strconv.Itoa(i)
	}
	return v
}()

func BenchmarkToMap1(b *testing.B) {
	data := benchmarkVStr

	b.Run("string", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for N := 0; N < b.N; N++ {
			ToMap1(data)
		}
	})
	runtime.GC()
	data2 := benchmarkV
	b.Run("int64", func(b *testing.B) {

		b.ReportAllocs()
		b.ResetTimer()
		for N := 0; N < b.N; N++ {
			ToMap2(data2)
		}
	})
}
