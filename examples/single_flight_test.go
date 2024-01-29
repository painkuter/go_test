package examples

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
	"golang.org/x/sync/singleflight"
)

var sfGroup singleflight.Group

func TestSingleFlight(t *testing.T) {
	sem := semaphore.NewWeighted(10)
	ctx := context.Background()
	for {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}
		go func() {
			defer sem.Release(1)
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			defer cancel()
			_, err := callWithTimeout(ctx)
			if err != nil {
				fmt.Println("err:", err.Error())
			}
		}()
	}
}

func callWithTimeout(ctx context.Context) (bool, error) {
	res, err, shared := sfGroup.Do("1", func() (interface{}, error) {
		ch := make(chan struct{})
		var result bool
		go func() {
			result = callDatabase(ctx, ch)
		}()
		select {
		case <-ch:
			fmt.Println("done")
		case <-ctx.Done():
			fmt.Println("halted callWithTimeout")
		}

		return result, nil
	})
	if shared {
		fmt.Println("shared")
	}
	return res.(bool), err
}

func callDatabase(ctx context.Context, ch chan struct{}) bool {
	time.Sleep(time.Millisecond * time.Duration(250+rand.Intn(500)))
	return true
}
