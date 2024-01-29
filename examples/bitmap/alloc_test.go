package bitmap

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/kelindar/bitmap"
)

var listLength int = 1e7

type memoryStats struct {
	allocated      string
	totalAllocated string
	system         string
}

// getMemory finds the allocated, totalAllocated and system memory
func getMemory() memoryStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return memoryStats{
		allocated:      fmt.Sprintf("%d KB", m.Alloc/1024),
		totalAllocated: fmt.Sprintf("%d KB", m.TotalAlloc/1024),
		system:         fmt.Sprintf("%d KB", m.Sys/1024),
	}
}
func TestAlloc(t *testing.T) {
	t.Run("struct array", func(t *testing.T) {

		sts := make([]struct{ b1, b2, b3, b4, b5, b6, b7, b8 bool }, 1_000_000)
		_ = sts
		//{1120 KB 1120 KB 12797 KB}
		//ints := make([]int32, 1_000_000)
		//_ = ints
		//{4051 KB 4051 KB 17405 KB}
		//bits := make([]bool, 1_000_000)
		//_ = bits
		//{31396 KB 31396 KB 46077 KB}
		fmt.Println(getMemory())
	})

	t.Run("int64 array", func(t *testing.T) {
		arr := make([]int32, 100_000_000)
		_ = arr
		fmt.Println(getMemory())
	})

	t.Run("kelindar/bitmap", func(t *testing.T) {
		var books bitmap.Bitmap
		books.Set(100_000_000)
		books.Contains(100_000_000)
		fmt.Println(getMemory())
	})

	t.Run("real test bitmap", func(t *testing.T) {
		var locCache []struct {
			bitMap bitmap.Bitmap
			id     int
		}

		for i := 0; i < 1500; i++ {
			var bm bitmap.Bitmap

			for wh := 0; wh < 1_000_000; wh++ {
				n := rand.Int31()
				if n%2 == 1 {
					bm.Set(uint32(wh))
				}
			}
			locCache = append(locCache, struct {
				bitMap bitmap.Bitmap
				id     int
			}{bitMap: bm, id: i})
		}
		fmt.Println(getMemory())

	})

	t.Run("real test slice", func(t *testing.T) {
		var locCache [][]int32

		for i := 0; i < 1500; i++ {
			loc := make([]int32, 1_000_000)
			for wh := 0; wh < 1_000_000; wh++ {
				n := rand.Int31()
				if n%2 == 1 {
					loc[wh] = int32(wh)
				}
			}
			locCache = append(locCache, loc)
		}
		fmt.Println(getMemory())

	})

}

//{136 KB 136 KB 11391 KB}
