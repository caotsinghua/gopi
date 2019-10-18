package main

import (
	"fmt"
)

func main() {
	// t1()
	// t2()
	// t3()
	// t4()
	t5()
}

func t1() {
	var x uint8 = 1<<1 | 1<<5 // "00100010", the set {1, 5}
	var y uint8 = 1<<1 | 1<<2
	to8b(x)
	to8b(y)
	to8b(x & y)  // 取交集
	to8b(x | y)  // 取并集
	to8b(x ^ y)  // 异或 // 不同
	to8b(x &^ y) // 位清空 （x为1，y为0）
	fmt.Println("===输出哪些位置为1")
	r := make([]uint, 0)
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			r = append(r, i)
		}
	}
	fmt.Println(r)

	var a int8 = -1
	// 1000 0001
	// 1111 1111 补码
	// 右移  1111 111 首位补1
	fmt.Printf("%08b\n", a>>8)
}
func to8b(n uint8) {
	fmt.Printf("%08b\n", n)
}

func t2() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}

func t3() {
	var f float32 = 3.141
	i := int(f)
	fmt.Println(f, i)
	f = 1.99
	fmt.Println(float64(f))
	f2 := 1e100
	i = int(f2)
	fmt.Println(f2, i)
}

func t4() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	// %[1]x 中的[1]表示再次使用第一个操作数
	// %#[1]x表示 输出时生成0 0x 或0X前缀
	// %X表示使用大写形式
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}

func t5() {
	var ascii rune = 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q %[1]T \n", ascii)
	fmt.Printf("%d %[1]c %[1]q \n", unicode)
	fmt.Printf("%d %[1]q \n", newline)
}
