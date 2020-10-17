package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"testing"
)

// объединяет 2 массива, удаляя дубликаты
func TestUnion(t *testing.T) {
	f1, err := ioutil.ReadFile("arr1")
	if err != nil {
		panic(err)
	}
	f2, err := ioutil.ReadFile("arr2")
	if err != nil {
		panic(err)
	}

	arr1 := strings.Split(string(f1), "\n")
	arr2 := strings.Split(string(f2), "\n")

	m := make(map[string]struct{})

	for _, elem := range append(arr1, arr2...) {
		z := strings.TrimSpace(elem)
		if z == "" || z == "\n" {
			continue
		}
		m[z] = struct{}{}
	}

	fmt.Println("len=", len(m))
	r := make([]string, 0, len(m))
	for key := range m {
		r = append(r, key)
	}
	sort.Strings(r)
	for _, value := range r {
		fmt.Println(value)
	}
}
