package main

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"go.etcd.io/etcd/pkg/v3/stringutil"
)

func TestGenerateRnd(t *testing.T) {
	// length + count
	result := stringutil.RandomStrings(16, 1)
	fmt.Println(result)
}

const myChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"

func TestGenerateUUID(t *testing.T) {
	// length + count
	result := uuid.New()
	fmt.Println(result)
}
