package app

import (
	"fmt"
	"time"
	"sync"
)

func test_1() {
	// expects random values for i
	mutex := sync.Mutex{}
	m := make(map[int]int)
	for i := 0; i < 1<<10; i++ {
		go func() {
			mutex.Lock()
			m[i]++
			mutex.Unlock()
		}()
	}
	recover()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Done. Len = %v\n", len(m))
}
