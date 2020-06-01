package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

//
func main() {
	f1, err := ioutil.ReadFile("./tools/arr1")
	if err != nil {
		panic(err)
	}
	f2, err := ioutil.ReadFile("./tools/arr2")
	if err != nil {
		panic(err)
	}

	arr1 := strings.Split(string(f1), "\n")
	arr2 := strings.Split(string(f2), "\n")

	m := make(map[string]struct{})

	for _, elem := range append(arr1, arr2...) {
		z := strings.TrimSpace(elem)
		if z == "" {
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
