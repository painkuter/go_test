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

func BenchmarkConst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		write2()
	}
}

func write1() {
	const msg1 = "test message"
	_ = fmt.Sprintln(msg1)
}

const msg2 = "test message"

func write2() {
	_ = fmt.Sprintln(msg2)
}
