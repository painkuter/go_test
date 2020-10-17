package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"testing"
)

// элементы, которые есть только во втором массиве, но нет в первом
func TestDiff(t *testing.T) {
	f1, err := ioutil.ReadFile("item")
	if err != nil {
		panic(err)
	}
	f2, err := ioutil.ReadFile("product_search")
	if err != nil {
		panic(err)
	}

	arr1 := strings.Split(string(f1), "\n")
	arr2 := strings.Split(string(f2), "\n")

	m := make(map[string]struct{})

	for _, elem := range arr1 {
		z := strings.TrimSpace(elem)
		if z == "" || z == "\n" {
			continue
		}
		m[z] = struct{}{}
	}

	var result []string
	for _, elem := range arr2 {
		z := strings.TrimSpace(elem)
		if z == "" || z == "\n" {
			continue
		}
		if _, ok := m[z]; !ok {
			result = append(result, z)
		}
	}

	fmt.Println("len=", len(result))

	sort.Strings(result)
	for _, value := range result {
		fmt.Println(value)
	}
}
