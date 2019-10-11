package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo() {
	start := time.Now()
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(start))
}
func echo2() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(start))
}
func test1() {
	for index, arg := range os.Args {
		fmt.Printf("index:%v , arg: %v \n", index, arg)
	}
}
func main() {
	// var s, sep string
	// for i := 0; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)
	// echo()
	// echo2()
	// test1()
	echo()
	echo2()
}
