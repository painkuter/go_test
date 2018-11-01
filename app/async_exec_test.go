package app

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestAsyncExec(t *testing.T) {
	j := int64(0) // execution counter

	taskCount := 100 // all tasks number
	maxGoroutines := 14

	concurrentGoroutines := make(chan struct{}, maxGoroutines)

	// fill channel's buffer
	for i := 0; i < maxGoroutines; i++ {
		concurrentGoroutines <- struct{}{}
	}

	done := make(chan bool)
	waitForAllJobs := make(chan bool)

	// goroutine starter:
	go func() {
		for i := 0; i < taskCount; i++ {
			<-done // wait for some DONE than start new goroutine
			concurrentGoroutines <- struct{}{}
		}
		waitForAllJobs <- true // waiting for all goroutines done
	}()

	for i := 0; i < taskCount; i++ {
		<-concurrentGoroutines
		go func() {
			exec(&j)
			done <- true
		}()
	}

	<-waitForAllJobs
}

func exec(j *int64) {
	time.Sleep(time.Duration(atomic.LoadInt64(j)) * time.Millisecond)
	atomic.AddInt64(j, 1)
	fmt.Printf("%d:%d\n", *j, runtime.NumGoroutine())
}
