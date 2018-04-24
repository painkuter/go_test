package app

import (
	"fmt"
	"time"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapInGoroutines(t *testing.T) {
	// expects random values for i
	mutex := sync.Mutex{}
	m := make(map[int]int)
	size := 1<<10
	for i := 0; i < size; i++ {
		go func() {
			mutex.Lock()
			m[i]++
			mutex.Unlock()
		}()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Len = %v,\t size = %v\n", len(m), size)

	assert.Equal(t, len(m)<size, true)
}
