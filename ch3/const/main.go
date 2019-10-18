package main

import (
	"fmt"
	"strings"
)

func main() {
	// t1()
	// t2()
	// t3()
	// t4()
	// t5()
	// t6()
	t7()
}

func t1() {
	const (
		e  = 2.71828182845904523536028747135266249775724709369995957496696763
		pi = 3.14159265358979323846264338327950288419716939937510582097494459
	)
	fmt.Println(e, pi)
}

func t2() {
	fmt.Println(parseIpv4("192.168.1.1"))
}

const IPV4Len = 4

func parseIpv4(s string) []string {

	p := strings.Split(s, ".")
	return p
}

type Weekday int

func t3() {
	const (
		Sunday Weekday = 1 << iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday,
		Wednesday,
		Thursday,
		Friday,
		Saturday)
}

type Flags uint

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadCast
	FlagLoopBack
	FlagPointToPoint
	FlagMulticast
)

func t4() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, isUp(v))
	turnDown(&v)
	fmt.Printf("%b %t \n", v, isUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, isUp(v))
	fmt.Printf("%b %t\n", v, isCast(v))

}

func isUp(v Flags) bool {
	return v&FlagUp == FlagUp
}
func turnDown(v *Flags) {
	*v &^= FlagUp // 位清除
	// 100100 a
	// 010101 b
	// 110001
	// 100000 r
	// a中位置为1且b中对应为0才为1
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadCast
}

func isCast(v Flags) bool {
	return v&(FlagBroadCast|FlagMulticast) != 0
}

func t5() {
	const (
		_   = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB // 1099511627776             (exceeds 1 << 32)
		PiB // 1125899906842624
		EiB // 1152921504606846976
		ZiB // 1180591620717411303424    (exceeds 1 << 64)
		YiB // 1208925819614629174706176
	)
	fmt.Printf("%b %[1]d \n", KiB)
	fmt.Printf("%b %[1]d \n", TiB)
	fmt.Printf("%b %[1]d \n", PiB)
	fmt.Printf("%b %[1]d \n", EiB)
	fmt.Println(YiB / ZiB)
	// fmt.Println(ZiB)
	// fmt.Printf("%b  \n", ZiB)
	// fmt.Printf("%b  \n", YiB)
}

func t6() {

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0";   5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float
	var f2 float64 = 3 + 0i
	// 例如0、0.0、0i和'\u0000'虽然有着相同的常量值，但是它们分别对应无类型的整数、无类型的浮点数、无类型的复数和无类型的字符等不同的常量类型
	fmt.Println(f2)
}

func t7() {
	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000')
}
