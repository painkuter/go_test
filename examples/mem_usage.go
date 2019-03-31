package examples

import (
	"runtime"

	"go_test/visualization_server/result"
)

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	//fmt.Printf("Alloc = %v kb", convertBytes(m.Alloc))
	//fmt.Printf("\tTotalAlloc = %v kb", convertBytes(m.TotalAlloc))
	//fmt.Printf("\tSys = %v kb", convertBytes(m.Sys))
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)
	result.ResultsKV(
		"Alloc", convertBytes(m.Alloc),
		"TotalAlloc", convertBytes(m.TotalAlloc),
		"Sys", convertBytes(m.Sys))
}

func convertBytes(b uint64) uint64 {
	return b / 1024
	//return b / 1024 / 1024
}
