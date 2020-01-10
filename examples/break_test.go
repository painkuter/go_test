package examples

import (
	"fmt"
	"runtime"
	"testing"
)

func TestReturnInsteadBreak(t *testing.T) {
	a := 10
	line()
	func() {
		fmt.Println(a)
		a++
		line()
		return
	}()
	fmt.Println(a)
	line()
}

func line() {
	_, _, line, _ := runtime.Caller(1)
	fmt.Printf("line: %d\n", line)
}
