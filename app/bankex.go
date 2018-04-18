package app

import "fmt"

func test_0(){
	foo := []string{"a","b","c","d"}
	bar := []string{"e","f","g","h"}
	copy (bar, foo[2:])
	fmt.Print(bar)
	//cd
}
