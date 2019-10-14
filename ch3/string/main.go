package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// test1()
	// testMultiple()
	testUtf8()
}

func test1() {
	s := "hello world"
	fmt.Println(s[0], s[7])
	s2 := "hello"
	fmt.Println(s == s2, s < s2)
}

func testMultiple() {
	var s = `这个
跨越多行
了。`
	fmt.Println(s)
}

func testUtf8() {
	s := "hello 时接"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); i++ {
		r, size := utf8.DecodeRuneInString()
		fmt.Println("%d\t%c\n")
	}
}
