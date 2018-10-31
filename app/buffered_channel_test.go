package app

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var jj = 0

func TestBufferedChannel(t *testing.T) {
	assert.Equal(t, 1, 1)

	ch := make(chan struct{}, 10)

	for i := 0; i < 1000; i++ {
		go func() {
			ch <- struct{}{}
			exe()
			<-ch
		}()
	}
	select {
	case <-ch:
		fmt.Println("ch")
	}
	fmt.Println(jj)
}

func exe() {
	time.Sleep(10 * time.Millisecond)
	jj++
}
