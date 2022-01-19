package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	file, err := ioutil.ReadFile("duplicates")
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(file), "\n")
	set := make(map[string]int)
	for _, s := range arr {
		if s == "" {
			continue
		}
		if _, ok := set[s]; !ok {
			set[s] = 0
		}
		set[s]++
	}
	for s, i := range set {
		if i > 1 {
			fmt.Println("duplicated ", s, " count ", i)
		}
	}
}
