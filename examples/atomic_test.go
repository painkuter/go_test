package examples

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
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
