package main

import (
	"fmt"
	"io/ioutil"
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
	for key := range m {
		fmt.Println(key)
	}
}
