package examples

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
)

func TestChan(t *testing.T) {
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
}

func TestChanSync(t *testing.T) {
	c := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()
	go func() {

		//for v1, ok := range c {

		//}

		wg.Done()

	}()
}

func TestFromEridanSilver(t *testing.T) {
	hydro, oxy, clo, sy := make(chan struct{}), make(chan struct{}), make(chan struct{}), make(chan struct{})
	defer close(clo)
	fu := func(c chan struct{}, el string) {
		select {
		case <-clo:
			return
		case c <- struct{}{}:
			fmt.Print(el)
			sy <- struct{}{}
		}
	}

	for i := 0; i < 50; i++ {
		go fu(hydro, "H")
		go fu(oxy, "O")
	}

	go func() {
		for {
			<-hydro
			<-hydro
			<-oxy
			for i := 0; i < 3; i++ {
				<-sy
			}
			fmt.Println("")
		}
	}()

	time.Sleep(2 * time.Second)
}

func TestH2O(t *testing.T) {
	o := semaphore.NewWeighted(2)
	//o.TryAcquire(2)
	h := semaphore.NewWeighted(2)
	//h.TryAcquire(2)
	pr := semaphore.NewWeighted(3)

	for i := 0; i < 50; i++ {
		go func() {
			ok := o.TryAcquire(1)
			if !ok {
				return
			}
			//time.Sleep(10)
			fmt.Print("H")
			h.Release(1)
			pr.TryAcquire(1)
		}()
		go func() {
			ok := h.TryAcquire(2)
			if !ok {
				return
			}
			//time.Sleep(10)

			fmt.Print("O")
			o.Release(2)
			pr.TryAcquire(1)
		}()
	}

	go func() {
		fmt.Println()
		pr.Release(3)
	}()
	time.Sleep(10 * time.Second)
}
