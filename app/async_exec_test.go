package app

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestAsyncExec(t *testing.T) {
	j := int64(0)

	taskCount := 100
	maxGoroutines := 14

	concurrentGoroutines := make(chan struct{}, maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		concurrentGoroutines <- struct{}{}
	}

	// done := make(chan bool)

	/*	waitForAllJobs := make(chan bool)

		go func() {

			waitForAllJobs <- true
		}()
	*/
	for i := 0; i < taskCount; i++ {
		<-concurrentGoroutines
		go func() {
			exec(&j)
			// done <- true
		}()
	}

	// for i := 0; i < taskCount; i++ {
	// 	<-done
	// 	concurrentGoroutines <- struct{}{}
	// }

	// <-waitForAllJobs
}

func exec(j *int64) {
	time.Sleep(10 * time.Millisecond)
	atomic.AddInt64(j, 1)
	fmt.Println(runtime.NumGoroutine())
}
