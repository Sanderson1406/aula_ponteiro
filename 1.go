package main

import "fmt"

func main() {
	x := 40
	y := 25
	var pntx *int = &x
	var pnty *int = &y
	fmt.Println(pntx, pnty)
	pntx, pnty = pnty, pntx
	fmt.Println(pntx, pnty)
}
