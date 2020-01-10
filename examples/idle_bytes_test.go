package examples

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"testing"
	"time"
)

type DifficultStruct struct {
	ID   int
	Name *string
}

func TestSegmentationFault(t *testing.T) {
	// make map
	// add element
	// get elem
	// use element
	// delete element
	// free memory
	// use elem

	name := "test  name"
	testMap := make(map[int]*DifficultStruct)
	testMap[1] = &DifficultStruct{
		ID:   2,
		Name: &name,
	}

	go func() {
		if testMap[1] == nil {
			return
		}
		time.Sleep(time.Second)
		fmt.Println(*testMap[1].Name)
	}()
	time.Sleep(time.Millisecond)
	delete(testMap, 1)
	runtime.GC()
	debug.FreeOSMemory()
	time.Sleep(time.Second)
}
