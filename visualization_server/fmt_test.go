package main

import "testing"

func TestResults(t *testing.T) {
	for i := 0; i < 10; i++ {
		Results("%d, %d\n", 10, 20)
	}
}
