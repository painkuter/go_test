package app

import (
	"fmt"
	"time"
)

func test_4(){
	var k int
	for i := 0; i < 1<<6; i ++{
		//i_ := i
		go func(j int) {
			k = j
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(k)
}