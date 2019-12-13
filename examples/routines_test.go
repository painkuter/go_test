package examples

import (
	"fmt"
	"testing"
	"time"
)

func TestRoutines(t *testing.T) {
	for i := 0; i < 1e6*4; i++ {
		go func(j int) {
			fmt.Printf("%d,", j)
			for {
				time.Sleep(time.Millisecond * 1e2)
			}
		}(i)
	}
	fmt.Println("DONE")
	time.Sleep(time.Minute * 1e3)
}
