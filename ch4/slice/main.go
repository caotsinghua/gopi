package main

import (
	"fmt"
	"unicode"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func t1() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:]) // 操作的是切片，改变了数组
	fmt.Println(a)
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func testappend() {
	var runes []rune
	p := []byte("hello,时接")
	for _, r := range "hello,时接" {
		runes = append(runes, r)
	}
	fmt.Printf("%q %q\n", string(runes), p)
}

func t3() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = append(x, i)
		fmt.Printf("%d cap%d \t %v\n", i, cap(y), y)
		x = y
	}
}

func nonempty(strs []string) []string {
	i := 0
	for _, s := range strs {
		if s != "" {
			strs[i] = s
			i++
		}
	}
	return strs[:i]
}

func remove(arr []int, i int) []int {
	copy(arr[i:], arr[i+1:])
	return arr[:len(arr)-1]
	// return arr
}
func t4() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(remove(arr, 1))
}
func main() {
	// t1()
	// s1 := make([]int, 1, 2)
	// s1 = make([]int, 0, 0)
	// var s1 []int
	// fmt.Println(s1, s1 == nil)
	// testappend()
	// t3()
	// t4()
	// fmt.Println(cleanSame([]int{1, 2, 3, 3, 3, 3, 4, 4, 5, 5, 5, 6}))
	fmt.Println(string(cleanSpace([]byte("cc   12 12   21    3c     "))))
}

// 123333445556
//
func cleanSame(s []int) []int {
	c := 0
	for i := 1; i < len(s); i++ {
		if s[c] != s[i] {
			c++
			s[c] = s[i]
		}
	}

	return s[:c+1]
}

func cleanSpace(s []byte) []byte {
	c := 0
	for i := 1; i < len(s); i++ {
		if !(unicode.IsSpace(rune(s[i])) && s[c] == s[i]) {
			c++
			s[c] = s[i]
		}
	}
	return s[:c+1]
}
