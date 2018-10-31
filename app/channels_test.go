package app

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestChannels(t *testing.T) {
	var done = make(chan bool, 1)
	var timer = make(chan bool, 1)

	go func() {
		fmt.Println("Timer started")
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Timer ended")
		timer <- true
		return
	}()

	go func() {
		fmt.Println("Update started")
		time.Sleep(time.Millisecond * 1000)
		fmt.Println("Update ended")
		done <- true
		return
	}()
	select {
	case <-done:
	case <-timer:
		{
			go func() {
				fmt.Println("Select started")
				time.Sleep(time.Millisecond * 1000)
				fmt.Println("Select ended")
				done <- true
			}()
		}
	}
	<-done

	//TODO: check output?
}

func TestWriteToClosedChannelBuffered(t *testing.T) { // expected panic
	ch := make(chan bool, 1)
	ch <- true
	close(ch)
	ch <- true
}

func TestWriteToClosedChannel(t *testing.T) { // expected lock
	ch := make(chan bool)
	ch <- true
	close(ch)
	ch <- true
}

func TestCh(t *testing.T) { // expected lock
	ch := make(chan bool)
	go func() {
		time.Sleep(100*time.Millisecond)
		runtime.Gosched()
		runtime.LockOSThread()
		runtime.Version()
	}()
	ch <- true
	// ch <- false
}
