package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"testing"
)

// пересечение 2х массивов
func TestIntersection(t *testing.T) {
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
		if _, ok := m[z]; ok {
			result = append(result, z)
		}
	}

	fmt.Println("len=", len(result))

	sort.Strings(result)
	for _, value := range result {
		fmt.Println(value)
	}
}
