package examples

import (
	"log"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testArgs struct {
	targetLoc bool
	whInfo    bool
	whAddr    bool
	whTT      bool
	dm        bool
	ts        bool
}

func (ta *testArgs) getWhInfo() {
	ta.whInfo = true
}

func (ta *testArgs) getTargetLoc() {
	ta.targetLoc = true
}

func (ta *testArgs) getWhAddr() {
	if !ta.whInfo {
		log.Fatal("whInfo")
	}
	ta.whAddr = true
}

func (ta *testArgs) getWhTT() {
	if !ta.whInfo {
		log.Fatal("whInfo")
	}
	ta.whTT = true
}

func (ta *testArgs) getDM() {
	if !ta.whInfo {
		log.Fatal("whInfo")
	}
	if !ta.targetLoc {
		log.Fatal("targetLoc")
	}
	ta.dm = true
}

func (ta *testArgs) calcTS() {
	if !ta.dm {
		log.Fatal("dm")
	}
	if !ta.whTT {
		log.Fatal("whTT")
	}
	if !ta.whAddr {
		log.Fatal("whAddr")
	}
	ta.ts = true
}

func TestConcurrencyCourier(t *testing.T) {
	var (
		wg1 sync.WaitGroup //  нужна для ожидания targetLoc
		wg2 sync.WaitGroup //  нужна для ожидания whInfo
		wg3 sync.WaitGroup //  нужна для ожидания dm, whAddr, whTT
	)
	ta := &testArgs{}

	wg1.Add(1)
	wg3.Add(1)
	go func() {
		ta.getTargetLoc()
		wg1.Done()
		wg3.Done()
	}()
	wg1.Add(1)
	wg2.Add(1)
	go func() {
		ta.getWhInfo()
		wg2.Done()
		wg1.Done()
	}()

	wg3.Add(2)

	go func() {
		wg2.Wait() // готов склад
		go func() {
			ta.getWhAddr()
			wg3.Done()
		}()
		go func() {
			ta.getWhTT()
			wg3.Done()
		}()
	}()

	wg1.Wait() // готова targetLoc
	wg2.Wait() // готов склад
	wg3.Add(1)
	go func() {
		ta.getDM()
		wg3.Done()
	}()

	wg3.Wait() // готов dm, whAddr, whTT

	ta.calcTS()
	assert.True(t, ta.ts)
}

/*

    |<====targetLoc
    |  |<==whInfo
|<==|==|whAddr
|<==|==|whTT
|<==dm
ts

*/
