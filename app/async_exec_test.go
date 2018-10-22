package app

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var j int

func TestAsyncExec(t *testing.T) {
	assert.Equal(t, 1, 1)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			exec()
		}()
		wg.Done()
	}
	wg.Wait()
}

func exec() {
	time.Sleep(10 * time.Millisecond)
	j++
}
