package benchmarks

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkFmt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprintf("%s-%d", "qwerty", 123456)
	}
}

func BenchmarkPlus(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = "qwerty" + "-" + strconv.Itoa(123456)
	}
}
