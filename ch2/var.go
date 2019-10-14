package main

import (
	"ch2/conv"
	"ch2/popcount"
	"fmt"
)

func main() {
	// testNew()
	// testAssign()
	// testConv()
	fmt.Println(popcount.PopCount(12))

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

func testConv() {
	c := conv.FToC(conv.Fahrenheit(234))
	fmt.Println(c)
}
