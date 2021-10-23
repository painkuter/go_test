package benchmarks

import (
	"testing"
)

func BenchmarkMapLen(b *testing.B) {

	arr1 := make([]int64, 900)
	arr2 := make([]int64, 1000)
	b.Run("1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make([]int64, len(arr2))
			_ = m
		}
	})

	b.Run("2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make([]int64, min(len(arr2), len(arr1)))
			_ = m
		}
	})

	b.Run("map1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make(map[int64]struct{}, len(arr2))
			_ = m
		}
	})

	b.Run("map2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make(map[int64]struct{}, min(len(arr2), len(arr1)))
			_ = m
		}
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
