package benchmarks

import (
	"testing"
	"time"
)

type wh struct {
	id        int
	name      string
	updatedAt time.Time
}

func BenchmarkAppend(b *testing.B) {
	size := 1000
	for i := 0; i < b.N; i++ {
		testData1 := make([]wh, size)
		testData2 := make([]wh, size)
		result := make([]wh, 0, size*2)
		result = append(testData1, testData2...)
		_ = result
	}
}
