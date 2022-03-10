package main

import (
	"fmt"
	"testing"

	"go.etcd.io/etcd/pkg/v3/stringutil"
)

func TestGenerateRnd(t *testing.T) {
	// length + count
	result := stringutil.RandomStrings(16, 1)
	fmt.Println(result)
}
