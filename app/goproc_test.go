package app

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSingleProcWithSleep(t *testing.T) {
	// 1 proc: with sleep
	// expects NOT full execution for goroutine
	runtime.GOMAXPROCS(1)
	k := 0
	size := 1 << 32
	go func() {
		for i := 0; i < size; i++ {
			time.Sleep(time.Nanosecond)
			k = i
		}
	}()
	time.Sleep(time.Millisecond)

	assert.True(t, k*1000 < size)
}

func TestSingleProcWithoutSleep(t *testing.T) {
	// 1 proc: without sleep
	// expects full execution for goroutine
	runtime.GOMAXPROCS(1)

	k := 0
	size := 1 << 32
	go func() {
		for i := 0; i < size; i++ {
			k = i
		}
	}()
	time.Sleep(time.Millisecond)

	assert.Equal(t, k, size-1)
}

func TestFourProcWithoutSleep(t *testing.T) {
	// 4 procs: without sleep
	// expects NOT full execution for goroutine
	runtime.GOMAXPROCS(4)

	k := 0
	size := 1 << 32
	go func() {
		for i := 0; i < size; i++ {
			k = i
		}
	}()
	time.Sleep(time.Millisecond)

	assert.True(t, k*1000 < size)

}

func TestAdditional(t *testing.T) {
	runtime.GOMAXPROCS(1)
	k := 0
	go func() {
		for i := 0; i < 1<<16; i++ {
			k = i
		}
	}()
	fmt.Println(k)
	time.Sleep(time.Nanosecond)
	fmt.Println(k)
}
