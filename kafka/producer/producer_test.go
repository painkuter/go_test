package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {

	buf := `2801
2786`
	messages := strings.Split(buf, "\n")
	fmt.Println(messages)
}
