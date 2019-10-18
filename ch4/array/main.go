package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// t1()
	// t2()
	// t3()
	// t4()
	// t5()
	t6()

}

func t1() {
	var a [3]int
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]
}

func t2() {
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"
	fmt.Println(q)
}

func t3() {
	q := [...]int{1, 2, 3, 4}
	i := 1
	fmt.Println(len(q), q, q[i])
}

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func t4() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
	s2 := [...]string{1: "aa"}
	fmt.Println(s2, len(s2))
	r := []int{99: -1}
	fmt.Println(r, len(r))
}

func t5() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	fmt.Println(a != b, a != c)
}
func t6() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte{'x', 'y'})
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(c1)
}
