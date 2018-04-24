package app

import (
	"fmt"
	"time"
)

func test_3() {
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
