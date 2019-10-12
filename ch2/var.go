package main

import (
	"fmt"
)

func main() {
	testNew()
	testAssign()
}

func testNew() {
	p := new(int)
	fmt.Println(p)
	*p = 21
	fmt.Println(*p)
}

func testAssign() {
	x := 1
	p := new(bool)
	*p = true
	count := make([]bool, 3)
	count[x] = *p
	fmt.Println(count)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
