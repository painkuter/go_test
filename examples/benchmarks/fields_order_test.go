package benchmarks

import (
	"testing"
)

type s1 struct {
	st1 string
	i1  int64
	st2 string
	i2  int64
	st3 string
	i3  int64
	st4 string
	i4  int64
	st5 string
	i5  int64
}

type s2 struct {
	st1 string
	st2 string
	st3 string
	st4 string
	st5 string
	i1  int64
	i2  int64
	i3  int64
	i4  int64
	i5  int64
}

func TestFields(t *testing.T) {
	r := make([]s1, 1e6)
	_ = r
}

func TestFields2(t *testing.T) {
	r := make([]s2, 1e6)
	_ = r
}

func BenchmarkFields(b *testing.B) {
	b.Run("s1", func(b *testing.B) {
		r := make([]s1, 1e6)
		_ = r
	})

	b.Run("s2", func(b *testing.B) {
		r := make([]s2, 1e6)
		_ = r
	})
}
