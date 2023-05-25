package main

import "fmt"

func main() {
	var bo bool = true
	var bopn *bool = &bo
	fmt.Println(bopn)
	bo2 := false
	var bopn2 *bool = &bo2
	fmt.Println(bopn2)
}
