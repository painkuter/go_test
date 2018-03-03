package app

import (
	"fmt"
	"runtime"
)

func test_1() {
	runtime.GOMAXPROCS(1)

	count := 1<<63 - 1

	fmt.Println(count)
}
