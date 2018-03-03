package app

import (
	"fmt"
	"runtime"
	"time"
)

func test_2() {
	// 1 proc: with sleep
	// expects NOT full execution for goroutine
	runtime.GOMAXPROCS(1)
	k := 0
	go func() {
		for i := 0; i < 1<<32-1; i++ {
			time.Sleep(time.Microsecond)
			k = i
		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(k)

	test_2_2()
}

func test_2_2() {
	// 1 proc: without sleep
	// expects full execution for goroutine
	runtime.GOMAXPROCS(1)

	k := 0
	go func() {
		for i := 0; i < 1<<32-1; i++ {
			k = i
		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(k)

	test_2_3()
}

func test_2_3(){
	// 4 procs: without sleep
	// expects NOT full execution for goroutine
	runtime.GOMAXPROCS(4)

	k := 0
	go func() {
		for i := 0; i < 1<<32-1; i++ {
			k = i
		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(k)
}

