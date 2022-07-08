package examples

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) { // падает при запуске с -race
	k := 0
	wg := sync.WaitGroup{}
	c := 4
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			k++
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(t, k, 4)
}

func TestSumAtomic(t *testing.T) {
	var k int32
	wg := sync.WaitGroup{}
	c := 4
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			atomic.AddInt32(&k, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	assert.EqualValues(t, k, 4)
}
