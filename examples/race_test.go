package examples

import (
	"fmt"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	var result int
	go func() {
		for i := 0; i < 1e6; i++ {
			result++
			fmt.Println(result, "+")
		}
	}()

	go func() {
		for i := 0; i < 1e6; i++ {
			result--
			fmt.Println(result, "-")
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(result)
}

func TestRaceInStruct(t *testing.T) {
	type s struct {
		f1 []int
		f2 []int
	}

	var result s
	go func() {
		for i := 0; i < 1e5; i++ {
			result.f1 = append(result.f1, i)
			fmt.Print(".")
		}
	}()

	go func() {
		for i := 0; i < 1e5; i++ {
			result.f2 = append(result.f2, i)
			fmt.Print(".")
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(len(result.f1))
	fmt.Println(len(result.f2))
}
