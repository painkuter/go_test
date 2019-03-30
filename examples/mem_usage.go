package examples

import (
	"fmt"
	"runtime"
)

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v kb", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v kb", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v kb", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024
	//return b / 1024 / 1024
}
